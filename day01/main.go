package main

import (
	"day01/healthcheck"
)

func main() {
	services := []healthcheck.Service{
		{Name: "auth-service"},
		{Name: "payment-service"},
	}

	healthcheck.ReportServiceStatus(services)
}
