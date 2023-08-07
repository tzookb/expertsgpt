package openai

import (
	"tzook/gogpt/internal/gpt"

	"github.com/sashabaranov/go-openai"
)

func ConvertMessagesToOpenAiThread(msgs []gpt.Message) []openai.ChatCompletionMessage {
	returned := []openai.ChatCompletionMessage{}

	for _, msg := range msgs {
		returned = append(returned, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: msg.Content,
		})
	}
	return returned
}
