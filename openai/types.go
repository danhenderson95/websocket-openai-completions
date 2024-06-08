package openai

type TOpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type TOpenAIData struct {
	Model    string           `json:"model"`
	Messages []TOpenAIMessage `json:"messages"`
}

type TOpenAIChoice struct {
	Index        int `json:"index"`
	Message      TOpenAIMessage
	LogProbs     string `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

type TOpenAIUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type TOpenAIResponse struct {
	ID                string          `json:"id"`
	Object            string          `json:"object"`
	Created           int             `json:"created"`
	Choices           []TOpenAIChoice `json:"choices"`
	Usage             TOpenAIUsage    `json:"usage"`
	SystemFingerprint string          `json:"system_fingerprint"`
}
