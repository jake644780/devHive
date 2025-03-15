package models

type OpenAIRequest struct {
	Model    string `json:"model"`
	Prompt   string `json:"prompt"`
	MaxTokens int   `json:"max_tokens"`
}

// Define a structure for the response body
type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}