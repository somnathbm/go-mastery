package main

import (
	"day04/contracts"
	"day04/loader"
	"day04/monitor"
	"day04/reporter"
	"log"
)

func main() {
	services, err := loader.LoadServices()
	databases, err := loader.LoadDatabases()

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// resources
	var resources []contracts.HealthChecker

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
