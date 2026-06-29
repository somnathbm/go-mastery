package main

import (
	"day03/loader"
	"day03/monitor"
	"day03/reporter"
	"log"
)

func main() {
	data, err := loader.LoadServices()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// monitor services
	monitor.Check(data)

	// report services
	reporter.ServiceStatus(data)
}
