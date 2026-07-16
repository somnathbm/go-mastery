package monitor

import (
	"context"
	"day09/health"
	"day09/metrics"
	"sync"
)

type Monitor struct{}

const WorkerPoolSize = 100
const ChannelCapacity = 1000

var myMetrics metrics.Metrics

// Consumer code
func checkWorker(ctx context.Context, resourceCh <-chan health.HealthChecker, resultsCh chan<- health.HealthStatus) {
	for {
		select {
		case resource, ok := <-resourceCh:
			if !ok {
				return
			}

			status := resource.Healthy()
			// myMetrics.Increment(status.Severity)
			resultsCh <- status
		case <-ctx.Done():
			return
		}
	}
}

// Driver code
func CheckResources(ctx context.Context, resources []health.HealthChecker) <-chan health.HealthStatus {
	// 1. Setup
	resourceCh := make(chan health.HealthChecker, ChannelCapacity)
	resultsCh := make(chan health.HealthStatus, ChannelCapacity)

	var wg sync.WaitGroup

	// 2. Worker pool
	for i := 1; i <= WorkerPoolSize; i++ {
		wg.Go(func() {
			// call worker code
			checkWorker(ctx, resourceCh, resultsCh)
		})
	}

	// 3. Co-ordinator block for synchronization
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	// 4. Producer code
	go func() {
		defer close(resourceCh)

		for _, resource := range resources {
			select {
			case <-ctx.Done():
				return
			case resourceCh <- resource:
			}
		}
	}()

	return resultsCh
}
