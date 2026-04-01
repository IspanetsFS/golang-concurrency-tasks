package main

// Task: Implement a merge function that combines two channels into one.
// When one of the input channels is closed — stop reading from it, but continue reading the other.
// When both are closed — close the output channel.
// Do NOT use sync.WaitGroup — use only the nil channel pattern.

import "fmt"

// TODO: implement
// merge combines two channels into one.
// When one channel is closed — continues reading the other.
// When both are closed — closes the output channel.
// Do NOT use sync.WaitGroup — use only the nil channel pattern.
func merge(a, b <-chan int) <-chan int {
	panic("implement me")
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		for _, v := range []int{1, 2, 3} {
			a <- v
		}
		close(a) // a closes first
	}()

	go func() {
		for _, v := range []int{4, 5, 6, 7, 8} {
			b <- v
		}
		close(b)
	}()

	for v := range merge(a, b) {
		fmt.Println(v)
	}
	// all 8 numbers must be received
}
