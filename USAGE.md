# Usage Guide - Abacus.AI ChatLLM Go Client

This guide demonstrates how to use the Abacus.AI ChatLLM Go client library.

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                     Your Application                        │
└────────────────────────┬────────────────────────────────────┘
                         │
                         │ Import package
                         ▼
┌─────────────────────────────────────────────────────────────┐
│          github.com/gittisak-go/gittisak-go/pkg/chatllm    │
│                                                              │
│  ┌─────────────┐    ┌──────────────┐   ┌─────────────┐    │
│  │   Client    │───▶│ ChatRequest  │───▶│ ChatResponse│    │
│  │  (Auth)     │    │  (Messages)  │   │  (Results)  │    │
│  └─────────────┘    └──────────────┘   └─────────────┘    │
└────────────────────────┬────────────────────────────────────┘
                         │
                         │ HTTPS POST
                         ▼
┌─────────────────────────────────────────────────────────────┐
│              Abacus.AI ChatLLM API                          │
│          https://api.abacus.ai/api/v0/getChatResponse       │
└─────────────────────────────────────────────────────────────┘
```

## Quick Start

### 1. Install the Library

```bash
go get github.com/gittisak-go/gittisak-go
```

### 2. Set Up Credentials

You need three credentials from your Abacus.AI account:

- **API Key**: From your Abacus.AI dashboard → API Keys
- **Deployment Token**: From your ChatLLM deployment settings
- **Deployment ID**: The unique ID of your ChatLLM deployment

Set them as environment variables:

```bash
export ABACUS_API_KEY="your-api-key"
export ABACUS_DEPLOYMENT_TOKEN="your-deployment-token"
export ABACUS_DEPLOYMENT_ID="your-deployment-id"
```

### 3. Create a Client

```go
import "github.com/gittisak-go/gittisak-go/pkg/chatllm"

client, err := chatllm.NewClient(chatllm.Config{
    APIKey:          os.Getenv("ABACUS_API_KEY"),
    DeploymentToken: os.Getenv("ABACUS_DEPLOYMENT_TOKEN"),
    DeploymentID:    os.Getenv("ABACUS_DEPLOYMENT_ID"),
})
```

### 4. Send a Message

```go
messages := []chatllm.ChatMessage{
    {
        IsUser: true,
        Text:   "Hello! How are you?",
    },
}

resp, err := client.GetChatResponse(context.Background(), messages)
if err != nil {
    log.Fatal(err)
}

// Print the assistant's response
for _, msg := range resp.Messages {
    if !msg.IsUser {
        fmt.Println("Assistant:", msg.GetTextContent())
    }
}
```

## Examples

### Example 1: Simple Question

```go
messages := []chatllm.ChatMessage{
    {IsUser: true, Text: "What is the capital of France?"},
}

resp, err := client.GetChatResponse(context.Background(), messages)
```

### Example 2: With Options (Model, Temperature, etc.)

```go
messages := []chatllm.ChatMessage{
    {IsUser: true, Text: "Explain quantum computing"},
}

resp, err := client.GetChatResponse(
    context.Background(),
    messages,
    chatllm.WithLLMName("gpt-4"),              // Choose the model
    chatllm.WithTemperature(0.7),               // Control randomness
    chatllm.WithSystemMessage("You are a physics teacher"),
    chatllm.WithNumCompletionTokens(200),      // Limit response length
)
```

### Example 3: Multi-turn Conversation

```go
// First message
messages := []chatllm.ChatMessage{
    {IsUser: true, Text: "I'm learning Go"},
}

resp, err := client.GetChatResponse(context.Background(), messages)

// Get assistant's reply
assistantReply := resp.Messages[len(resp.Messages)-1].GetTextContent()

// Continue conversation
messages = append(messages, chatllm.ChatMessage{
    IsUser: false,
    Text:   assistantReply,
})
messages = append(messages, chatllm.ChatMessage{
    IsUser: true,
    Text:   "What should I learn first?",
})

resp, err = client.GetChatResponse(context.Background(), messages)
```

## Available Options

| Option | Description | Example |
|--------|-------------|---------|
| `WithLLMName(name)` | Set the LLM model | `WithLLMName("gpt-4")` |
| `WithTemperature(temp)` | Set sampling temperature (0-2) | `WithTemperature(0.7)` |
| `WithSystemMessage(msg)` | Set system instructions | `WithSystemMessage("You are helpful")` |
| `WithNumCompletionTokens(n)` | Limit response tokens | `WithNumCompletionTokens(100)` |
| `WithChatConfig(config)` | Custom configuration | `WithChatConfig(map[string]interface{}{...})` |

## Response Structure

The `ChatResponse` contains:

```go
type ChatResponse struct {
    DeploymentConversationID string            // Unique conversation ID
    Messages                 []ResponseMessage // Full message history
    DocIDs                   []string          // Referenced documents
    KeywordArguments         map[string]string // Additional metadata
}
```

Each `ResponseMessage` has:

```go
type ResponseMessage struct {
    IsUser    bool        // true = user, false = assistant
    Text      interface{} // Message content (string or []string)
    Timestamp string      // When the message was sent
}
```

Use `msg.GetTextContent()` to safely extract text from any message.

## Running the Examples

### Basic Example

```bash
cd examples/basic
go run main.go
```

### Demo Example (Multiple Scenarios)

```bash
cd examples/demo
go run main.go
```

This will show:
1. Simple chat
2. Chat with options
3. Multi-turn conversation

## Error Handling

Always check for errors:

```go
resp, err := client.GetChatResponse(ctx, messages)
if err != nil {
    log.Printf("Error: %v", err)
    // Handle error appropriately
}
```

Common errors:
- Missing credentials
- Invalid deployment ID
- Network issues
- API rate limits

## Testing

Run the test suite:

```bash
go test ./...
```

With coverage:

```bash
go test -cover ./...
```

Current coverage: **81.0%**

## Support

- 📚 [Abacus.AI Documentation](https://abacus.ai/help/api/ref/predict/getChatResponse)
- 💬 [API Reference](https://abacus.ai/help/ref)
- 🐛 [Report Issues](https://github.com/gittisak-go/gittisak-go/issues)

## License

MIT License - See [LICENSE](../LICENSE) file
