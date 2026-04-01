package main

// Task: Implement singleflight — if multiple goroutines concurrently request the same data,
// the actual request is executed exactly once; the rest wait and receive the same result.
// Unlike Cache: singleflight does not persist the result after completion — it only deduplicates in-flight requests.

import (
	"fmt"
	"sync"
	"time"
)

type response struct {
	data interface{}
	err  error
}

type call struct {
	res    *response
	finish chan struct{}
}

func newCall() *call {
	return &call{
		finish: make(chan struct{}),
	}
}

func (c *call) do(fn func() (interface{}, error)) {
	data, err := fn()
	c.res = &response{data, err}
	close(c.finish)
}

func (c *call) wait() <-chan struct{} {
	return c.finish
}

func (c *call) result() (interface{}, error) {
	return c.res.data, c.res.err
}

type SingleFlight struct {
	mu    sync.Mutex
	calls map[string]*call
}

func NewSingleFlight() *SingleFlight {
	return &SingleFlight{
		calls: make(map[string]*call),
	}
}

// TODO: implement
// Do executes fn if there is no active call for the given key.
// If one is already in flight — waits for it to finish and returns the same result.
func (sf *SingleFlight) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	var c *call
	sf.mu.Lock()
	c, ok := sf.calls[key]
	if ok {
		sf.mu.Unlock()
		<-c.wait()
		return c.result()
	}
	c = newCall()
	sf.calls[key] = c
	sf.mu.Unlock()
	c.do(fn)
	sf.mu.Lock()
	delete(sf.calls, key)
	sf.mu.Unlock()
	return c.result()
}

// simulate — simulates a slow request
func simulate(key string) (interface{}, error) {
	time.Sleep(100 * time.Millisecond)
	return "result_" + key, nil
}

func main() {
	sf := NewSingleFlight()
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
