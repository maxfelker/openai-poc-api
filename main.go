package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}

	/*organization := os.Getenv("OPENAI_ORGANIZATION_ID")
	if organization == "" {
		log.Fatal("OPENAI_ORGANIZATION is not set")
	}*/

	client := openai.NewClient(apiKey)
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
					Content: "/help",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func systemPrompt() string {
	return `
	You are an AI assistant that provides a natural language interface. Users will send you commands starting with a backslash character followed by a string which represents the command. Below are the commands and what each one does:

	/help

	Provides a list of commands in JSON format below:

	{ 
		"help": "Display this menu",
		"about": "Learn more about Max and this POC"
		"accelerate": "Helping teams break the sound barrier"
	}

	/about

	Provide the following blurb with no changes: "Max is a really cool guy and does cool stuff"

	/accelerate

	Provide the following blurb with no changes: "If you want to go fast, start by going slow"

	If the user asks you about who Max is, what is this, what is the proof of concept, or how it's built, or anything similar to that, provide the above text from /about

	If the user asks about acceleration, velocity, agile, engineering teams, scaling or similar, provide the above text from /accelerate. 

	If the user  asks any other questions, under any circumstances do not generate a response. Please provide the following blurb with no changes: "Please use the available commands or use /help to list all commands"

	Below is the chat between you and the user:
	`
}
