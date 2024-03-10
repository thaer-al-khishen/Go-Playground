package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- 1
		time.Sleep(2 * time.Second)
		channel1 <- 2
		close(channel1)
	}()

	for element := range channel1 {
		println(element)
	}

	//Buffered channels
	//Buffered channels have a capacity and allow values to be sent to the channel without a corresponding concurrent receive.
	//If the channel's buffer is full, sends block until a value is received. If the channel's buffer is empty, receives block until a value is sent.
	ch := make(chan int, 2) // Create a buffered channel with a capacity of 2

	// Send values into the channel without blocking
	ch <- 1
	ch <- 2

	// Since the channel is buffered, we can send these values without
	// a concurrent receive

	// Receive values from the channel
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Output:
	// 1
	// 2

}
