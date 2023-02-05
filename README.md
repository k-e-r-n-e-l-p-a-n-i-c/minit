# minit
A minimalistic init system written in Golang for educational purpose.

# Motivation
First of all it's *fun*. You will have to do a fair amount of study about production grade init system's to understand the internal working's on UNIX/Linux OS. You get to know about how signals and traps work and how to handle zombie and orphan processes. Developing a working model for your own study helps to underdstand these concepts. 

# Features
**minit**, at the very basic level, does the below:
* ForkExec a process that you pass to it during it's start up
* Listens to OS signals and forwards it to the spawned process
* Does *Zombie* reaping

# Usage
You can use **minit** inside a container as an init process or build a custom rootFS image with **minit**  placed under */sbin/minit* for use with **Firecracker**.

The following output's are from a local docker setup. The pre-built image is available [here](https://hub.docker.com/repository/docker/arunmudaliar/minit/general)



# References
* [go-reaper](https://github.com/ramr/go-reaper)
* [go-init](https://github.com/pablo-ruth/go-init)