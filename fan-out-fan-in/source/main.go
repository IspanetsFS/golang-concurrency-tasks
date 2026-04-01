package main

// Task: Implement a function that accepts a channel of numbers, launches N workers to process them
// (squaring each value), and merges the results into a single output channel.

import (
	"fmt"
)

// TODO: implement
// fanOut distributes values from the input channel across n output channels
func fanOut(input <-chan int, n int) []<-chan int {
	panic("implement me")
}

// TODO: implement
// fanIn merges multiple input channels into a single output channel
func fanIn(channels ...(<-chan int)) <-chan int {
	panic("implement me")
}

// TODO: implement
// process reads from the input channel, squares each value, and writes to the output channel
func process(input <-chan int) <-chan int {
	panic("implement me")
}

func main() {
	input := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()

	// Split into 3 workers, process, then merge back
	outputs := fanOut(input, 3)

	processed := make([]<-chan int, len(outputs))
	for i, ch := range outputs {
		processed[i] = process(ch)
	}

	for result := range fanIn(processed...) {
		fmt.Println(result)
	}
}
