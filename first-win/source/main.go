package main

// Task: Launch N goroutines, each performing a request.
// Return the result of the first goroutine that succeeds; cancel all others via context.

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// TODO: implement
// firstWin launches len(funcs) goroutines concurrently.
// Returns the result of the first one to succeed.
// All other goroutines receive a cancelled ctx and must exit.
// If ALL goroutines return an error — returns the last error.
func firstWin(ctx context.Context, funcs []func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	panic("implement me")
}

func main() {
	// Simulate multiple servers — take the response from the fastest one
	servers := []func(context.Context) (interface{}, error){
		makeRequest("server1", 300*time.Millisecond, false),
		makeRequest("server2", 100*time.Millisecond, false),
		makeRequest("server3", 200*time.Millisecond, false),
	}

	result, err := firstWin(context.Background(), servers)
	fmt.Println(result, err) // "server2", nil — the fastest one

	// All fail with an error
	failing := []func(context.Context) (interface{}, error){
		makeRequest("s1", 50*time.Millisecond, true),
		makeRequest("s2", 100*time.Millisecond, true),
	}
	result, err = firstWin(context.Background(), failing)
	fmt.Println(result, err) // nil, error
}

func makeRequest(name string, delay time.Duration, fail bool) func(context.Context) (interface{}, error) {
	return func(ctx context.Context) (interface{}, error) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(delay):
		}

		if fail {
			return nil, errors.New(name + " failed")
		}
		return name, nil
	}
}
