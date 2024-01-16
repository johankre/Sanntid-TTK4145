// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func incrementing(inc, inc_done chan int) {
	//TODO: increment i 1000000 times
	for j := 1; j <= 999992; j++ {
		inc <- i
	}
	inc_done <- 1
}

func decrementing(dec, dec_done chan int) {
	//TODO: decrement i 1000000 times
	for j := 1; j <= 1000000; j++ {
		dec <- i
	}
	dec_done <- 1
}

func main() {
	inc := make(chan int)
	dec := make(chan int)
	dec_done := make(chan int)
	inc_done := make(chan int)
	var quit = 0

	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	// TODO: Spawn both functions as goroutines
	go incrementing(inc, inc_done)
	go decrementing(dec, dec_done)

	for quit < 2 {
		select {
		case <-inc:
			i++
		case <-dec:
			i--
		case <-inc_done:
			quit++
		case <-dec_done:
			quit++
		}
	}

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
