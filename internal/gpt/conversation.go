package gpt

type Conversation struct {
	msgs []Message
}

func NewConversation() *Conversation {
	conv := Conversation{}
	return &conv
}

func (c *Conversation) AddMessage(msg Message) {
	c.msgs = append(c.msgs, msg)
}

func (c *Conversation) GetMessages() []Message {
	return c.msgs
}
