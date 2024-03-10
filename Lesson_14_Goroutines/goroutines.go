package main

import (
	"fmt"
	"time"
)

func main() {
	go say("world")
	say("hello")

	//You use channels to communicate between goroutines

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) //You're creating a channel of integers
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	//You can also create a buffered channel like this:
	bufferedC := make(chan int, 100)
	go func() {
		time.Sleep(200)
		bufferedC <- 2
		bufferedC <- 3
		bufferedC <- 5
	}()

	bufferedResult := <-bufferedC
	fmt.Println(bufferedResult)

	//You can close a channel like this
	close(bufferedC)

	//You can check if a channel is closed like this
	_, ok := <-bufferedC //The first value _ is the value stored in the channel
	fmt.Println(ok)      //Should return false if the channel is empty or closed

	channel := make(chan int)

	go writeCh(channel)
	readCh(channel)

	//A select statement is unique to channels, and invokes the code of the channel that receives a signal first

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	//Perform an operation before sending a sum to a channel
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// Write-only channel
func writeCh(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

// Read-only channel
func readCh(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
