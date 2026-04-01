package main

// Task: Implement a rate limiter that restricts the number of requests to N per second.
// Use a ticker. The Wait function must block until a token is available.

import (
	"context"
	"fmt"
	"time"
)

type RateLimiter struct {
	// TODO: add fields
}

// TODO: implement
func NewRateLimiter(rps int) *RateLimiter {
	panic("implement me")
}

// TODO: implement
// Wait blocks until a token is available or ctx is cancelled.
// Returns ctx.Err() if the context is cancelled.
func (r *RateLimiter) Wait(ctx context.Context) error {
	panic("implement me")
}

// TODO: implement
// Stop releases the underlying ticker and signals the limiter to shut down.
func (r *RateLimiter) Stop() {
	panic("implement me")
}

func main() {
	limiter := NewRateLimiter(5) // 5 запросов в секунду
	defer limiter.Stop()

	ctx := context.Background()
	for i := 0; i < 10; i++ {
		if err := limiter.Wait(ctx); err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Printf("request %d at %v\n", i, time.Now().Format("15:04:05.000"))
	}
}
