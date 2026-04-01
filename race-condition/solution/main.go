package main

// Task: The code below contains a race condition.
// Find the problem and fix it in three ways: using sync.Mutex, using sync/atomic, and using a channel.

import (
	"fmt"
	"sync"
	"sync/atomic"
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
	m     sync.Mutex
	value int
}

func (c *MutexCounter) Increment() {
	c.m.Lock()
	defer c.m.Unlock()
	c.value++
}

func (c *MutexCounter) Value() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.value
}

// TODO: fix using sync/atomic
type AtomicCounter struct {
	value atomic.Int64
}

func (c *AtomicCounter) Increment() {
	c.value.Add(1)
}

func (c *AtomicCounter) Value() int {
	return int(c.value.Load())
}

// TODO: fix using a channel
type ChanCounter struct {
	inc  chan struct{}
	get  chan int
	done chan struct{}
}

func NewChanCounter() *ChanCounter {
	c := &ChanCounter{
		inc:  make(chan struct{}),
		get:  make(chan int),
		done: make(chan struct{}),
	}
	go func() {
		var value int
		for {
			select {
			case <-c.done:
				return
			case <-c.inc:
				value++
			case c.get <- value:
			}
		}
	}()
	return c
}

func (c *ChanCounter) Increment() {
	select {
	case c.inc <- struct{}{}:
	case <-c.done:
	}
}

func (c *ChanCounter) Value() int {
	select {
	case val := <-c.get:
		return val
	case <-c.done:
		return -1
	}
}

func (c *ChanCounter) Stop() {
	close(c.done)
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
	// c := &UnsafeCounter{}
	// runTest(c.Increment, c.Value)

	mc := &MutexCounter{}
	runTest(mc.Increment, mc.Value)

	ac := &AtomicCounter{}
	runTest(ac.Increment, ac.Value)

	cc := NewChanCounter()
	runTest(cc.Increment, cc.Value)
	cc.Stop()
	cc.Increment()
	println(cc.Value())
}
