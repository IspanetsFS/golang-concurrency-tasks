package main

// Task: Implement a three-stage pipeline: generate → square → filter.
// The pipeline must shut down cleanly when the context is cancelled (no goroutine leaks).

import (
	"context"
	"fmt"
)

// TODO: generates numbers from 1 to n, closes the channel on ctx.Done()
func generate(ctx context.Context, n int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for i := range n {
			select {
			case <-ctx.Done():
				return
			case result <- i + 1:
			}
		}
	}()
	return result
}

// TODO: squares each number from the input channel
func square(ctx context.Context, in <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for val := range in {
			select {
			case <-ctx.Done():
				return
			case result <- val * val:
			}
		}
	}()
	return result
}

// TODO: passes only numbers greater than threshold
func filter(ctx context.Context, in <-chan int, threshold int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for val := range in {
			if val <= threshold {
				continue
			}
			select {
			case <-ctx.Done():
				return
			case result <- val:
			}
		}
	}()
	return result
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nums := generate(ctx, 20)
	squared := square(ctx, nums)
	filtered := filter(ctx, squared, 50)
	var i int
	for val := range filtered { // range over the channel
		fmt.Printf("[%d] %d\n", i, val)
		i++
		if i == 4 { // take only the first 5 results
			cancel() // cancel the pipeline
			break
		}
	}

	// After cancel() — no goroutines should remain
}
