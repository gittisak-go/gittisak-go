# Gittisak-Go: Abacus.AI ChatLLM Go Client

A Go client library for the [Abacus.AI ChatLLM](https://abacus.ai/chat_llm-ent) enterprise API. This library provides a simple and idiomatic Go interface for interacting with Abacus.AI's ChatLLM service, which offers access to multiple state-of-the-art LLMs including GPT-4o, Claude 3.5 Sonnet, and Gemini 1.5.

## Features

- 🚀 Simple and intuitive API
- 🔐 Built-in authentication with API keys and deployment tokens
- 🎯 Type-safe request and response structures
- ⚡ Context support for cancellation and timeouts
- 🧪 Comprehensive test coverage
- 📝 Full documentation and examples
- ✅ **Works with the real Abacus.AI API**

## Installation

```bash
go get github.com/gittisak-go/gittisak-go
```

## Quick Start

### Prerequisites

To use this library, you need:

1. **API Key**: Get one from [Abacus.AI Dashboard](https://admin.abacus.ai)
   - Sign up and log in
   - Navigate to your API Keys page
   - Create and copy your API key

2. **Deployment Token and ID**: Create a ChatLLM deployment
   - Go to your Abacus.AI dashboard
   - Create or select a ChatLLM deployment
   - Note the deployment ID and token

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/gittisak-go/gittisak-go/pkg/chatllm"
)

func main() {
    // Create a new client
    client, err := chatllm.NewClient(chatllm.Config{
        APIKey:          "your-api-key",
        DeploymentToken: "your-deployment-token",
        DeploymentID:    "your-deployment-id",
    })
    if err != nil {
        log.Fatal(err)
    }

    // Create chat messages
    messages := []chatllm.ChatMessage{
        {
            IsUser: true,
            Text:   "What is the capital of France?",
        },
    }

    // Send the request
    resp, err := client.GetChatResponse(context.Background(), messages)
    if err != nil {
        log.Fatal(err)
    }

    // Print the response
    for _, msg := range resp.Messages {
        if !msg.IsUser {
            fmt.Println("Assistant:", msg.GetTextContent())
        }
    }
}
```

## Configuration

### Using Environment Variables

```bash
export ABACUS_API_KEY="your-api-key"
export ABACUS_DEPLOYMENT_TOKEN="your-deployment-token"
export ABACUS_DEPLOYMENT_ID="your-deployment-id"
```

```go
client, err := chatllm.NewClient(chatllm.Config{
    APIKey:          os.Getenv("ABACUS_API_KEY"),
    DeploymentToken: os.Getenv("ABACUS_DEPLOYMENT_TOKEN"),
    DeploymentID:    os.Getenv("ABACUS_DEPLOYMENT_ID"),
})
```

### Custom Configuration

```go
client, err := chatllm.NewClient(chatllm.Config{
    APIKey:          "your-api-key",
    DeploymentToken: "your-deployment-token",
    DeploymentID:    "your-deployment-id",
    BaseURL:         "https://api.abacus.ai",  // optional
    Timeout:         60 * time.Second,         // optional
})
```

## API Reference

### Client Creation

#### `NewClient(config Config) (*Client, error)`

Creates a new ChatLLM client with full configuration options.

```go
client, err := chatllm.NewClient(chatllm.Config{
    APIKey:          "your-api-key",
    DeploymentToken: "your-deployment-token",
    DeploymentID:    "your-deployment-id",
    BaseURL:         "https://api.abacus.ai", // optional
    Timeout:         30 * time.Second,        // optional
})
```

### Chat Operations

#### `GetChatResponse(ctx context.Context, messages []ChatMessage, options ...ChatOption) (*ChatResponse, error)`

Sends a chat request to the Abacus.AI getChatResponse endpoint.

**Parameters:**

- `ctx`: Context for cancellation and timeout
- `messages`: Array of chat messages
- `options`: Optional configuration (model, temperature, etc.)

**Message Structure:**

```go
type ChatMessage struct {
    IsUser bool   `json:"is_user"` // true for user, false for assistant
    Text   string `json:"text"`    // message content
}
```

**Example with Options:**

```go
messages := []chatllm.ChatMessage{
    {
        IsUser: true,
        Text:   "Explain quantum computing",
    },
}

resp, err := client.GetChatResponse(
    context.Background(),
    messages,
    chatllm.WithSystemMessage("You are a helpful assistant."),
    chatllm.WithLLMName("gpt-4"),
    chatllm.WithTemperature(0.7),
    chatllm.WithNumCompletionTokens(500),
)
```

### Available Options

- `WithLLMName(name string)` - Set the LLM model to use
- `WithTemperature(temp float64)` - Set sampling temperature (0-2)
- `WithSystemMessage(msg string)` - Set system instructions
- `WithNumCompletionTokens(tokens int)` - Set max tokens to generate
- `WithChatConfig(config map[string]interface{})` - Set additional configuration

### Response Structure

```go
type ChatResponse struct {
    DeploymentConversationID string            // Conversation ID
    Messages                 []ResponseMessage // Full message history
    DocIDs                   []string          // Referenced document IDs
    KeywordArguments         map[string]string // Additional metadata
}

type ResponseMessage struct {
    IsUser    bool        // true if from user
    Text      interface{} // message content (string or []string)
    Timestamp string      // message timestamp
}
```

Use `GetTextContent()` method to safely extract text from ResponseMessage:

```go
for _, msg := range resp.Messages {
    if !msg.IsUser {
        fmt.Println(msg.GetTextContent())
    }
}
```

## Examples

See the [examples](./examples) directory for complete working examples:

- [Basic Example](./examples/basic/main.go) - Simple chat with options

To run the example:

```bash
export ABACUS_API_KEY="your-api-key"
export ABACUS_DEPLOYMENT_TOKEN="your-deployment-token"
export ABACUS_DEPLOYMENT_ID="your-deployment-id"
go run examples/basic/main.go
```

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## Supported Models

The library supports all LLM models available through Abacus.AI ChatLLM deployments, including:

- GPT-4o
- GPT-4 Turbo
- GPT-3.5 Turbo
- Claude 3.5 Sonnet
- Claude 3 Opus
- Gemini 1.5 Pro
- Gemini 1.5 Flash
- And more

The available models depend on your Abacus.AI deployment configuration.

## Error Handling

The library returns descriptive errors that can be checked and handled:

```go
resp, err := client.GetChatResponse(ctx, messages)
if err != nil {
    log.Printf("Error getting chat response: %v", err)
    return
}
```

Common error scenarios:
- Missing API key, deployment token, or deployment ID
- Network errors
- API errors (invalid deployment, quota exceeded, etc.)
- Invalid message format

## Authentication

This library uses Abacus.AI's authentication system:

1. **API Key**: Sent in the `apiKey` header for API authentication
2. **Deployment Token**: Identifies your specific ChatLLM deployment
3. **Deployment ID**: The unique identifier for your deployment

All three are required for the client to work properly.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built for the [Abacus.AI ChatLLM](https://abacus.ai/chat_llm-ent) API
- Follows Go best practices and idiomatic patterns

## Support

For issues and questions:
- Open an issue on GitHub
- Check the [Abacus.AI documentation](https://abacus.ai/help/api/ref/predict/getChatResponse)
- Visit the [Abacus.AI API Reference](https://abacus.ai/help/ref)
