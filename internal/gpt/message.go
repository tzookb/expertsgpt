package gpt

type Role int

const (
	System Role = iota
	Assistant
	User
)

type Message struct {
	User    Role
	Content string
}

func NewMessage(role Role, content string) *Message {
	msg := Message{
		role,
		content,
	}
	return &msg
}
