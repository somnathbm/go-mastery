package main

import (
	"context"
	"day10/filter"
	"day10/health"
	"day10/loader"
	"day10/monitor"
	"day10/reporter"
	"day10/scheduler"
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

	// +++++++++++ LOAD RESOURCES ++++++++++++
	services, err := loader.LoadServices("data/services.json")
	if err != nil {
		log.Printf("Error: %T", err)
	}
	databases, err := loader.LoadDatabases("data/databases.json")

	if err != nil {
		log.Printf("Error: %v", err)
	}

	// ++++++++++ 2. Prepare resource list ++++++++++
	// resources
	var resources []health.HealthChecker

	// iterate through resources
	for i := range services {
		resources = append(resources, &services[i])
	}

	for i := range databases {
		resources = append(resources, &databases[i])
	}

	// ++++++++ 3. Monitor cycle as callback
	callbackFunc := func(ctx context.Context) {
		resultCh := monitor.CheckResources(ctx, resources)
		// filter out only unhealthy status
		statusCh := filter.Filter(ctx, resultCh)

		// report services
		reporter.Report(ctx, statusCh)
	}

	// 4. Scheduler to execute the monitoring cycle
	sched := scheduler.New(10*time.Second, callbackFunc)
	sched.Start(ctx)
}
