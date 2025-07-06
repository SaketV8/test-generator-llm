package main

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/saketV8/test-generator-llm/pkg/llm"
	"github.com/saketV8/test-generator-llm/pkg/models"
	"github.com/saketV8/test-generator-llm/pkg/prompts"
	"github.com/saketV8/test-generator-llm/pkg/utils"
)

var ERROR_DURING_BUILD = false

func main() {
	// Get the GitHub PAT from an environment variable
	utils.PrintBoxedText("Test Generator LLM Started")
	fmt.Print("--------------------------------\n\n\n")

	// "" --> empty build error for first time run
	startTestGenerator("")
	buildTestBinary()
	runTestBinary()
	cleanRemainFiles()

	if ERROR_DURING_BUILD {
		// run again but with extra error message
		// formatting the error message

		error_prompt := "\n\n" + utils.ReadFile("./pkg/prompts/raw_prompt/error_prompt.txt")
		build_error := "\n" + utils.ReadFile("./build-logs/build-cmake-make.log") + "\n\n"

		startTestGenerator(error_prompt + build_error)
		buildTestBinary()
		runTestBinary()
		cleanRemainFiles()

		ERROR_DURING_BUILD = false
	}

}

func startTestGenerator(extra_build_error string) {
	// generating the Code from LLM
	// STEP 1
	utils.PrintBoxedText("Requeting LLM to generate Test cases")

	system_prompt := utils.ReadFile("./pkg/prompts/raw_prompt/system_prompt.txt")
	llmOutput, err := llm.ChatWithLLM(system_prompt, prompts.FINAL_PROMPT_MAIN+extra_build_error, true)
	// llmOutput, err := llm.ChatWithLLM(system_prompt, prompts.DefaultPrompt, true)
	if err != nil {
		log.Fatalf("error occured in llm: %v", err)
	} else {
		utils.PrintBoxedText("Test cases generated")
	}

	// STEP 2
	// structure the lllmOutputl
	utils.PrintBoxedText("Parsing Test cases")
	var filesToWrite []models.FileToGenerate = llm.ParseLLmOutPut(llmOutput)

	// STEP 3
	utils.PrintBoxedText("--- Writing Files to temp test folder ---")
	for _, file := range filesToWrite {
		err := utils.WriteFile(file.Path, file.Content)
		if err != nil {
			fmt.Printf("Failed to write file %s. Error: %v\n", file.Path, err)
		} else {
			fmt.Printf("Successfully wrote file: %s\n", file.Path)
		}

	}
	fmt.Println("---------------------------")

	// STEP 4
	// renaming the current orgChartApi/test/test_controllers.cc to orgChartApi/test/test_controllers.cc.bkp
	// rename <.cc> to <.bkp> 🌸
	utils.PrintBoxedText("Copying File to main project")
	utils.RenameFile("./orgChartApi/test/test_controllers.cc", "./orgChartApi/test/test_controllers.cc.bkp")
	// copying the generated file to ./orgChartApi/test/test_controllers.cc
	utils.CopyFile("./test/test_controllers.cc", "./orgChartApi/test/test_controllers.cc")

}

func buildTestBinary() {
	// STEP 5
	// generating the test binary
	buildDir := "./orgChartApi/build"
	logFilePath := "./build-logs/build-cmake-make.log"
	err := utils.RunCMakeAndMake(buildDir, logFilePath)
	if err != nil {
		ERROR_DURING_BUILD = true
		fmt.Println("Error Occured during build")
	}
}

func runTestBinary() {
	if !ERROR_DURING_BUILD {
		// running the test binary
		runDir := "./orgChartApi/build/test"
		runLogFilePath := "./build-logs/test-run.log"
		utils.RunTestBuildLogging(runDir, runLogFilePath)
	} else {
		fmt.Println("Error during build, skipping the RUN (-_-)")
	}
}

func cleanRemainFiles() {
	// STEP 6
	// Cleaning
	// deleting <.cc> 🌸
	utils.PrintBoxedText("File's Remain Cleaning Started")
	err := utils.DeleteFile("./orgChartApi/test/test_controllers.cc")
	if err != nil {
		fmt.Println("Error:", err)
	}
	// rename <.bkp> to <.cc> 🌸
	utils.RenameFile("./orgChartApi/test/test_controllers.cc.bkp", "./orgChartApi/test/test_controllers.cc")

	utils.PrintBoxedText("File's Remain Cleaning Complete")
}
