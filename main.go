package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// OpenAIClient represents the client for OpenAI API
type OpenAIClient struct {
	apiKey       string
	organization string
}

// NewOpenAIClient creates a new OpenAI API client
func NewOpenAIClient(apiKey string, organization string) *OpenAIClient {
	return &OpenAIClient{apiKey: apiKey, organization: organization}
}

// ChatCompletionRequest represents the request for chat completions
type ChatCompletionRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

// Message represents a single message in the conversation
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletions sends a request to the chat completions API endpoint
func (c *OpenAIClient) ChatCompletions(req ChatCompletionRequest) (string, error) {
	// Marshal the request payload to JSON
	payloadBytes, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	// Create a new HTTP request
	httpRequest, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	// Set request headers
	httpRequest.Header.Add("Content-Type", "application/json")
	httpRequest.Header.Add("Authorization", "Bearer "+c.apiKey)
	httpRequest.Header.Add("OpenAI-Organization", c.organization)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(httpRequest)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	// Initialize the OpenAI client
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}
	organization := os.Getenv("OPENAI_ORGANIZATION_ID")
	if organization == "" {
		log.Fatal("OPENAI_ORGANIZATION is not set")
	}
	openaiClient := NewOpenAIClient(apiKey, organization)

	// Prepare the chat completion request
	request := ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "system",
				Content: SystemPrompt(),
			},
			{
				Role:    "user",
				Content: "/help",
			},
		},
		Temperature: 0.2,
	}

	// Make the chat completions request
	response, err := openaiClient.ChatCompletions(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
