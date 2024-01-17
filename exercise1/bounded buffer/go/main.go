package main

import "fmt"
import "time"

func producer(ch chan int, done chan bool) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		// TODO: push real value to buffer
		ch <- i
	}
	done <- true
}

func consumer(ch chan int, done chan bool) {

	time.Sleep(1 * time.Second)
	for {
		i := <-ch
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
		if i >= 9 {
			break
		}
	}
	done <- true
}

func main() {

	// TODO: make a bounded buffer
	ch := make(chan int, 5)
	done := make(chan bool)

	go producer(ch, done)
	go consumer(ch, done)

	<-done
	<-done
}
