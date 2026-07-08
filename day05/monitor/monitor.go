package monitor

import (
	"day05/health"
	"sync"
)

type Monitor struct{}

func Check(resources []health.HealthChecker) <-chan health.HealthStatus {
	var wg sync.WaitGroup
	statusChannel := make(chan health.HealthStatus)

	for _, resource := range resources {
		wg.Go(func() {
			// return resource status onto channel
			statusChannel <- resource.Healthy()
		})
	}

	go func() {
		// wait for all goroutines to complete, then close the channel
		wg.Wait()
		close(statusChannel)
	}()

	return statusChannel
}
