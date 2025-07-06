package utils

import (
	"fmt"
	"os"
)

// DeleteFile deletes the file at the given path and returns an error if it fails.
func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("failed to delete file %s: %w", path, err)
	}
	return nil
}
