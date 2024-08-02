package main

import "fmt"

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
