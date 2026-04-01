package main

// Task: Implement a semaphore using channels.
// The semaphore must limit the number of concurrently executing operations.

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type token struct{}

type Semaphore struct {
	sem chan token
}

// TODO: implement
func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		sem: make(chan token, n),
	}
}

// TODO: implement
// Acquire acquires a semaphore slot, blocking if all slots are taken.
// Returns an error if ctx is cancelled.
func (s *Semaphore) Acquire(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.sem <- token{}:
		return nil
	}
}

// TODO: implement
// Release releases a semaphore slot.
func (s *Semaphore) Release() {
	<-s.sem
}

func main() {
	sem := NewSemaphore(3)
	var wg sync.WaitGroup
	var mu sync.Mutex
	maxConcurrent := 0
	current := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem.Acquire(context.Background())
			defer sem.Release()

			mu.Lock()
			current++
			if current > maxConcurrent {
				maxConcurrent = current
			}
			mu.Unlock()

			time.Sleep(10 * time.Millisecond)

			mu.Lock()
			current--
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("Max concurrent:", maxConcurrent) // must be <= 3
}
