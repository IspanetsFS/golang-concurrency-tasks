package main

// Task: Implement a Circuit Breaker — a pattern that protects against cascading failures.
// If threshold consecutive errors occur — "open the circuit" and immediately return an error without calling fn.
// After resetTimeout elapses, attempt a probe call (half-open state).

import (
	"errors"
	"fmt"
	"time"
)

type State int

const (
	Closed   State = iota // operating normally
	Open                  // circuit is open, requests are rejected
	HalfOpen              // probe call allowed to test recovery
)

var ErrCircuitOpen = errors.New("circuit breaker: circuit is open")

type CircuitBreaker struct {
	// TODO: add fields
}

// TODO: implement
func NewCircuitBreaker(threshold int, resetTimeout time.Duration) *CircuitBreaker {
	panic("implement me")
}

// TODO: implement
// Do executes fn if the circuit is closed or half-open.
// Returns ErrCircuitOpen immediately if the circuit is open.
func (cb *CircuitBreaker) Do(fn func() error) error {
	panic("implement me")
}

func main() {
	cb := NewCircuitBreaker(3, 100*time.Millisecond)

	// Simulate failures
	for i := 0; i < 5; i++ {
		err := cb.Do(func() error {
			return errors.New("service unavailable")
		})
		fmt.Printf("attempt %d: %v\n", i+1, err)
		// 1,2,3: "service unavailable"
		// 4,5:   "circuit breaker: circuit is open"
	}

	// Wait for reset timeout
	time.Sleep(150 * time.Millisecond)

	// Probe call (HalfOpen state)
	err := cb.Do(func() error { return nil })
	fmt.Println("after reset:", err) // nil — circuit is closed again
}
