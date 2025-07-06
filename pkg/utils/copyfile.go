package utils

import (
	"io"
	"os"
	"path/filepath"
)

// CopyFile copies the contents from src to dst.

// It creates the parent directory of dst if needed.
func CopyFile(src, dst string) error {
	// Open source file
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	// Ensure destination directory exists
	dir := filepath.Dir(dst)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Create destination file
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy contents
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	// Ensure contents are flushed
	return out.Sync()
}
