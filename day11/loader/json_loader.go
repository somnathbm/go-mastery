package loader

import (
	"cmp"
	"day11/model"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// Load all services
func LoadServices(path string) ([]model.Service, error) {
	// prepare the file path object
	servicePath := filepath.Join(strings.Split(path, ",")...)

	// read the file in memory
	content, err := os.ReadFile(servicePath)
	if err != nil {
		return nil, err
	}

	// it is like, converting the json string to go data structure
	var data []model.Service
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, fmt.Errorf("JSON parse error: %w", err)
	}

	// sort the struct slices alphabetically
	slices.SortFunc(data, func(a, b model.Service) int {
		return cmp.Compare(a.SName, b.SName)
	})
	return data, nil
}

// Load all databases
func LoadDatabases(path string) ([]model.Database, error) {
	// prepare the file path object
	databasePath := filepath.Join(strings.Split(path, ",")...)

	// read the file in memory
	content, err := os.ReadFile(databasePath)
	if err != nil {
		return nil, fmt.Errorf("file read error: %w", err)
	}

	// it is like, converting the json string to go data structure
	var data []model.Database
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, fmt.Errorf("JSON parse error: %w", err)
	}

	// sort the struct slices alphabetically
	slices.SortFunc(data, func(a, b model.Database) int {
		return cmp.Compare(a.DBName, b.DBName)
	})
	return data, nil
}
