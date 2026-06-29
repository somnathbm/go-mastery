package healthcheck

import "fmt"

type Service struct {
	Name   string
	CPU    int
	Memory int
}

func getUsage(usage int) string {
	switch {
	case usage >= 90:
		return "CRITICAL"
	case usage >= 80:
		return "WARNING"
	default:
		return "NORMAL"
	}
}

func ReportServiceStatus(services []Service) {
	for i := range services {
		fmt.Println("---------------------")
		fmt.Println("Service: ", services[i].Name)
		fmt.Printf("CPU: %v%%, SEVERITY: %v\n", services[i].CPU, getUsage(services[i].CPU))
		fmt.Printf("Memory: %v%%, SEVERITY: %v\n", services[i].Memory, getUsage(services[i].Memory))
	}
}
