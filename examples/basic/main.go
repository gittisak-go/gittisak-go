package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gittisak-go/gittisak-go/pkg/chatllm"
)

func main() {
	// Get API key from environment variable
	apiKey := os.Getenv("ABACUS_API_KEY")
	if apiKey == "" {
		log.Fatal("ABACUS_API_KEY environment variable is required")
	}

	// Create a new ChatLLM client
	client, err := chatllm.NewClientWithAPIKey(apiKey)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create a chat completion request
	req := chatllm.ChatCompletionRequest{
		Model: "gpt-4o",
		Messages: []chatllm.Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant.",
			},
			{
				Role:    "user",
				Content: "What is the capital of France?",
			},
		},
	}

	// Send the request
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to create chat completion: %v", err)
	}

	// Print the response
	fmt.Printf("Model: %s\n", resp.Model)
	fmt.Printf("Response: %s\n", resp.Choices[0].Message.Content)
	fmt.Printf("Usage: %d prompt tokens, %d completion tokens, %d total tokens\n",
		resp.Usage.PromptTokens,
		resp.Usage.CompletionTokens,
		resp.Usage.TotalTokens,
	)
}
