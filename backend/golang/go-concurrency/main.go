package main

import "fmt"

func main() {

	// ExecuteGreetMorning()
	// ExampleChannel()
	// ChanAndGoRoutine()
	// BuffChan()
	// UnBufferedChannel()
	// CloseChannel()
	// BiAndUniChannel()
	// BuffWorking()
	// WaitGroupAndSyncOnce()
	// SyncChannels()
	// SelectWithChannels()
	// TimeoutChannelWithSelect()
	AtomicIncrement()
}

// * if the channel is closed and we still receive it wont throw error
func BuffWorking() {
	chann := make(chan int, 2)
	// chann2 := make(chan string, 2)

	go func(chn chan<- int) {

		chn <- 1
		chn <- 1
		chn <- 1
		// chn <- "hello"
		// chn <- "hello"
		// chn <- "hello"
		close(chn)
	}(chann)

	fmt.Println("value listened", <-chann)
	fmt.Println("value listened", <-chann)
	fmt.Println("value listened", <-chann)
	fmt.Println("value listened", <-chann) // 0
	fmt.Println("value listened", <-chann) // 0

}
