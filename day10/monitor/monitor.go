package monitor

import (
	"context"
	"day10/health"
	"day10/workerpool"
	"time"
)

type Monitor struct{}

const WorkerPoolSize = 100
const ChannelCapacity = 1000
const RetryCount = 3

// var myMetrics metrics.Metrics

// Retry wrapper for a service
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
	return workerpool.Run(ctx, resources, processResource)
}

// job function
func processResource(resource health.HealthChecker, startTime time.Time) health.HealthStatus {
	status := retryWith(resource)
	end := time.Since(startTime)
	status.ResponseTime = end
	return status
}
