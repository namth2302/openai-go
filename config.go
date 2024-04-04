package openai

const openaiAPIURLv1 = "https://api.openai.com/v1"

type ClientConfig struct {
	OpenAiKey string
	OrgID     string
	BaseURL   string
}

func DefaultConfig(openAiKey string) ClientConfig {
	return ClientConfig{
		OpenAiKey: openAiKey,
		BaseURL:   openaiAPIURLv1,
	}
}

func (ClientConfig) String() string {
	return "<OpenAI API ClientConfig>"
}
