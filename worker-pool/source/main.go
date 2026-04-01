package main

// Task: Implement a worker pool. There are N jobs to process using at most M concurrent workers.
// Each worker simulates work via time.Sleep. Results must be returned in the same order as the input jobs.

import (
	"fmt"
	"time"
)

type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobID  int
	Output int
}

// TODO: implement
// workerPool runs at most maxWorkers goroutines concurrently.
// Returns results IN THE SAME ORDER as the input jobs.
func workerPool(jobs []Job, maxWorkers int) []Result {
	panic("implement me")
}

func processJob(job Job) Result {
	time.Sleep(10 * time.Millisecond) // simulate work
	return Result{
		JobID:  job.ID,
		Output: job.Value * 2,
	}
}

func main() {
	jobs := make([]Job, 10)
	for i := range jobs {
		jobs[i] = Job{ID: i, Value: i + 1}
	}

	results := workerPool(jobs, 3)
	for _, r := range results {
		fmt.Printf("Job %d: %d\n", r.JobID, r.Output)
	}
}
