package loader

import (
	"encoding/json"
	"errors"
	"io/fs"
	"testing"
)

func TestLoadService(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		wantErr     bool
		wantErrType string
	}{
		{"resource-404", "../data/servicess.json", true, "path"},
		{"resource-ok", "../data/services.json", false, ""},
		{"resource-bad", "../data/services-malformed.json", true, "json"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data, gotErr := LoadServices(tc.path)

			// Case 1: we expect no error. because both path + json is OK
			if !tc.wantErr {
				if gotErr != nil {
					t.Errorf("Expected no error, but got: %v", gotErr)
				}
				if len(data) == 0 {
					t.Errorf("Expected data to be loaded, but got empty slice")
				}
				return
			}

			// Case 2: we expect an error
			if gotErr == nil {
				t.Errorf("Expected an error, but got nil")
				return
			}

			// Case 3: we expected error, got error, now check what kind of error we got
			switch tc.wantErrType {
			case "path":
				var pathErr *fs.PathError
				if !errors.As(gotErr, &pathErr) {
					t.Errorf("Expected *fs.PathError, but got: %v", gotErr)
				}
			case "json":
				var syntaxErr *json.SyntaxError
				if !errors.As(gotErr, &syntaxErr) {
					t.Errorf("Expected *json.SyntaxError, but got: %v", gotErr)
				}
			}
		})
	}
}
