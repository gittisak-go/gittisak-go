package chatllm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
				APIKey: "test-api-key",
			},
			wantErr: false,
		},
		{
			name: "valid config with custom base URL",
			config: Config{
				APIKey:  "test-api-key",
				BaseURL: "https://custom.api.com",
			},
			wantErr: false,
		},
		{
			name: "missing API key",
			config: Config{
				BaseURL: "https://api.abacus.ai",
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

func TestCreateChatCompletion(t *testing.T) {
	mockResponse := ChatCompletionResponse{
		ID:      "chatcmpl-123",
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   "gpt-4o",
		Choices: []Choice{
			{
				Index: 0,
				Message: Message{
					Role:    "assistant",
					Content: "The capital of France is Paris.",
				},
				FinishReason: "stop",
			},
		},
		Usage: Usage{
			PromptTokens:     20,
			CompletionTokens: 10,
			TotalTokens:      30,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/v1/chat/completions" {
			t.Errorf("Expected path /v1/chat/completions, got %s", r.URL.Path)
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer test-api-key" {
			t.Errorf("Expected Authorization header 'Bearer test-api-key', got '%s'", authHeader)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	client, err := NewClient(Config{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	req := ChatCompletionRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{
				Role:    "user",
				Content: "What is the capital of France?",
			},
		},
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateChatCompletion() error = %v", err)
	}

	if resp.ID != mockResponse.ID {
		t.Errorf("CreateChatCompletion() ID = %v, want %v", resp.ID, mockResponse.ID)
	}
	if resp.Model != mockResponse.Model {
		t.Errorf("CreateChatCompletion() Model = %v, want %v", resp.Model, mockResponse.Model)
	}
	if len(resp.Choices) != 1 {
		t.Errorf("CreateChatCompletion() Choices length = %v, want 1", len(resp.Choices))
	}
	if resp.Choices[0].Message.Content != mockResponse.Choices[0].Message.Content {
		t.Errorf("CreateChatCompletion() Content = %v, want %v", resp.Choices[0].Message.Content, mockResponse.Choices[0].Message.Content)
	}
}

func TestCreateChatCompletionEmptyMessages(t *testing.T) {
	client, err := NewClientWithAPIKey("test-api-key")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	req := ChatCompletionRequest{
		Model:    "gpt-4o",
		Messages: []Message{},
	}

	_, err = client.CreateChatCompletion(context.Background(), req)
	if err == nil {
		t.Error("CreateChatCompletion() expected error for empty messages, got nil")
	}
}
