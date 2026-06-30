package reporter

import (
	"day04/contracts"
	"fmt"
)

func Report(resources []contracts.HealthChecker) {
	// severityMap := make(map[string]int, 3)
	for _, resource := range resources {

		// severity := resource.Severity()
		// severityMap[severity] += 1
		fmt.Println("-----------------")
		for k, v := range resource.Info() {
			fmt.Printf("%v: %v\n", k, v)
		}
		// fmt.Printf("Severity: %v\n", severity)

		// if service.RestartCount > 0 {
		// 	fmt.Printf("Service: %v was restarted...\n", service.ServiceName)
		// 	fmt.Printf("Restart count: %v\n", service.RestartCount)
		// }
	}

	// fmt.Println("\n-----SUMMARY--------")
	// fmt.Println("Services: ", len(services))
	// fmt.Printf("NORMAL: %v\n", severityMap["NORMAL"])
	// fmt.Printf("WARNING: %v\n", severityMap["WARNING"])
	// fmt.Printf("CRITICAL: %v\n", severityMap["CRITICAL"])
}
