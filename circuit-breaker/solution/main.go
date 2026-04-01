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

type reqState struct {
	resp chan bool
}

type CircuitBreaker struct {
	failedCallCh  chan struct{}
	successCallCh chan struct{}
	resetCh       chan struct{}
	allowCh       chan reqState
	state         State
	resetTimeout  time.Duration
	threshold     int
}

// TODO: implement
func NewCircuitBreaker(threshold int, resetTimeout time.Duration) *CircuitBreaker {
	cb := &CircuitBreaker{
		threshold:     threshold,
		resetTimeout:  resetTimeout,
		failedCallCh:  make(chan struct{}),
		successCallCh: make(chan struct{}),
		resetCh:       make(chan struct{}),
		allowCh:       make(chan reqState),
	}
	go cb.serve()
	return cb
}

func (cb *CircuitBreaker) Allow() bool {
	req := reqState{make(chan bool, 1)}
	cb.allowCh <- req
	return <-req.resp
}

func (cb *CircuitBreaker) FailedCall() {
	cb.failedCallCh <- struct{}{}
}

func (cb *CircuitBreaker) SuccessCall() {
	cb.successCallCh <- struct{}{}
}

func (cb *CircuitBreaker) serve() {
	var errCount int
	var halfOpenInFlight bool
	for {
		select {
		case <-cb.failedCallCh:
			errCount++
			halfOpenInFlight = false
			if cb.state == HalfOpen || errCount >= cb.threshold {
				errCount = 0
				cb.state = Open
				time.AfterFunc(cb.resetTimeout, func() {
					cb.resetCh <- struct{}{}
				})
			}
		case <-cb.successCallCh:
			errCount = 0
			halfOpenInFlight = false
			if cb.state == HalfOpen {
				cb.state = Closed
			}
		case <-cb.resetCh:
			errCount = 0
			cb.state = HalfOpen
		case req := <-cb.allowCh:
			switch cb.state {
			case Open:
				req.resp <- false
			case Closed:
				req.resp <- true
			case HalfOpen:
				if !halfOpenInFlight {
					halfOpenInFlight = true
					req.resp <- true
				} else {
					req.resp <- false
				}
			}
		}
	}
}

// TODO: implement
// Do executes fn if the circuit is closed or half-open.
// Returns ErrCircuitOpen immediately if the circuit is open.
func (cb *CircuitBreaker) Do(fn func() error) error {
	if !cb.Allow() {
		return ErrCircuitOpen
	}
	err := fn()
	if err != nil {
		cb.FailedCall()
	} else {
		cb.SuccessCall()
	}
	return err
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
