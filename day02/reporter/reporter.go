package reporter

import (
	"day02/model"
	"fmt"
)

func Status(services []model.Service) {
	severityMap := make(map[string]int, 3)
	for _, service := range services {
		severity := service.Severity()
		severityMap[severity] += 1
		fmt.Println("-----------------")
		fmt.Printf("Service: %v\n", service.Name)
		fmt.Printf("CPU usage: %v%%\n", service.CPU)
		fmt.Printf("Memory usage: %v%%\n", service.Memory)
		fmt.Printf("Severity: %v\n", severity)

		if !service.Healthy() {
			service.Restart()
			fmt.Println("Restarting...")
			newSeverity := service.Severity()
			severityMap[severity] -= 1
			severityMap[newSeverity] += 1
			fmt.Printf("Service: %v was restarted...\n", service.Name)
			fmt.Printf("Restart count: %v\n", service.RestartCount)
			fmt.Printf("CPU: %v%%\n", service.CPU)
			fmt.Printf("Memory: %v%%\n", service.Memory)
			fmt.Printf("Post restart severity: %v\n", newSeverity)
		}
	}

	fmt.Println("\n-----SUMMARY--------")
	fmt.Println("Services: ", len(services))
	fmt.Printf("NORMAL: %v\n", severityMap["NORMAL"])
	fmt.Printf("WARNING: %v\n", severityMap["WARNING"])
	fmt.Printf("CRITICAL: %v\n", severityMap["CRITICAL"])
}
