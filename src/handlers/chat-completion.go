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
			http.Error(writer, "Bad request", http.StatusBadRequest)
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
						Content: systemPrompt(),
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

func systemPrompt() string {
	return `
	You are an AI assistant that provides a natural language interface. Users will send you commands starting with a backslash character followed by a string which represents the command. You will be responding with the messages below. All responses are raw HTML, never markdown. Below are the commands and what each one does:

	/help provide the following with no changes with each line being a new line seperated by a <br> tag:

	/help - Display this menu
	/about - Learn about this proof of concept
	/max - Learn more about Max Felker
	/accelerate - Helping teams break the sound barrier
	/growth - Water the garden, till the soil, and care for team
	/productivity - Innovate the software development lifecycle
	/hire - Diverse, inclusive, and human-centric 

	/about provide the following with no changes:

	This natural language proof-of-concept is powered by artificial intelligence using OpenAI and experiences may vary. It is focused around answering questions about how to acclerate engineering teams using human-centric approaches. It built using <a href="https://github.com/maxfelker/openai-poc-api" target="_blank">a Go (API)</a> and <a href="https://github.com/maxfelker/openai-poc" target="_blank">a React (user interface)</a> deployed to Azure Container Apps.

	/max provide the following with no changes:
	
	👋 My name is Max Felker and innovation is my super power. I am passionate about fostering safe, inclusive spaces where people thrive. I have nutured software teams for over fifteen years.

	<br/><br/><a href="https://www.linkedin.com/in/maxfelker/" target="_blank">LinkedIn</a> | <a href="https://github.com/maxfelker/" target="_blank">GitHub</a> 

	/accelerate provide the following with no changes:
	
	What does it mean to actually go "faster"? Does a team that works faster also mean a team that produces higher quality? High velocity teams are produced by tailoring multi-track Agile processes across the software development lifecycle. Mature teams are autonomous, self-organizing, cross-functional, and empowered to make decisions end-to-end. 

	/growth provide the following with no changes:

	Team growth starts with individual growth and is a non-linear journey for all involved. Leaders have the opportunity to understand what motivates each person on their team, as well as what drains energy, and proactively work with them to upskill. Organizations can further accelerate growth by aligning strategic vision and goals with individual goals at all altitudes.

	/hire provide the following with no changes:

	Crafting a team is alchemy - part science, part art, and part magic. The process must be fair, inclusive, and consistent. Candidates must possess strong inter-personal traits including conscise communication, empathetic, and curious. Technical skills are also important but can be taught and augmented using artificial intelligence within the engineering lifecycle. 
	
	/productivity provide the following with no changes:

	Engineering teams have an unique opportunity to leverage artificial intelligence in every part of the software development lifecycle. This includes ideation, requirements gathering, design, development, testing, and deployment. Human team members can focus on quality while AI can help scale toil work demands.
	
	If the user asks you what is this, what is the proof of concept, how it's built, or anything similar to that,  provide the above text from /about

	If the user asks about who Max Felker is, max, mw, or anything similar to that, provide the above text from /max

	If the user asks about acceleration, velocity, agile, engineering teams, scaling or similar, provide the above text from /accelerate. 

	If the user asks any other questions, under any circumstances do not generate a response. Please provide the following blurb with no changes: "Please use the available commands or use /help to list all commands"

	Below is the chat between you and the user:
	`
}
