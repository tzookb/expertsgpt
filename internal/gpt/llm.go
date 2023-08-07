package gpt

type Llm interface {
	AskRespondChat(conv Conversation) (error, string)
}
