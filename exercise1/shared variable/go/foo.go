// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	//"time"
)

var i = 0

func incrementing(incr, quit chan int) {
	for j := 0; j <= 1000000; j++ {
		// i += 1
		incr <- 0
	}
	quit <- 0
}

func decrementing(decr, quit chan int) {
	for j := 0; j <= 1000001; j++ {
		// i -= 1
		decr <- 0
	}
	quit <- 0

}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)
	n_quit := 0

	// TODO: Spawn both functions as goroutines
	incr := make(chan int)
	decr := make(chan int)
	quit := make(chan int)

	go incrementing(incr, quit)
	go decrementing(decr, quit)

	for n_quit < 2 {
		select {
		case <-incr:
			i++
		case <-decr:
			i--
		case <-quit:
			n_quit++
		}
	}

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	// time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
