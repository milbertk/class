package class

import (
	"encoding/json"
	"os"
)

type JSONReader struct {
	Data map[string]interface{}
}

// NewJSONReader reads a JSON file and stores it in a map
func NewJSONReader(filePath string) (*JSONReader, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return &JSONReader{Data: data}, nil
}

// GetJSON returns the whole data as JSON bytes (or map)
func (r *JSONReader) GetJSON() map[string]interface{} {
	return r.Data
}

// GetValue gets a specific value by key
func (r *JSONReader) GetValue(key string) (string, bool) {
	if val, ok := r.Data[key]; ok {
		return val.(string), true
	}
	return "", false
}
