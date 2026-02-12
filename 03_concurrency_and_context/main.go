package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job represents a unit of work
type Job struct {
	ID    int
	Value int
}

// Result represents the outcome of a job
type Result struct {
	JobID  int
	Output int
	Err    error
}

func main() {
	fmt.Println("=== 03 Concurrency & Context ===")

	// 1. Context with Timeout helps prevent goroutine leaks
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Channels for jobs and results
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	// 2. Worker Pool using WaitGroup
	var wg sync.WaitGroup
	numWorkers := 3

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, &wg, jobs, results)
	}

	// 3. Send jobs asynchronously
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- Job{ID: i, Value: rand.Intn(100)}
		}
		close(jobs)
	}()

	// 4. Collect results in a separate goroutine to avoid blocking
	go func() {
		wg.Wait()
		close(results)
	}()

	// 5. Process results
	for res := range results {
		if res.Err != nil {
			fmt.Printf("❌ Job %d failed: %v\n", res.JobID, res.Err)
		} else {
			fmt.Printf("✅ Job %d processed: %d\n", res.JobID, res.Output)
		}
	}

	fmt.Println("All jobs completed or timed out.")
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("⚠️ Worker %d stopping: %v\n", id, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}

			// Simulate variable processing time
			select {
			case <-time.After(time.Duration(rand.Intn(500)) * time.Millisecond):
				results <- Result{
					JobID:  job.ID,
					Output: job.Value * 2,
					Err:    nil,
				}
			case <-ctx.Done():
				results <- Result{
					JobID: job.ID,
					Err:   ctx.Err(),
				}
			}
		}
	}
}
