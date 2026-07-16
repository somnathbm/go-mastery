package main

import (
	"context"
	"day09/filter"
	"day09/health"
	"day09/loader"
	"day09/monitor"
	"day09/reporter"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// go func() {
	// 	time.Sleep(50 * time.Microsecond)
	// 	cancel()
	// }()

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
	resultCh := monitor.CheckResources(ctx, resources)

	// filter out only unhealthy status
	statusCh := filter.Filter(ctx, resultCh)

	// report services
	reporter.Report(ctx, statusCh)
}
