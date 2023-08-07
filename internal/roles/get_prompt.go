package roles

import "fmt"

type GetPrompter struct {
	intro string
}

func (g *GetPrompter) GetPrompt(requirements string) string {
	return fmt.Sprintf(intro, requirements)
}
