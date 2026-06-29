package main

import (
	"day02/loader"
	"day02/reporter"
	"log"
)

func main() {
	data, err := loader.LoadServices()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	reporter.Status(data)
}
