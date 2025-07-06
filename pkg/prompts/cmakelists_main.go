package prompts

import (
	"github.com/saketV8/test-generator-llm/pkg/utils"
)

// var CMakeLists_MAIN = utils.ReadFile("./orgChartApi/CMakeLists.txt")

// func createCMakeListsMainPrompt() string {
// 	var cmakelists_content = utils.ReadFile("./orgChartApi/CMakeLists.txt")

// 	builder := strings.Builder{}
// 	builder.WriteString("\n")
// 	builder.WriteString("CMakeLists.txt")
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

// var CMakeLists_MAIN = createCMakeListsMainPrompt()
var CMakeLists_MAIN = utils.CreatePromptFromFile("./orgChartApi/CMakeLists.txt", "File: CMakeLists.txt")
