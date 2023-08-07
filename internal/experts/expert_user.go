package experts

type ExpertUser interface {
	GetPrompt(data string) string
}
