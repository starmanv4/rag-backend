package rag

import (
	"context" // ✅ Add missing import
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

var openaiClient *openai.Client

func InitOpenAI() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}
	openaiClient = openai.NewClient(apiKey)
}

func GenerateResponse(contextText, query string) string {
	resp, err := openaiClient.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{ // ✅ Fix struct name
			{Role: "system", Content: "You are a helpful assistant answering questions based on provided context."},
			{Role: "user", Content: "Context:\n" + contextText + "\n\nQuery: " + query},
		},
	})

	if err != nil {
		log.Println("Error generating response:", err)
		return "Error generating response"
	}

	return resp.Choices[0].Message.Content
}
