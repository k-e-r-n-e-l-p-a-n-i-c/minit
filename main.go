package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {

	pid := os.Getpid()
	if pid != 1 {
		fmt.Printf("[init] PID [%v] running in spawner-only mode", pid)
	} else {
		fmt.Printf("[init] PID [%v] running in reaper mode", pid)
		go reapZombies()
	}

	command := os.Args

	if len(command) == 1 {
		fmt.Println("[init] nothing to run")
		os.Exit(1)
	}
	bin := command[1]
	args := command[2:]
	fmt.Printf("[init] %v %v\n", bin, strings.Join(args, " "))
	cmd := exec.Command(bin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Create a channel of type os.signal to receive the signals
	sigs := make(chan os.Signal, 1)
	defer close(sigs)

	// Use signal.Notify() to trap and relay required signals to our channel
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM)
	defer signal.Reset()

	// Set ProcessGroupID for child process as init process. Both will be under same process group
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// Start a routine that will receive the signals on the channel and will forward it to child process
	go func() {
		sig := <-sigs
		fmt.Printf("[init] received  %v signal for PID %v\n", sig, cmd.Process.Pid)
		cmd.Process.Signal(sig)
	}()

	// Uncomment for testing zombie process creation
	//n := 1
	//for n < 4 {
	//	_, _, _ = syscall.StartProcess("/usr/bin/sleep", []string{"sleep", "40000"}, &syscall.ProcAttr{})
	//	n++
	//}

	// Start the command
	err := cmd.Start()
	if err != nil {
		fmt.Printf("[init] failed to start process  %v\n", err)
		os.Exit(1)
	}

	// uncomment below delay for testing reaping. This simulates a delay after interrupt signal is trapped
	// you will find the reaping still happening.
	//time.Sleep(10 * time.Second)

	// Blocking code using wait()
	cmd.Wait()

}

func reapZombies() {
	for {
		var wstatus syscall.WaitStatus

		pid, err := syscall.Wait4(-1, &wstatus, syscall.WNOHANG, nil)

		// Below block is required for busy systems
		for syscall.EINTR == err {
			pid, err = syscall.Wait4(pid, &wstatus, syscall.WNOHANG, nil)
		}

		if pid <= 0 {
			time.Sleep(1 * time.Second)
		} else {
			fmt.Printf("[init] reaping zombie %v\n", pid)
			continue
		}
	}
}
