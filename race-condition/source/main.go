package main

// Task: The code below contains a race condition.
// Find the problem and fix it in three ways: using sync.Mutex, using sync/atomic, and using a channel.

import (
	"fmt"
	"sync"
)

// BUG: there is a data race here!
type UnsafeCounter struct {
	value int
}

func (c *UnsafeCounter) Increment() {
	c.value++ // NOT atomic!
}

func (c *UnsafeCounter) Value() int {
	return c.value
}

// TODO: fix using sync.Mutex
type MutexCounter struct {
	// ...
}

// TODO: fix using sync/atomic
type AtomicCounter struct {
	// ...
}

// TODO: fix using a channel
type ChanCounter struct {
	// ...
}

func runTest(increment func(), getValue func() int) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("Result:", getValue(), "(expected 1000)")
}

func main() {
	// This is BROKEN:
	c := &UnsafeCounter{}
	runTest(c.Increment, c.Value)

	// TODO: test your implementations
}
