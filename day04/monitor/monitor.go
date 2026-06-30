package monitor

import (
	"day04/contracts"
)

type Monitor struct{}

func Check(resources []contracts.HealthChecker) {
	for i := range resources {
		resource := resources[i]
		if !resource.Healthy() {
			resource.Restart()
		}
	}
}
