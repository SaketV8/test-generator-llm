package utils

import "strings"

func CreatePromptFromFile(filepath string, filename string) string {
	var content = ReadFile(filepath)

	builder := strings.Builder{}
	builder.WriteString("\n\n")
	builder.WriteString("\n---\n")
	builder.WriteString(filename)
	builder.WriteString("\n---\n")
	builder.WriteString("\n")
	builder.WriteString(content)
	builder.WriteString("\n")

	// fmt.Println(result)
	result := builder.String()
	return result
}
