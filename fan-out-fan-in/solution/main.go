package main

// Task: Implement a function that accepts a channel of numbers, launches N workers
// to process them (squaring each value), and merges the results into a single output channel.

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

// TODO: implement
// fanOut distributes values from the input channel across n output channels
func fanOut(input <-chan int, n int) []<-chan int {
	chans := make([]<-chan int, 0, n)
	for range n {
		ch := make(chan int)
		chans = append(chans, ch)

		go func() {
			defer close(ch)
			for i := range input {
				ch <- i
			}
		}()
	}
	return chans
}

// TODO: implement
// fanIn merges multiple input channels into a single output channel
func fanIn(channels ...(<-chan int)) <-chan int {
	result := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for i := range ch {
				result <- i
			}
		}()

	}
	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}

// TODO: implement
// process reads from the input channel, squares each value, and writes to the output channel
func process(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for i := range input {
			output <- i * i
		}
	}()
	return output
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

	// Collect all results and verify:
	results := []int{}
	for result := range fanIn(processed...) {
		results = append(results, result)
	}

	sort.Ints(results)
	expected := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	fmt.Println(reflect.DeepEqual(results, expected)) // true
}
