package rag

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

var openaiClient *openai.Client

func InitOpenAI() {
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}
	openaiClient = openai.NewClient(apiKey)
}

func GenerateResponse(contextText, query string) string {
	resp, err := openaiClient.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "Please give me answer about my query in native spoken English, based on the context, and only answer."},
			{Role: "user", Content: "Context:\n" + contextText + "\n\nQuery: " + query},
		},
	})

	if err != nil {
		log.Println("Error generating response:", err)
		return "Error generating response"
	}

	return resp.Choices[0].Message.Content
}
