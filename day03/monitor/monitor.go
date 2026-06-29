package monitor

import (
	"day03/model"
)

type Monitor struct{}

func Check(services []model.Service) {
	for i := range services {
		service := &services[i]
		if !service.Healthy() {
			service.Restart()
		}
	}
}
