package prompts

import "github.com/saketV8/test-generator-llm/pkg/utils"

var START_PROMPT = utils.ReadFile("./pkg/prompts/raw_prompt/start_prompt.txt")
var FILE_CONTENT_PROMPT = utils.ReadFile("./pkg/prompts/raw_prompt/file_content_prompt.txt")

var API_ENDPOINT_PROMPT = utils.ReadFile("./pkg/prompts/raw_prompt/end_prompt.txt")
var FILE_TREE_STRUCTURE_PROMPT = utils.ReadFile("./pkg/prompts/raw_prompt/file_struct.txt")

var END_PROMPT = utils.ReadFile("./pkg/prompts/raw_prompt/end_prompt.txt")

// use this in LLM
var FINAL_PROMPT_MAIN = START_PROMPT + FILE_CONTENT_PROMPT + API_ENDPOINT_PROMPT + FILE_TREE_STRUCTURE_PROMPT + END_PROMPT
