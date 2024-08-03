package main

import (
	"fmt"
	"time"
)

// Buffered and UnBuffered channels
// Unbuffered Channels: Communication happens only when both sender and receiver are ready. They are useful for synchronization.
// Buffered Channels: Allow sending and receiving to proceed independently up to the buffer's capacity. They are useful for decoupling the timing of sending and receiving.

func BuffChan() {
	channel := make(chan int, 2) // Buffered channel with capacity 2

	channel <- 42 // This will not block
	channel <- 43 // This will not block

	go func() {
		channel <- 44 // This will block because the buffer is full
	}()

	value1 := <-channel // This will unblock the goroutine
	value2 := <-channel
	value3 := <-channel

	fmt.Println("Received:", value1, value2, value3)
}

// unbuff channels is supposed to be received/send immediately

func UnBufferedChannel() {
	channel := make(chan int) // Unbuffered channel

	go func() {
		channel <- 42 // This will block until the value is received
	}()

	value := <-channel // This will block until a value is sent
	fmt.Println("Received:", value)
}

// Close Channel Example
func CloseChannel() {
	channel := make(chan string, 5) // buffered bidirectional channel

	go func() {
		for i := 0; i < 5; i++ {
			channel <- "money"
		}

		fmt.Println("length of channel is  ", len(channel))

		// Sender should always close channel
		close(channel)
	}()

	for {
		// Check if channel is closed
		value, ok := <-channel

		if !ok {
			fmt.Println("Channel is emptied")
			break
		}

		fmt.Println("Incoming", value)
	}

}

// Bidirectional channels and unidirectional channels
// by default all channels are bidirectional
// to make a channel unidirectional, you have to change the direction of channel in function signature

func BiAndUniChannel() {

	// bidirectional UnBuffered channel
	channel := make(chan string)

	go sendOnlyChan(channel, "50 apples")
	receiveOnlyChan(channel)
}

func sendOnlyChan(chn chan<- string, message string) {
	chn <- message

	fmt.Println("Value sent")
}

func receiveOnlyChan(chn <-chan string) {
	fmt.Println("Received value of  ", <-chn)
}

/*
-----
SELECT with channels
----
*/

//* Handling Multiple Channels

func SelectWithChannels() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "message from chan1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		chan2 <- "message from chan2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		case <-time.After(4 * time.Second):
			fmt.Println("Timeout")
		}

	}
}

func TimeoutChannelWithSelect() {
	channel := make(chan string)

	select {
	case msg := <-channel:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received")
	}
}
