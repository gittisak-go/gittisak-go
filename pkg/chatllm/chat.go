package chatllm

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

// ChatMessage represents a message in the Abacus.AI format
type ChatMessage struct {
	IsUser bool   `json:"is_user"` // true if message is from user, false if from assistant
	Text   string `json:"text"`    // The message content
}

// ChatRequest represents a request to the Abacus.AI getChatResponse API
type ChatRequest struct {
	DeploymentToken     string                 `json:"deploymentToken"`
	DeploymentID        string                 `json:"deploymentId"`
	Messages            []ChatMessage          `json:"messages"`
	LLMName             string                 `json:"llmName,omitempty"`
	NumCompletionTokens *int                   `json:"numCompletionTokens,omitempty"`
	SystemMessage       string                 `json:"systemMessage,omitempty"`
	Temperature         *float64               `json:"temperature,omitempty"`
	ChatConfig          map[string]interface{} `json:"chatConfig,omitempty"`
}

// ChatResponse represents a response from the Abacus.AI getChatResponse API
type ChatResponse struct {
	DeploymentConversationID string              `json:"deploymentConversationId"`
	Messages                 []ResponseMessage   `json:"messages"`
	DocIDs                   []string            `json:"docIds,omitempty"`
	KeywordArguments         map[string]string   `json:"keywordArguments,omitempty"`
}

// ResponseMessage represents a message in the chat response
type ResponseMessage struct {
	IsUser           bool     `json:"isUser"`
	Text             interface{} `json:"text"` // Can be string or []string
	Timestamp        string   `json:"timestamp,omitempty"`
	IsUseful         *bool    `json:"isUseful,omitempty"`
	Feedback         string   `json:"feedback,omitempty"`
	DocIDs           []string `json:"docIds,omitempty"`
	KeywordArguments map[string]string `json:"keywordArguments,omitempty"`
}

// GetTextContent extracts text content from ResponseMessage.Text
func (r *ResponseMessage) GetTextContent() string {
	switch v := r.Text.(type) {
	case string:
		return v
	case []interface{}:
		if len(v) > 0 {
			if str, ok := v[0].(string); ok {
				return str
			}
		}
	}
	return ""
}

// GetChatResponse sends a chat request to the Abacus.AI getChatResponse endpoint
func (c *Client) GetChatResponse(ctx context.Context, messages []ChatMessage, options ...ChatOption) (*ChatResponse, error) {
	if len(messages) == 0 {
		return nil, fmt.Errorf("at least one message is required")
	}

	req := ChatRequest{
		DeploymentToken: c.deploymentToken,
		DeploymentID:    c.deploymentID,
		Messages:        messages,
	}

	// Apply options
	for _, opt := range options {
		opt(&req)
	}

	path := "/api/v0/getChatResponse"

	resp, err := c.doRequest(ctx, "POST", path, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &chatResp, nil
}

// ChatOption is a function that modifies a ChatRequest
type ChatOption func(*ChatRequest)

// WithLLMName sets the LLM model to use
func WithLLMName(name string) ChatOption {
	return func(r *ChatRequest) {
		r.LLMName = name
	}
}

// WithTemperature sets the sampling temperature
func WithTemperature(temp float64) ChatOption {
	return func(r *ChatRequest) {
		r.Temperature = &temp
	}
}

// WithSystemMessage sets the system message
func WithSystemMessage(msg string) ChatOption {
	return func(r *ChatRequest) {
		r.SystemMessage = msg
	}
}

// WithNumCompletionTokens sets the maximum number of tokens to generate
func WithNumCompletionTokens(tokens int) ChatOption {
	return func(r *ChatRequest) {
		r.NumCompletionTokens = &tokens
	}
}

// WithChatConfig sets additional chat configuration
func WithChatConfig(config map[string]interface{}) ChatOption {
	return func(r *ChatRequest) {
		r.ChatConfig = config
	}
}
