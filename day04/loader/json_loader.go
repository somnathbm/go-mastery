package loader

import (
	"cmp"
	"day04/model"
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
)

var servicePath string = filepath.Join("data", "services.json")
var databasePath string = filepath.Join("data", "databases.json")

// Load all services
func LoadServices() ([]model.Service, error) {
	// read the file in memory
	content, err := os.ReadFile(servicePath)
	if err != nil {
		return nil, err
	}

	// it is like, converting the json string to go data structure
	var data []model.Service
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	// sort the struct slices alphabetically
	slices.SortFunc(data, func(a, b model.Service) int {
		return cmp.Compare(a.ServiceName, b.ServiceName)
	})
	return data, nil
}

// Load all databases
func LoadDatabases() ([]model.Database, error) {
	// read the file in memory
	content, err := os.ReadFile(databasePath)
	if err != nil {
		return nil, err
	}

	// it is like, converting the json string to go data structure
	var data []model.Database
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	// sort the struct slices alphabetically
	slices.SortFunc(data, func(a, b model.Database) int {
		return cmp.Compare(a.DBName, b.DBName)
	})
	return data, nil
}
