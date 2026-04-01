package main

// Task: Implement a thread-safe cache with lazy loading. On the first access to a key — the loader function is called.
// Subsequent accesses return the cached value.
// Concurrent requests for the same key must invoke the loader exactly once.

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type value struct {
	once sync.Once
	val  interface{}
	err  error
}

type Cache struct {
	mu     sync.RWMutex
	cache  map[string]*value
	loader func(key string) (interface{}, error)
}

func NewCache(loader func(key string) (interface{}, error)) *Cache {
	return &Cache{
		cache:  make(map[string]*value),
		loader: loader,
	}
}

// TODO: implement
// Get returns the cached value for the key, calling loader on the first access.
// Concurrent calls with the same key must invoke loader exactly once.
func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	v, ok := c.cache[key]
	c.mu.RUnlock()

	if !ok {
		c.mu.Lock()
		v, ok = c.cache[key]
		if !ok {
			v = &value{}
			c.cache[key] = v
		}
		c.mu.Unlock()
	}

	v.once.Do(func() {
		v.val, v.err = c.loader(key)
	})

	return v.val, v.err
}

func main() {
	var loadCount sync.Map // for safe per-key counting

	cache := NewCache(func(key string) (interface{}, error) {
		actual, _ := loadCount.LoadOrStore(key, new(int))
		counter := actual.(*int)
		// use atomic to avoid a race on the counter
		_ = counter
		return "value_" + key, nil
	})

	// Test 1: one key, 100 goroutines → loader called exactly once
	var count1 atomic.Int64
	cache1 := NewCache(func(key string) (interface{}, error) {
		count1.Add(1)
		return "val_" + key, nil
	})

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache1.Get("key1")
		}()
	}
	wg.Wait()
	fmt.Println("Test 1 — loader called once:", count1.Load() == 1) // true

	// Test 2: different keys → loader called for each key
	var count2 atomic.Int64
	cache2 := NewCache(func(key string) (interface{}, error) {
		count2.Add(1)
		return "val_" + key, nil
	})
	cache2.Get("key_a")
	cache2.Get("key_b")
	cache2.Get("key_a")                                                // repeated — does not call loader
	fmt.Println("Test 2 — loader called 2 times:", count2.Load() == 2) // true

	// Test 3: values are correct
	v, _ := cache.Get("key1")
	fmt.Println("Test 3 — value is correct:", v == "value_key1") // true
}
