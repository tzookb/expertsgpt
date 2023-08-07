package runner

import (
	"fmt"
	internalopenai "tzook/gogpt/api/openai"
	"tzook/gogpt/internal/gpt"
	"tzook/gogpt/internal/roles"
)

type App struct {
	llm gpt.Llm
}

func NewApp(apiToken string) *App {
	llm := internalopenai.NewAdapter(apiToken)
	a := App{llm: llm}
	return &a
}

func (a *App) Boom() {
	// prompt := roles.GetPrompt("rent estimator tool")
	prompt := roles.GetPromptMarketing("home cleaning service in boca raton")
	conv := gpt.NewConversation()
	// conv.AddMessage(*gpt.NewMessage(gpt.Assistant, "be a chill guy and respond accordingly"))
	conv.AddMessage(*gpt.NewMessage(gpt.User, prompt))

	_, resp := a.llm.AskRespondChat(*conv)
	fmt.Println(resp)
}
