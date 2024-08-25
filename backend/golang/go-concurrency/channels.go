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

/*
----
When deadlock occurs
----
**/
// * if the channel is closed and we still sent it value why wont this throw an error?
//The reason this doesn't cause an error is that the goroutine sending values to the channel will block after sending the first two values (since the channel's buffer is full), but the program will continue executing in the main goroutine.

func BuffWorking() {
	chann := make(chan int, 2)
	// chann2 := make(chan string, 2)

	go func(chn chan<- int) {
		defer close(chn)

		// After sending the first 2 values our buffered channel is full.
		// Thus in our main go routine we receive it and call it a day
		chn <- 1
		chn <- 1

		// But what happens  when we send a new values to buffer channel after being full? why wont it cause problems
		// goroutine sending values to the channel will block after sending the first two values

		chn <- 1 // blocks from sending to buffer
		chn <- 1 // blocks from sending to buffer
		chn <- 1 // blocks from sending to buffer

		// chn <- "hello"
		// chn <- "hello"
		// chn <- "hello"

	}(chann)

	fmt.Println("value listened", <-chann) // 1
	fmt.Println("value listened", <-chann) // 1

	// THIS WONT CAUSE ERROR as our buffered channel is full, so listening to a buffered channel when its full wont cause any error in "main-go-routine", it will cause error to our "go-routine" which is sending more values to it. But we aren't listening to them in main.go. so below code doesn't panics
	// fmt.Println("value listened", <-chann) // no value
	// fmt.Println("value listened", <-chann) // no value
	// fmt.Println("value listened", <-chann) // no value

}

// UnBuffered channel will block but main go-routine wont
// if its not closed and not received immediately

/*
	Unbuffered channel: unBuff is an unbuffered channel (created with make(chan string) without a capacity). This means that sends on this channel will block until there's a corresponding receive.

Goroutine behavior: The goroutine will start executing, but it will block on the first send operation (unBuff <- "bomb created") until there's a receive operation.
Main goroutine: The main goroutine performs only one receive operation (<-unBuff). This receive will unblock the first send in the other goroutine.
Program termination: After printing the received value, the main function ends, which terminates the program. This happens before the goroutine has a chance to send the remaining 9 values.

Here's the sequence of events:

The goroutine starts and attempts to send "bomb created" to the channel.
The send operation blocks because there's no receiver yet.
The main goroutine reaches the receive operation and takes the first "bomb created" from the channel.
The main goroutine prints "work is done bomb created" and then exits.
The program terminates before the goroutine can send any more values.

The reason this doesn't throw an error is that Go doesn't consider it an error for goroutines to be left hanging when the program exits. The runtime simply terminates all running goroutines when the main function returns.
*/
func UnBuff() {

	unBuff := make(chan string)
	go func(unBuff chan<- string) {
		defer close(unBuff)

		for i := 0; i < 10; i++ {
			unBuff <- "bomb created"
		}

	}(unBuff)

	fmt.Println("work is done", <-unBuff)

}

// for val := range unBuff {
// 	fmt.Println("work is done", val)
// }
