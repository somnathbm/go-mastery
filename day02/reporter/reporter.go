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
	}

	fmt.Println("\n-----SUMMARY--------")
	fmt.Println("Services: ", len(services))
	fmt.Printf("NORMAL: %v\n", severityMap["NORMAL"])
	fmt.Printf("WARNING: %v\n", severityMap["WARNING"])
	fmt.Printf("CRITICAL: %v\n", severityMap["CRITICAL"])
}
