package utils

import (
	"os"
	"path/filepath"
)

// RenameFile renames (or moves) a file from oldPath to newPath.

// It creates the destination directory if needed.
func RenameFile(oldPath, newPath string) error {
	// Ensure parent directory of newPath exists
	dir := filepath.Dir(newPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Rename the file
	return os.Rename(oldPath, newPath)
}
