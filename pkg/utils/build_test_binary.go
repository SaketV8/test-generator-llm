package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func RunCMakeAndMake(buildDir string, logFilePath string) error {
	PrintBoxedText("Building Test Binary")
	// Create or truncate the log file
	logFile, err := CreateEmptyFile(logFilePath)
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}
	defer logFile.Close()

	// Setup multiwriter to write to both stdout and log
	multiOut := io.MultiWriter(os.Stdout, logFile)

	// Run cmake
	cmakeCmd := exec.Command("cmake", "..", "-DBUILD_TESTING=ON")
	cmakeCmd.Dir = buildDir
	cmakeCmd.Stdout = multiOut
	cmakeCmd.Stderr = multiOut

	fmt.Fprintln(multiOut, "Running: cmake .. -DBUILD_TESTING=ON")
	if err := cmakeCmd.Run(); err != nil {
		fmt.Fprintf(multiOut, "cmake failed: %v\n", err)
		return fmt.Errorf("cmake failed: %w", err)
	}
	fmt.Fprintln(multiOut, "cmake succeeded.")

	// Run make
	makeCmd := exec.Command("make")
	makeCmd.Dir = buildDir
	makeCmd.Stdout = multiOut
	makeCmd.Stderr = multiOut

	fmt.Fprintln(multiOut, "Running: make")
	if err := makeCmd.Run(); err != nil {
		fmt.Fprintf(multiOut, "make failed: %v\n", err)
		return fmt.Errorf("make failed: %w", err)
	}
	fmt.Fprintln(multiOut, "\nmake succeeded.")
	PrintBoxedText("Building Test Binary succeeded")

	return nil
}
