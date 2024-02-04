package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/maxfelker/openai-poc-api/src/utils"
	"github.com/sashabaranov/go-openai"
)

type PostBody struct {
	Message string `json:"message"`
}

func ChatCompletion(client *openai.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var body PostBody
		err := json.NewDecoder(request.Body).Decode(&body)
		if err != nil {
			utils.ReturnError(writer, "Bad request", http.StatusBadRequest)
			return
		}

		if len(body.Message) < 3 {
			utils.ReturnError(writer, "Message must contain at least 3 characters", http.StatusBadRequest)
			return
		}

		if len(body.Message) > 50 {
			utils.ReturnError(writer, "Message length cannot exceed 50 characters", http.StatusBadRequest)
			return
		}

		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("OPENAI_API_KEY is not set")
		}

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleSystem,
						Content: SystemPrompt(),
					},
					{
						Role:    openai.ChatMessageRoleUser,
						Content: body.Message,
					},
				},
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}
		completionResponse := resp.Choices[0].Message.Content

		response, e := json.Marshal(completionResponse)
		if e != nil {
			utils.ReturnError(writer, e.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		writer.Write(response)
	}
}
