package reporter

import (
	"context"
	"day09/health"
	"fmt"
)

func Report(ctx context.Context, statusChn <-chan health.HealthStatus) {
	for status := range statusChn {
		fmt.Println("-----------------")
		fmt.Println(status.Name)
		fmt.Printf("Severity: %v\n", status.Severity)
		fmt.Printf("Reason: %v\n", status.Reason)
		fmt.Println("-----------------------------------")
		// fmt.Printf("Resource: %v, Status: %v, Severity: %v, Reason: %v\n", status.Name, status.Healthy, status.Severity, status.Reason)
	}
}

// for {
// 	select {
// 	case <-ctx.Done():
// 		return
// 	case status, ok := <-statusChn:
// 		if !ok {
// 			return
// 		}
// 		fmt.Println("-----------------")
// 		fmt.Printf("Resource: %v, Status: %v, Severity: %v, Reason: %v\n", status.Name, status.Healthy, status.Severity, status.Reason)
// 	}
// }
// }
