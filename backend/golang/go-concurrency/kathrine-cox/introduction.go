package kathrinecox

import (
	"fmt"
	"sync"
	"time"
)

// Why is Concurrency hard
/*
1. race condition -> data race
*/

func race_data() {

	var data int
	go func() { data++ }()
	time.Sleep(1 * time.Second) // This is bad!
	if data == 0 {

		fmt.Printf("the value is %v.\n", data)
	}

	// Memory synchronization
	func() {
		var memoryAccess sync.Mutex
		var value int
		go func() {
			memoryAccess.Lock()
			value++
			memoryAccess.Unlock()
		}()
		memoryAccess.Lock()
		if value == 0 {
			fmt.Printf("the value is %v.\n", value)

		} else {
			fmt.Printf("the value is %v.\n", value)
		}
		memoryAccess.Unlock()
	}()
}

// A Coffman condition
func deadlock() {

	type value struct {
		mu    sync.Mutex
		value int
	}

	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()
		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}
	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

func K_output() {
	// race_data()
	deadlock()
}
