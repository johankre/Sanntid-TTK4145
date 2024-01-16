Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Paralellism run multiple tasks at the same time on multiple cpus
Concurrency is running multiple tasks on one cpu.

What is the difference between a *race condition* and a *data race*? 
> Data race is when two threads access the same object at the same time
Race condition is when the order of events affects the outcome.
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> Assigning resources to solve tasks.


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> Allows multiple tasks to be completed at the same time.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> *Your answer here*

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> *Your answer here*

What do you think is best - *shared variables* or *message passing*?
> *Your answer here*


