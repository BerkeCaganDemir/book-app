package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

// ...existing code...
type JSONStore struct {
	FilePath string
	mu       sync.Mutex
}

func NewJSONStore(path string) *JSONStore {
	// Ensure directory exists and file is initialized
	dir := filepath.Dir(path)
	if dir != "." && dir != "/" {
		_ = os.MkdirAll(dir, 0755)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.WriteFile(path, []byte("[]"), 0644)
	}
	return &JSONStore{FilePath: path}
}

func (s *JSONStore) Read(data interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.ReadFile(s.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, data)
}

func (s *JSONStore) Write(data interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FilePath, bytes, 0644)
}

// ...existing code...
