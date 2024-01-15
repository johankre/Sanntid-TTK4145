// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	//"time"
)

var i = 0

func incrementing(incr chan int) {
	for j := 0; j <= 1000000; j++ {
		i += 1
		incr <- i
	}
}

func decrementing(decr chan int) {
	for j := 0; j <= 1000001; j++ {
		i -= 1
		decr <- i
	}

}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	// TODO: Spawn both functions as goroutines
	incr := make(chan int)
	decr := make(chan int)

	go incrementing(incr)
	go decrementing(decr)

	for {
		select {
		case <-incr:
			i := <-incr
		case <-decr:
			i := <-decr
		}
	}

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	// time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
