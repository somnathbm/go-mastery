package reporter

import (
	"day04/health"
	"fmt"
)

func Report(resources []health.HealthChecker) {
	for _, resource := range resources {
		fmt.Println("-----------------")
		fmt.Println(resource)
	}
}
