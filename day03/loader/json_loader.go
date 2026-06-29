package loader

import (
	"cmp"
	"day03/model"
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
)

var servicePath string = filepath.Join("data", "services.json")

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
		return cmp.Compare(a.Name, b.Name)
	})
	return data, nil
}
