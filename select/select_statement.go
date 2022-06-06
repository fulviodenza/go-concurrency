package main

import (
	"fmt"
	"time"
)

func simple_select() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(time.Second * 5)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

func more_channels_ready() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d, c2Count: %d\n", c1Count, c2Count)
}

func no_channel_ready() {
	var c <-chan int
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}

func no_channel_ready_and_something_meantime() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}

func main() {
	fmt.Println("Quick example of select statement")
	simple_select()
	fmt.Println("***************")
	fmt.Println("What happens when multiple channels have something to read?")
	more_channels_ready()
	fmt.Println("***************")
	fmt.Println("What happens when multiple channels have something to read?")
	no_channel_ready()
	fmt.Println("***************")
	fmt.Println("What happens when no channel is ready, and we need to do something in the meantime?")
	no_channel_ready_and_something_meantime()
	fmt.Println("***************")
}
