# Result: Sharing a variable
***Explain what happened, and why*** 
- Go  
The result is different every time we run the programs.
The concurrent program is ***not*** synchronized, thus the result may deviate from zero after 500 milliseconds.  

- C  


### [Go runtime package](https://pkg.go.dev/runtime)

- Contains operations that interact with Go's runtime system.
- The following variables control the run-time behavior of Go programs:
    - ***GOGC***   
    Sets the initial garbage collection target percentage.
    - ***GOMEMLIMIT***   
    Sets a soft memory limit for the runtime.
    Numeric value in bytes.
    Default setting disables the memory limit.
    - ***GODEBUG***  
    Controls debugging variables within the runtime.
    - ***GOMAXPROCS***   
    Variable that limits the number of operating system threads that can execute Go code simultaneously.
    - ***GORACE***   
    Configures the race detector.
    - ***GOTRACBACK***  
    Controls the amount of output generated when a Go program fails.
    Default prints a tack trace for the current goroutine, eliding functions internal to the runtime system.

