# Result: Sharing a variable
***Explain what happened, and why*** 
- Go  
The result is different every time we run the programs.
The concurrent program is ***not*** synchronized, thus the result may deviate from zero after 500 milliseconds.  

- C  
Creating two threads: one for increasing *i* and one for decreasing *i*. 
The program lets both threads finish, thus the result is always 0.

### C - Passing Arguments to Threads
- pthread_create()
    - All arguments mus be passed by reference and cast to (void *)
    - Only one argument to the thread start routine
    - For multiple arguments
        - Creating a structure that contains all of the arguments
        - Pass a pointer to that structure in pthread_create()

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

