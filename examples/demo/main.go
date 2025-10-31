package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gittisak-go/gittisak-go/pkg/chatllm"
)

func main() {
	fmt.Println("=== Abacus.AI ChatLLM Go Client Demo ===")

	// Get credentials from environment variables
	apiKey := os.Getenv("ABACUS_API_KEY")
	deploymentToken := os.Getenv("ABACUS_DEPLOYMENT_TOKEN")
	deploymentID := os.Getenv("ABACUS_DEPLOYMENT_ID")

	if apiKey == "" || deploymentToken == "" || deploymentID == "" {
		log.Fatal("Required environment variables:\n" +
			"  ABACUS_API_KEY\n" +
			"  ABACUS_DEPLOYMENT_TOKEN\n" +
			"  ABACUS_DEPLOYMENT_ID\n")
	}

	// Create a new ChatLLM client
	fmt.Println("✓ Creating Abacus.AI ChatLLM client...")
	client, err := chatllm.NewClient(chatllm.Config{
		APIKey:          apiKey,
		DeploymentToken: deploymentToken,
		DeploymentID:    deploymentID,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	fmt.Println("✓ Client created successfully")

	// Example 1: Simple chat
	fmt.Println("--- Example 1: Simple Chat ---")
	simpleChat(client)

	fmt.Println("\n--- Example 2: Chat with Options ---")
	chatWithOptions(client)

	fmt.Println("\n--- Example 3: Multi-turn Conversation ---")
	multiTurnChat(client)

	fmt.Println("\n=== Demo Complete ===")
}

func simpleChat(client *chatllm.Client) {
	messages := []chatllm.ChatMessage{
		{
			IsUser: true,
			Text:   "What is 2+2?",
		},
	}

	fmt.Println("User: What is 2+2?")
	resp, err := client.GetChatResponse(context.Background(), messages)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// Print assistant's response
	for _, msg := range resp.Messages {
		if !msg.IsUser {
			fmt.Printf("Assistant: %s\n", msg.GetTextContent())
		}
	}
	fmt.Printf("(Conversation ID: %s)\n", resp.DeploymentConversationID)
}

func chatWithOptions(client *chatllm.Client) {
	messages := []chatllm.ChatMessage{
		{
			IsUser: true,
			Text:   "Explain machine learning in one sentence.",
		},
	}

	fmt.Println("User: Explain machine learning in one sentence.")
	fmt.Println("Using: GPT-4, Temperature: 0.3, Max tokens: 50")

	resp, err := client.GetChatResponse(
		context.Background(),
		messages,
		chatllm.WithLLMName("gpt-4"),
		chatllm.WithTemperature(0.3),
		chatllm.WithNumCompletionTokens(50),
		chatllm.WithSystemMessage("You are a concise technical expert."),
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	for _, msg := range resp.Messages {
		if !msg.IsUser {
			fmt.Printf("Assistant: %s\n", msg.GetTextContent())
		}
	}
}

func multiTurnChat(client *chatllm.Client) {
	// Simulate a multi-turn conversation
	messages := []chatllm.ChatMessage{
		{
			IsUser: true,
			Text:   "I'm learning Go programming.",
		},
	}

	fmt.Println("User: I'm learning Go programming.")

	// First turn
	resp, err := client.GetChatResponse(
		context.Background(),
		messages,
		chatllm.WithSystemMessage("You are a helpful programming tutor."),
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// Get the assistant's response
	var assistantReply string
	for _, msg := range resp.Messages {
		if !msg.IsUser {
			assistantReply = msg.GetTextContent()
			fmt.Printf("Assistant: %s\n", assistantReply)
		}
	}

	// Second turn - add to conversation
	messages = append(messages, chatllm.ChatMessage{
		IsUser: false,
		Text:   assistantReply,
	})
	messages = append(messages, chatllm.ChatMessage{
		IsUser: true,
		Text:   "What's the best resource to learn?",
	})

	fmt.Println("\nUser: What's the best resource to learn?")

	resp, err = client.GetChatResponse(
		context.Background(),
		messages,
		chatllm.WithSystemMessage("You are a helpful programming tutor."),
	)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	for _, msg := range resp.Messages {
		if !msg.IsUser && msg.GetTextContent() != assistantReply {
			fmt.Printf("Assistant: %s\n", msg.GetTextContent())
		}
	}
}
