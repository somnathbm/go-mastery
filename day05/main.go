package main

import (
	"day05/health"
	"day05/loader"
	"day05/monitor"
	"day05/reporter"
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
	statusCh := monitor.Check(resources)

	// report services
	reporter.Report(statusCh)
}
