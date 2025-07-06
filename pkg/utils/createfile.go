package utils

import (
	"os"
	"path/filepath"
)

// CreateEmptyFile ensures the file exists and is empty.

// It creates parent directories if needed.
func CreateEmptyFile(path string) (*os.File, error) {
	// Ensure parent directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	// Create or truncate the file
	return os.Create(path)
}
