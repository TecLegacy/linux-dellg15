package main

import (
	"fmt"
	"time"
)

func GreetMorning() {
	fmt.Println("Good Morning")
}

func ExecuteGreetMorning() {
	go GreetMorning()
	time.Sleep(1 * time.Microsecond)
	fmt.Print("Learning concurrency")

	// you have to wait for go routine to finish its execution
	// so you have to wait till the time that function is completed executing its task, it could be for 1 seconds for 1 hr
	// in this case 1 microseconds is miniscule for greetMorning() to complete its execution thus "Learning concurrency gets printed first"
	/**In your current code, the main function starts a goroutine to execute GreetMorning and then sleeps for 1 microsecond. However, 1 microsecond is likely too short for the goroutine to complete its execution, so "Learning concurrency" might be printed before "Good Morning".

	To ensure that the goroutine has enough time to complete, you should sleep for a longer duration. However, using time.Sleep is not a reliable way to wait for goroutines to finish. Instead, you should use synchronization mechanisms like sync.WaitGroup.
	*/
}

/*
------------------------------
//*working with channels
*/

// When you declare a channel using var, it creates a channel variable with a nil value. This means the channel is not initialized and does not point to any actual channel in memory. On the other hand, when you use make to create a channel, it initializes the channel and allocates memory for it, returning a reference to the newly created channel.
func ExampleChannel() {
	var channel chan int

	// <nil>
	fmt.Println("value of channel first", channel)

	// chan int
	fmt.Printf("type of channel second %T\n", channel)

	channel2 := make(chan int)

	// 0xc0000281220
	fmt.Println("value of channel second", channel2)

	// chan int
	fmt.Printf("type of channel  second %T \n", channel2)

}

// The idea of channels is send and receive values
// With go-routines
func ChanAndGoRoutine() {
	// make keyword creates and initializes
	// in this case it creates a channel and initializes with a memory address
	channel := make(chan int)
	go multiplyByTen(channel)
	// sending value to channel that is supposed to be multiplied
	channel <- 2

	// DEADLOCK example - Read notion for more explanation
	// channel <- 2
	//go multiplyByTen(channel)

}

func multiplyByTen(chn chan int) {
	fmt.Printf("Multiplied by 10 value = \t  %d \n", 10*<-chn)
}
