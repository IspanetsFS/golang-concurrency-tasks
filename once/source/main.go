package main

// Task: Implement a thread-safe cache with lazy loading. On the first access to a key — the loader function is called.
// Subsequent accesses return the cached value.
// Concurrent requests for the same key must invoke the loader exactly once.

import (
	"fmt"
	"sync"
)

type Cache struct {
	// TODO: add fields
}

func NewCache(loader func(key string) (interface{}, error)) *Cache {
	panic("implement me")
}

// TODO: implement
// Get returns the cached value for the key, calling loader on the first access.
// Concurrent calls with the same key must invoke loader exactly once.
func (c *Cache) Get(key string) (interface{}, error) {
	panic("implement me")
}

func main() {
	loadCount := 0
	var mu sync.Mutex

	cache := NewCache(func(key string) (interface{}, error) {
		mu.Lock()
		loadCount++
		mu.Unlock()
		return "value_" + key, nil
	})

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Get("key1")
		}()
	}
	wg.Wait()

	fmt.Println("loadCount:", loadCount) // must be 1
}
