package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func initialize() {
	fmt.Println("Initializing...")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	once.Do(initialize)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

}

func WaitGroupAndSyncOnce() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {

		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers completed")
}

// SYNCHRONIZATION with channels
// order is preserved, data communication & synchronization between go routines
// channels are not light weight compared to wait-groups
func SyncChannels() {
	channels := make(chan bool) // bidirectional unBuffered channels

	for i := 0; i < 5; i++ {
		go LoadTruck(i, channels)
		<-channels
	}

	//? this creates race condition
	// <-channels

}

func LoadTruck(id int, chn chan bool) {
	fmt.Println("Starting my work id  ", id)

	time.Sleep(1 * time.Second) // imitate work logic

	fmt.Printf("Work done by %d \n", id)

	chn <- true
}
