package workerpool

import (
	"context"
	"sync"
)

const WorkerPoolSize = 100
const ChannelCapacity = 1000

// Worker pool driver
func Run[J any, R any](ctx context.Context, jobsCh <-chan J, processFunc func(job J) R) <-chan R {
	// 1. main provides monitor.go -> CheckResources the resources list to monitor
	// 2. CheckResources produces the jobs i.e. owns the jobs channel -> calls workerpool, passing context, jobs channel and a process function
	// 3. process function simply takes a resource, do health check and return the result - no channels
	// 4. worker pool takes the jobs channel, ranges through it, calls process fucntion, passing each job
	// 5. processor function processes, return the result. the worker pool writes it to the result channel that the worker pool owns and return the channel
	var wg sync.WaitGroup
	resultCh := make(chan R, ChannelCapacity)

	for i := 1; i <= WorkerPoolSize; i++ {
		wg.Go(func() {
			for job := range jobsCh {
				status := processFunc(job)
				resultCh <- status
			}
		})
	}

	// wait for all routines to finish writing, then close the channel
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}
