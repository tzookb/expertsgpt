package roles

type RoleUser interface {
	GetPrompt(data string) string
}
