package class

import (
	"bufio"
	"os"
	"strings"
)

type Entry struct {
	Route string
	Key   string
	Value string
}

type TextReader struct {
	Entries []Entry
}

// NewTextReader creates a new reader and loads entries from file
func NewTextReader(filePath string) (*TextReader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []Entry
	scanner := bufio.NewScanner(file)

	var current Entry
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			// Empty line = end of one block
			if current.Route != "" || current.Key != "" || current.Value != "" {
				entries = append(entries, current)
				current = Entry{}
			}
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "route":
			current.Route = val
		case "key":
			current.Key = val
		case "value":
			current.Value = val
		}
	}

	if current.Route != "" || current.Key != "" || current.Value != "" {
		entries = append(entries, current)
	}

	return &TextReader{Entries: entries}, nil
}

// GetValue finds value by route and key
func (r *TextReader) GetValue(route, key string) (string, bool) {
	for _, entry := range r.Entries {
		if entry.Route == route && entry.Key == key {
			return entry.Value, true
		}
	}
	return "", false
}
