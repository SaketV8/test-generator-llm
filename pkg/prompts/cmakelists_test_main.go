package prompts

import "github.com/saketV8/test-generator-llm/pkg/utils"

// var CMakeLists_TEST_MAIN = utils.ReadFile("./orgChartApi/test/CMakeLists.txt")

// func createCMakeListsTestPrompt() string {
// 	var cmakelists_content = utils.ReadFile("./orgChartApi/test/CMakeLists.txt")

// 	builder := strings.Builder{}
// 	builder.WriteString("\n")
// 	builder.WriteString("CMakeLists.txt (./test/CMakeLists.txt)")
// 	builder.WriteString("\n")
// 	builder.WriteString("---")
// 	builder.WriteString("\n")
// 	builder.WriteString("\n")
// 	builder.WriteString(cmakelists_content)
// 	builder.WriteString("\n")
// 	builder.WriteString("\n")
// 	builder.WriteString("---")
// 	builder.WriteString("\n")

// 	// fmt.Println(result)
// 	result := builder.String()
// 	return result
// }

var CMakeLists_TEST_MAIN = utils.CreatePromptFromFile("./orgChartApi/test/CMakeLists.txt", "File: ./test/CMakeLists.txt")
