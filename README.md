# minit
A minimalistic init system written in Golang for educational purpose.

# Motivation
First of all it's *fun*. You will have to do a fair amount of study about production grade init system's to understand its internal working's. You'll get to know about how signals and traps work and how to handle zombie and orphan processes. Developing a working model for your own study helps to underdstand these concepts. Golang hides most of the internal complexities and provides high level API's that are easy to follow.

# Features
**minit**, at the very basic level, does the below:
* ForkExec a process that you pass to it during it's start up
* Listens to OS signals and forwards it to the spawned process
* Does *Zombie* reaping

# Usage
You can use **minit** inside a container as an init process or build a custom rootFS image with **minit**  placed under */sbin/minit* for use with **Firecracker**.

The following output's are from a local docker setup. The pre-built image is available [here](https://hub.docker.com/repository/docker/arunmudaliar/minit/general)


```# docker run arunmudaliar/minit:1.0.0```

Inspecting the process tree inside the running container gives:
```
# ps -ef
UID          PID    PPID  C STIME TTY          TIME CMD
root           1       0  0 05:12 ?        00:00:00 minit sleep 10000
root          14       1  0 05:12 ?        00:00:00 sleep 10000
```
Key observations:
* Running the above container starts **minit** as PID 1 aka init process. 
* The process *sleep 10000* was spawned by minit. 
* **minit** will now relay OS signals to this process.

Send an Interrupt to this container (kill or Ctrl-C) and it will be forwarded to *sleep*.
```
# docker run arunmudaliar/minit:1.0.0
[minit] sleep 10000
^C[minit] received  interrupt signal for PID 14
```

# References
* [go-reaper](https://github.com/ramr/go-reaper)
* [go-init](https://github.com/pablo-ruth/go-init)