package reporter

import (
	"day08/health"
	"fmt"
)

func Report(statusChn <-chan health.HealthStatus) {
	for status := range statusChn {
		fmt.Println("-----------------")
		fmt.Printf("Resource: %v, Status: %v, Severity: %v, Reason: %v\n", status.Name, status.Healthy, status.Severity, status.Reason)
	}
}
