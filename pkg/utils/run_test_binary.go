package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func RunTestBuildLogging(buildDir string, logFilePath string) error {
	// Create (or truncate) the log file
	PrintBoxedText("Running Test Binary")
	logFile, err := CreateEmptyFile(logFilePath)
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}
	defer logFile.Close()

	// Prepare make command
	cmd := exec.Command("./org_chart_test")
	cmd.Dir = buildDir

	// Send output to both stdout and the log file
	multiOut := io.MultiWriter(os.Stdout, logFile)
	cmd.Stdout = multiOut
	cmd.Stderr = multiOut

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(multiOut, "test run failed 🐳: %v\n", err)
		return fmt.Errorf("test run failed: %w", err)
	}

	fmt.Fprintln(multiOut, "\nTest run succeeded")
	PrintBoxedText("Test Binary run succeeded")
	return nil
}
