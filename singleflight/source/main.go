package main

// Task: Implement singleflight — if multiple goroutines concurrently request the same data,
// the actual request is executed exactly once; the rest wait and receive the same result.
// Unlike Cache: singleflight does not persist the result after completion — it only deduplicates in-flight requests.

import (
	"fmt"
	"sync"
	"time"
)

type call struct {
	// TODO: add fields
}

type SingleFlight struct {
	// TODO: add fields
}

// TODO: implement
// Do executes fn if there is no active call for the given key.
// If one is already in flight — waits for it to finish and returns the same result.
func (sf *SingleFlight) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	panic("implement me")
}

// simulate — simulates a slow request
func simulate(key string) (interface{}, error) {
	time.Sleep(100 * time.Millisecond)
	return "result_" + key, nil
}

func main() {
	sf := &SingleFlight{}
	var wg sync.WaitGroup

	callCount := 0
	var mu sync.Mutex

	// 10 goroutines request the same key concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val, err := sf.Do("key1", func() (interface{}, error) {
				mu.Lock()
				callCount++
				mu.Unlock()
				return simulate("key1")
			})
			fmt.Println(val, err)
		}()
	}

	wg.Wait()
	fmt.Println("Actual calls:", callCount) // must be 1
}
