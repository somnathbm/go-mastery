package main

import (
	"day07/health"
	"day07/loader"
	"day07/monitor"
	"day07/reporter"
	"log"
)

func main() {
	services, err := loader.LoadServices("data/services.json")
	if err != nil {
		log.Printf("Error: %T", err)
	}
	databases, err := loader.LoadDatabases("data/databases.json")

	if err != nil {
		log.Printf("Error: %v", err)
	}

	// resources
	var resources []health.HealthChecker

	// iterate through resources
	for i := range services {
		resources = append(resources, &services[i])
	}

	for i := range databases {
		resources = append(resources, &databases[i])
	}

	// monitor services
	resultCh := monitor.CheckResources(resources)

	// report services
	reporter.Report(resultCh)
}
