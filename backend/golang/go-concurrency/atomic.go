package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// causes race condition
// outputs any value the completes is execution

/*
Imagine a busy ice cream shop with multiple servers. There's a single whiteboard where they keep track of how many ice creams they've sold. When a server sells an ice cream, they need to update this count.
Problem without atomic operations:

Server A reads the count: 100
Server B reads the count: 100
Server A adds 1 and writes 101
Server B adds 1 and writes 101
*/

func __without_atomic() {
	count := 0
	var wg sync.WaitGroup

	increment := func() {
		defer wg.Done()
		current := count // Read
		current++        // Modify
		count = current  // Write
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Println("Final count:", count)
}

// consistently outputs 1000
// no race conditions
func AtomicIncrement() {
	var count int64 = 0
	var wg sync.WaitGroup

	increment := func() {
		defer wg.Done()
		atomic.AddInt64(&count, 1) // Atomic increment
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Println("Final count:", atomic.LoadInt64(&count))
}
