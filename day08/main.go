package main

import (
	"context"
	"day08/health"
	"day08/loader"
	"day08/monitor"
	"day08/reporter"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
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

	// report services
	now := time.Now()
	reporter.Report(resultCh)
	fmt.Printf("Total time taken: %v", time.Since(now))
}
