package utils

import (
	"os"
	"path/filepath"
)

// func WriteFile(path string, content string) error {
// 	return os.WriteFile(path, []byte(content), 0644)
// }

func WriteFile(path string, content string) error {
	// Ensure the parent directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Write the file
	return os.WriteFile(path, []byte(content), 0644)
}
