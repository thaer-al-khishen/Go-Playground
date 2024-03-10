package main

import (
	"fmt"
	"sync"
)

// A shared counter variable
var (
	counter int
	mutex   sync.Mutex
	//rwmutex   sync.RWMutex
)

func main() {
	var wg sync.WaitGroup

	// Start 5 goroutines that increment the counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait() // Wait for all goroutines to finish
}

func increment() {
	mutex.Lock()         // Before accessing the counter, lock the mutex
	defer mutex.Unlock() // Ensure the mutex gets unlocked when the function exits
	counter++            // Critical section that modifies the shared variable
	fmt.Println(counter)
}
