package main

import (
	"day04/health"
	"day04/loader"
	"day04/monitor"
	"day04/reporter"
	"log"
)

func main() {
	services, err := loader.LoadServices()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	databases, err := loader.LoadDatabases()

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
	monitor.Check(resources)

	// report services
	reporter.Report(resources)
}
