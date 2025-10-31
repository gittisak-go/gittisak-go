package chatllm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the default base URL for Abacus.AI ChatLLM API
	DefaultBaseURL = "https://api.abacus.ai"
	// DefaultTimeout is the default HTTP client timeout
	DefaultTimeout = 30 * time.Second
)

// Client represents an Abacus.AI ChatLLM client
type Client struct {
	apiKey          string
	deploymentToken string
	deploymentID    string
	baseURL         string
	httpClient      *http.Client
}

// Config holds the configuration for creating a new Client
type Config struct {
	APIKey          string
	DeploymentToken string
	DeploymentID    string
	BaseURL         string
	Timeout         time.Duration
}

// NewClient creates a new ChatLLM client with the given configuration
func NewClient(config Config) (*Client, error) {
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	timeout := config.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	return &Client{
		apiKey:          config.APIKey,
		deploymentToken: config.DeploymentToken,
		deploymentID:    config.DeploymentID,
		baseURL:         baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

// NewClientWithAPIKey creates a new ChatLLM client with just an API key
func NewClientWithAPIKey(apiKey string) (*Client, error) {
	return NewClient(Config{
		APIKey: apiKey,
	})
}

// doRequest performs an HTTP request with the configured client
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apiKey", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return resp, nil
}
