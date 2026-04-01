package main

// Task: Implement a rate limiter that restricts the number of requests to N per second.
// Use a ticker. The Wait function must block until a token is available.

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrStopped = errors.New("rate limiter: stopped")

type RateLimiter struct {
	ticker *time.Ticker
	done   chan struct{}
	once   sync.Once
}

func NewRateLimiter(rps int) *RateLimiter {
	return &RateLimiter{
		ticker: time.NewTicker(time.Second / time.Duration(rps)),
		done:   make(chan struct{}),
	}
}

func (r *RateLimiter) Wait(ctx context.Context) error {
	select {
	case <-r.ticker.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-r.done:
		return ErrStopped
	}
}

func (r *RateLimiter) Stop() {
	r.once.Do(func() {
		close(r.done)
		r.ticker.Stop()
	})
}

func main() {
	limiter := NewRateLimiter(10)
	defer limiter.Stop()

	ctx := context.Background()
	start := time.Now()
	for i := 0; i < 10; i++ {
		if err := limiter.Wait(ctx); err != nil {
			fmt.Println("error:", err)
			return
		}
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed >= 900*time.Millisecond) // true
	fmt.Println(elapsed < 1200*time.Millisecond) // true

	ctx2, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	limiter2 := NewRateLimiter(1)
	defer limiter2.Stop()
	limiter2.Wait(ctx2)
	err := limiter2.Wait(ctx2)
	fmt.Println(errors.Is(err, context.DeadlineExceeded)) // true
}
