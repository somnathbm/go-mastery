package main

import (
	"day04/loader"
	"log"
)

func main() {
	data, err := loader.LoadServices("data/services.json")
	if err != nil {
		log.Printf("Error: %T", err)
	}
	// databases, err := loader.LoadDatabases()

	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// }

	// // resources
	// var resources []health.HealthChecker

	// // iterate through resources
	// for i := range services {
	// 	resources = append(resources, &services[i])
	// }

	// for i := range databases {
	// 	resources = append(resources, &databases[i])
	// }

	// // monitor services
	// monitor.Check(resources)

	// // report services
	// reporter.Report(resources)
}
