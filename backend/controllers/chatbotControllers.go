package controllers

import (
	"backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)
func GenerateNewIdeas(ideaList string) (string, error) {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("couldn't load dotenv")
	}
	// Your OpenAI API key
	apiKey := godotenv.Load("API_KEY")

	// Set up the API URL for GPT
	apiURL := "https://api.openai.com/v1/completions"

	// Define the prompt with the role and list of ideas
	prompt := fmt.Sprintf(`
You are a creative assistant tasked with generating new and unique ideas based on a list of previously done ideas.
Your goal is to come up with fresh ideas related to the following:

%s

Please generate 5 new, creative ideas that are inspired by these themes or topics.
`, ideaList)

	// Set up the request body with model, prompt, and other parameters
	requestBody := models.OpenAIRequest{
		Model:    "text-davinci-003", // Or another model if preferred
		Prompt:   prompt,
		MaxTokens: 200,
	}

	// Marshal the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	// Create a new HTTP request with the API URL and the JSON body
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	// Set the necessary headers for authorization and content-type
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Create an HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the JSON response into OpenAIResponse struct
	var openAIResponse models.OpenAIResponse
	if err := json.Unmarshal(body, &openAIResponse); err != nil {
		return "", err
	}

	// Return the generated text (new ideas)
	if len(openAIResponse.Choices) > 0 {
		return openAIResponse.Choices[0].Text, nil
	}

	return "", fmt.Errorf("No response from OpenAI API")
}
