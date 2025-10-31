package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gittisak-go/gittisak-go/pkg/chatllm"
)

func main() {
	// Get credentials from environment variables
	apiKey := os.Getenv("ABACUS_API_KEY")
	deploymentToken := os.Getenv("ABACUS_DEPLOYMENT_TOKEN")
	deploymentID := os.Getenv("ABACUS_DEPLOYMENT_ID")

	if apiKey == "" {
		log.Fatal("ABACUS_API_KEY environment variable is required")
	}
	if deploymentToken == "" {
		log.Fatal("ABACUS_DEPLOYMENT_TOKEN environment variable is required")
	}
	if deploymentID == "" {
		log.Fatal("ABACUS_DEPLOYMENT_ID environment variable is required")
	}

	// Create a new ChatLLM client
	client, err := chatllm.NewClient(chatllm.Config{
		APIKey:          apiKey,
		DeploymentToken: deploymentToken,
		DeploymentID:    deploymentID,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create chat messages
	messages := []chatllm.ChatMessage{
		{
			IsUser: true,
			Text:   "What is the capital of France?",
		},
	}

	// Send the request with options
	resp, err := client.GetChatResponse(
		context.Background(),
		messages,
		chatllm.WithSystemMessage("You are a helpful assistant."),
		chatllm.WithLLMName("gpt-4"),
		chatllm.WithTemperature(0.7),
	)
	if err != nil {
		log.Fatalf("Failed to get chat response: %v", err)
	}

	// Print the response
	fmt.Printf("Conversation ID: %s\n", resp.DeploymentConversationID)
	fmt.Printf("\nChat History:\n")
	for i, msg := range resp.Messages {
		role := "User"
		if !msg.IsUser {
			role = "Assistant"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, role, msg.GetTextContent())
	}
}
