package main

// Task: Implement a retry function that retries calling fn up to maxAttempts times with exponential backoff.
// The entire process is bounded by the context. On success — returns immediately.

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type RetryConfig struct {
	MaxAttempts int
	InitialWait time.Duration // delay before the 2nd attempt
	Factor      float64       // multiplier: wait *= factor on each iteration
}

// TODO: implement
// Retry calls fn up to cfg.MaxAttempts times.
// Between attempts it waits InitialWait * Factor^attempt.
// Stops early if ctx is cancelled.
// Returns the last error if all attempts are exhausted.
func Retry(ctx context.Context, cfg RetryConfig, fn func(ctx context.Context) error) error {
	waitTimeout := cfg.InitialWait
	var lastErr error
	for attempt := 0; attempt < cfg.MaxAttempts; attempt++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		lastErr = fn(ctx)
		if lastErr == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(waitTimeout):
		}
		waitTimeout *= time.Duration(cfg.Factor)
	}
	return lastErr
}

func main() {
	attempts := 0
	err := Retry(context.Background(), RetryConfig{
		MaxAttempts: 5,
		InitialWait: 10 * time.Millisecond,
		Factor:      2.0,
	}, func(ctx context.Context) error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary error")
		}
		return nil // success on the 3rd attempt
	})

	fmt.Println(err)      // nil
	fmt.Println(attempts) // 3
}
