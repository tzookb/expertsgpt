package openai

import (
	"context"
	"tzook/gogpt/internal/gpt"

	"github.com/sashabaranov/go-openai"
)

type OpenAiAdapter struct {
	client *openai.Client
}

var _ gpt.Llm = (*OpenAiAdapter)(nil)

func NewAdapter(token string) *OpenAiAdapter {
	client := openai.NewClient(token)
	ad := OpenAiAdapter{client: client}
	return &ad
}
func (oa *OpenAiAdapter) AskRespondChat(conv gpt.Conversation) (error, string) {
	convereted := ConvertMessagesToOpenAiThread(conv.GetMessages())
	resp, err := oa.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Temperature: 1,
			Messages:    convereted,
		},
	)

	if err != nil {
		return err, ""
	}

	return nil, resp.Choices[0].Message.Content
}
