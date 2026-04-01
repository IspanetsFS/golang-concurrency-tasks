package main

// Task: Implement a service that launches several background workers.
// When the context is cancelled — all workers must finish their current task gracefully and stop.
// The main goroutine must wait for all workers to complete before returning.

import (
	"context"
	"fmt"
	"time"
)

// TODO: implement
// runWorkers launches n workers; each worker:
// - processes a task (prints a message) every tickInterval
// - exits when ctx is cancelled
// The function blocks until ALL workers have finished.
func runWorkers(ctx context.Context, n int, tickInterval time.Duration) {
	panic("implement me")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	runWorkers(ctx, 3, 30*time.Millisecond)
	fmt.Println("all workers finished")
}
