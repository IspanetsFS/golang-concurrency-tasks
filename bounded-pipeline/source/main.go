package main

// Task: Implement a pausable generator. The generator produces numbers, but if the downstream
// consumer cannot keep up, the generator must pause (backpressure) instead of buffering or dropping data.

import (
	"context"
	"fmt"
	"time"
)

// TODO: implement
// generator produces numbers from 1 to n
// pause/resume channels control the pause state:
//   - a signal on pause  — stops sending
//   - a signal on resume — resumes sending
//
// use the nil channel pattern for pausing
func generator(ctx context.Context, n int, pause, resume <-chan struct{}) <-chan int {
	panic("implement me")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pause := make(chan struct{}, 1)
	resume := make(chan struct{}, 1)

	nums := generator(ctx, 20, pause, resume)

	for i := 0; i < 5; i++ {
		fmt.Println("received:", <-nums)
	}

	pause <- struct{}{} // pause the generator
	fmt.Println("generator paused")
	time.Sleep(50 * time.Millisecond)

	resume <- struct{}{} // resume the generator
	fmt.Println("generator resumed")

	for v := range nums {
		fmt.Println("received:", v)
	}
}
