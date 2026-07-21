package monitor

import (
	"context"
	"day09/health"
	"day09/metrics"
	"day09/workerpool"
	"time"
)

type Monitor struct{}

const WorkerPoolSize = 100
const ChannelCapacity = 1000
const RetryCount = 3

var myMetrics metrics.Metrics

// Consumer code
func checkWorker(ctx context.Context, resourceCh <-chan health.HealthChecker, resultsCh chan<- health.HealthStatus) {
	for {
		start := time.Now()
		select {
		case resource, ok := <-resourceCh:
			if !ok {
				return
			}

			status := retryWith(resource)
			end := time.Since(start)
			status.ResponseTime = end
			// myMetrics.Increment(status.Severity)

			resultsCh <- status
		case <-ctx.Done():
			return
		}
	}
}

func retryWith(resource health.HealthChecker) health.HealthStatus {
	var hStatus health.HealthStatus

	for i := 1; i <= RetryCount; i++ {
		status, err := resource.CheckHealth()

		// if healthcheck SUCCEEDS, no retry, immediately return the status
		if err == nil {
			hStatus = status
			break
		}

		// if healthcheck FAILS, keep retrying and set the reason & return the status
		hStatus = status
		// set the reason here
		hStatus.Reason = err.Error()
	}
	return hStatus
}

// Driver code
func CheckResources(ctx context.Context, resources []health.HealthChecker) <-chan health.HealthStatus {
	// 1. Setup
	jobsCh := make(chan health.HealthChecker, ChannelCapacity)
	// resultsCh := make(chan health.HealthStatus, ChannelCapacity)

	// var wg sync.WaitGroup

	// // 2. Worker pool
	// for i := 1; i <= WorkerPoolSize; i++ {
	// 	wg.Go(func() {
	// 		// call worker code
	// 		checkWorker(ctx, jobsCh, resultsCh)
	// 	})
	// }

	// // 3. Co-ordinator block for synchronization
	// go func() {
	// 	wg.Wait()
	// 	close(resultsCh)
	// }()

	resultsCh := workerpool.Run(ctx, jobsCh, processResource)

	// 4. Producer code
	go func() {
		defer close(jobsCh)

		for _, resource := range resources {
			select {
			case <-ctx.Done():
				return
			case jobsCh <- resource:
			}
		}
	}()

	return resultsCh
}

// job function
func processResource(resource health.HealthChecker) health.HealthStatus {
	status, err := resource.CheckHealth()
	if err != nil {
		// deal with it
	}
	return status
}
