// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"sync"
	"time"
)

var i = 0

// Mutexes for synchronization
var mutex sync.Mutex

func incrementing() {
	//TODO: increment i 1000000 times
	for j := 0; j <= 1000000; j++ {
		mutex.Lock()
		i++
		mutex.Unlock()
	}
}

func decrementing() {
	//TODO: decrement i 1000000 times
	for j := 1000000; j >= 0; j-- {
		mutex.Lock()
		i--
		mutex.Unlock()
	}

}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	// 1 gives a sequential program.
	// 2 or more gives a concurrently executing program.
	// Answer in results.md
	runtime.GOMAXPROCS(2)

	// TODO: Spawn both functions as goroutines
	go incrementing()
	go decrementing()

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
