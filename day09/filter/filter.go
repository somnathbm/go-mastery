package filter

import (
	"context"
	"day09/health"
)

func Filter(ctx context.Context, statusCh <-chan health.HealthStatus) <-chan health.HealthStatus {
	filterCh := make(chan health.HealthStatus, 100)

	go func() {
		defer close(filterCh)

		for {
			select {
			case <-ctx.Done():
				// if parent cancels, return immediately
				return
			case status, ok := <-statusCh:
				if !ok {
					// if input channel is closed upstream, return immediately
					return
				}
				if status.Severity != health.SeverityNormal {
					select {
					case filterCh <- status:
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}()
	return filterCh
}
