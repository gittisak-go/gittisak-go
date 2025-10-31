package chatllm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "valid config with API key",
			config: Config{
				APIKey:          "test-api-key",
				DeploymentToken: "test-token",
				DeploymentID:    "test-id",
			},
			wantErr: false,
		},
		{
			name: "valid config with custom base URL",
			config: Config{
				APIKey:          "test-api-key",
				DeploymentToken: "test-token",
				DeploymentID:    "test-id",
				BaseURL:         "https://custom.api.com",
			},
			wantErr: false,
		},
		{
			name: "missing API key",
			config: Config{
				DeploymentToken: "test-token",
				DeploymentID:    "test-id",
				BaseURL:         "https://api.abacus.ai",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if client == nil {
					t.Error("NewClient() returned nil client")
				}
				if client.apiKey != tt.config.APIKey {
					t.Errorf("NewClient() apiKey = %v, want %v", client.apiKey, tt.config.APIKey)
				}
				if client.deploymentToken != tt.config.DeploymentToken {
					t.Errorf("NewClient() deploymentToken = %v, want %v", client.deploymentToken, tt.config.DeploymentToken)
				}
				if client.deploymentID != tt.config.DeploymentID {
					t.Errorf("NewClient() deploymentID = %v, want %v", client.deploymentID, tt.config.DeploymentID)
				}
				expectedBaseURL := tt.config.BaseURL
				if expectedBaseURL == "" {
					expectedBaseURL = DefaultBaseURL
				}
				if client.baseURL != expectedBaseURL {
					t.Errorf("NewClient() baseURL = %v, want %v", client.baseURL, expectedBaseURL)
				}
			}
		})
	}
}

func TestNewClientWithAPIKey(t *testing.T) {
	apiKey := "test-api-key"
	client, err := NewClientWithAPIKey(apiKey)
	if err != nil {
		t.Fatalf("NewClientWithAPIKey() error = %v", err)
	}
	if client == nil {
		t.Fatal("NewClientWithAPIKey() returned nil client")
	}
	if client.apiKey != apiKey {
		t.Errorf("NewClientWithAPIKey() apiKey = %v, want %v", client.apiKey, apiKey)
	}
	if client.baseURL != DefaultBaseURL {
		t.Errorf("NewClientWithAPIKey() baseURL = %v, want %v", client.baseURL, DefaultBaseURL)
	}
}

func TestGetChatResponse(t *testing.T) {
	mockResponse := ChatResponse{
		DeploymentConversationID: "conv-123",
		Messages: []ResponseMessage{
			{
				IsUser:    true,
				Text:      "What is the capital of France?",
				Timestamp: "2025-10-31T10:00:00Z",
			},
			{
				IsUser:    false,
				Text:      "The capital of France is Paris.",
				Timestamp: "2025-10-31T10:00:01Z",
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/api/v0/getChatResponse" {
			t.Errorf("Expected path /api/v0/getChatResponse, got %s", r.URL.Path)
		}

		apiKeyHeader := r.Header.Get("apiKey")
		if apiKeyHeader != "test-api-key" {
			t.Errorf("Expected apiKey header 'test-api-key', got '%s'", apiKeyHeader)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	client, err := NewClient(Config{
		APIKey:          "test-api-key",
		DeploymentToken: "test-token",
		DeploymentID:    "test-deployment",
		BaseURL:         server.URL,
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	messages := []ChatMessage{
		{
			IsUser: true,
			Text:   "What is the capital of France?",
		},
	}

	resp, err := client.GetChatResponse(context.Background(), messages)
	if err != nil {
		t.Fatalf("GetChatResponse() error = %v", err)
	}

	if resp.DeploymentConversationID != mockResponse.DeploymentConversationID {
		t.Errorf("GetChatResponse() DeploymentConversationID = %v, want %v", resp.DeploymentConversationID, mockResponse.DeploymentConversationID)
	}
	if len(resp.Messages) != len(mockResponse.Messages) {
		t.Errorf("GetChatResponse() Messages length = %v, want %v", len(resp.Messages), len(mockResponse.Messages))
	}
}

func TestGetChatResponseEmptyMessages(t *testing.T) {
	client, err := NewClient(Config{
		APIKey:          "test-api-key",
		DeploymentToken: "test-token",
		DeploymentID:    "test-deployment",
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	messages := []ChatMessage{}

	_, err = client.GetChatResponse(context.Background(), messages)
	if err == nil {
		t.Error("GetChatResponse() expected error for empty messages, got nil")
	}
}

func TestChatOptions(t *testing.T) {
	req := ChatRequest{
		Messages: []ChatMessage{{IsUser: true, Text: "Hello"}},
	}

	WithLLMName("gpt-4")(&req)
	if req.LLMName != "gpt-4" {
		t.Errorf("WithLLMName() did not set LLMName correctly")
	}

	temp := 0.7
	WithTemperature(temp)(&req)
	if req.Temperature == nil || *req.Temperature != temp {
		t.Errorf("WithTemperature() did not set Temperature correctly")
	}

	WithSystemMessage("You are helpful")(&req)
	if req.SystemMessage != "You are helpful" {
		t.Errorf("WithSystemMessage() did not set SystemMessage correctly")
	}

	tokens := 100
	WithNumCompletionTokens(tokens)(&req)
	if req.NumCompletionTokens == nil || *req.NumCompletionTokens != tokens {
		t.Errorf("WithNumCompletionTokens() did not set NumCompletionTokens correctly")
	}
}

func TestResponseMessageGetTextContent(t *testing.T) {
	tests := []struct {
		name     string
		message  ResponseMessage
		expected string
	}{
		{
			name:     "string text",
			message:  ResponseMessage{Text: "Hello world"},
			expected: "Hello world",
		},
		{
			name:     "array text",
			message:  ResponseMessage{Text: []interface{}{"First message"}},
			expected: "First message",
		},
		{
			name:     "empty array",
			message:  ResponseMessage{Text: []interface{}{}},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.message.GetTextContent()
			if result != tt.expected {
				t.Errorf("GetTextContent() = %v, want %v", result, tt.expected)
			}
		})
	}
}
