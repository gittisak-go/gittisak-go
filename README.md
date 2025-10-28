# Gittisak-Go: Abacus.AI ChatLLM Go Client

A Go client library for the [Abacus.AI ChatLLM](https://abacus.ai/chat_llm-ent) enterprise API. This library provides a simple and idiomatic Go interface for interacting with Abacus.AI's ChatLLM service, which offers access to multiple state-of-the-art LLMs including GPT-4o, Claude 3.5 Sonnet, and Gemini 1.5.

## Features

- 🚀 Simple and intuitive API
- 🔐 Built-in authentication handling
- 🎯 Type-safe request and response structures
- ⚡ Context support for cancellation and timeouts
- 🧪 Comprehensive test coverage
- 📝 Full documentation and examples

## Installation

```bash
go get github.com/gittisak-go/gittisak-go
```

## Quick Start

### Prerequisites

You'll need an Abacus.AI API key. To get one:

1. Visit [Abacus.AI](https://admin.abacus.ai)
2. Sign up and log in
3. Navigate to your API Keys page
4. Copy your API key

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
    client, err := chatllm.NewClientWithAPIKey("your-api-key")
    if err != nil {
        log.Fatal(err)
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
        log.Fatal(err)
    }

    // Print the response
    fmt.Println(resp.Choices[0].Message.Content)
}
```

## Configuration

### Using Environment Variables

```bash
export ABACUS_API_KEY="your-api-key"
```

```go
apiKey := os.Getenv("ABACUS_API_KEY")
client, err := chatllm.NewClientWithAPIKey(apiKey)
```

### Custom Configuration

```go
client, err := chatllm.NewClient(chatllm.Config{
    APIKey:  "your-api-key",
    BaseURL: "https://api.abacus.ai",
    Timeout: 60 * time.Second,
})
```

## API Reference

### Client Creation

#### `NewClient(config Config) (*Client, error)`

Creates a new ChatLLM client with full configuration options.

```go
client, err := chatllm.NewClient(chatllm.Config{
    APIKey:  "your-api-key",
    BaseURL: "https://api.abacus.ai", // optional
    Timeout: 30 * time.Second,        // optional
})
```

#### `NewClientWithAPIKey(apiKey string) (*Client, error)`

Creates a new ChatLLM client with just an API key, using default configuration.

```go
client, err := chatllm.NewClientWithAPIKey("your-api-key")
```

### Chat Completion

#### `CreateChatCompletion(ctx context.Context, req ChatCompletionRequest) (*ChatCompletionResponse, error)`

Sends a chat completion request to the ChatLLM API.

**Request Parameters:**

- `Model` (string, optional): The model to use (e.g., "gpt-4o", "claude-3.5-sonnet")
- `Messages` ([]Message, required): Array of messages in the conversation
- `Temperature` (*float64, optional): Sampling temperature (0-2)
- `MaxTokens` (*int, optional): Maximum tokens to generate
- `Stream` (bool, optional): Enable streaming responses

**Message Structure:**

- `Role` (string): "user", "assistant", or "system"
- `Content` (string): The message content

**Example:**

```go
req := chatllm.ChatCompletionRequest{
    Model: "gpt-4o",
    Messages: []chatllm.Message{
        {
            Role:    "system",
            Content: "You are a helpful assistant.",
        },
        {
            Role:    "user",
            Content: "Explain quantum computing in simple terms.",
        },
    },
}

resp, err := client.CreateChatCompletion(context.Background(), req)
```

## Examples

See the [examples](./examples) directory for complete working examples:

- [Basic Example](./examples/basic/main.go) - Simple chat completion

To run an example:

```bash
export ABACUS_API_KEY="your-api-key"
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

The library supports all models available through Abacus.AI ChatLLM, including:

- GPT-4o
- GPT-4 Turbo
- Claude 3.5 Sonnet
- Claude 3 Opus
- Gemini 1.5 Pro
- And more

Check the [Abacus.AI documentation](https://abacus.ai/chat_llm-ent) for the latest list of available models.

## Error Handling

The library returns descriptive errors that can be checked and handled:

```go
resp, err := client.CreateChatCompletion(ctx, req)
if err != nil {
    // Handle error
    log.Printf("Error creating chat completion: %v", err)
    return
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built for the [Abacus.AI ChatLLM](https://abacus.ai/chat_llm-ent) API
- Inspired by the OpenAI Go client library design patterns

## Support

For issues and questions:
- Open an issue on GitHub
- Check the [Abacus.AI documentation](https://abacus.ai/chat_llm-ent)
