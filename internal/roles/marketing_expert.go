package roles

import "fmt"

// ## Search Information
// %s
// ## Format example
// {format_example}
const marketing_expert_info = `
Role: You are a professional seo expert; the goal is to design a concise, usable, efficient seo plan.
According to the context, fill in the following missing information, note that each sections are returned in json string.
If the requirements are unclear, ensure minimum viability and avoid excessive design
ATTENTION: return only a json string, not before messaging or anything that will break parsing the json message.

# Context
## Original Requirements
%s
-----


Fields:
## Original Requirements: Provide as Plain text, place the polished complete original requirements here

## SEO Keywords: Provided as an array, up to 10 search keywords to focus on, could be long tails as well.

## Keywords Blog Topics: Provided as an object of keyword=>array of subjects, the keywords are from the "SEO keywords" above.

"""
FORMAT_EXAMPLE = """
---
{
	"originalRequirements": "The boss...",
	"seoKeywords": [...],
	"keywordsBlogTopics": {key1: [...]},
}
`

func GetPromptMarketing(requirements string) string {
	return fmt.Sprintf(marketing_expert_info, requirements)
}
