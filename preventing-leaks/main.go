package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ReceivingNicely()
	RandStream(4)
	RandStreamCorrect(4)
}

func ReceivingNicely() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done")
}

func RandStream(n int) {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			// print never gets run, after the n-th iteration
			// the goroutine blocks trying to send the next random
			// integer to a channel that is no longer being read from,
			// we have no way to tell the producer to stop.
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()

		return randStream
	}

	randStream := newRandStream()
	fmt.Printf("%d random ints:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}

func RandStreamCorrect(n int) {

	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

	time.Sleep(1 * time.Second)
}
