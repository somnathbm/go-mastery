package monitor

import (
	"day08/health"
	"sync"
)

type Monitor struct{}

const WorkerPoolSize = 100
const ChannelCapacity = 1000

// Consumer code
func checkWorker(resourceCh <-chan health.HealthChecker, resultsCh chan<- health.HealthStatus) {
	for resource := range resourceCh {
		resultsCh <- resource.Healthy()
	}
}

// Driver code
func CheckResources(resources []health.HealthChecker) <-chan health.HealthStatus {
	// 1. Setup
	resourceCh := make(chan health.HealthChecker, ChannelCapacity)
	resultsCh := make(chan health.HealthStatus, ChannelCapacity)

	var wg sync.WaitGroup

	// 2. Worker pool
	for i := 1; i <= WorkerPoolSize; i++ {
		wg.Go(func() {
			// call worker code
			checkWorker(resourceCh, resultsCh)
		})
	}

	// 3. Co-ordinator block for synchronization
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	// 4. Producer code
	go func() {
		for _, resource := range resources {
			resourceCh <- resource
		}
		close(resourceCh)
	}()

	return resultsCh
}
