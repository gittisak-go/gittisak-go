package mcp

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestNewServer(t *testing.T) {
	server := NewServer("test-server", "1.0.0")
	
	if server.name != "test-server" {
		t.Errorf("Expected name 'test-server', got '%s'", server.name)
	}
	
	if server.version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got '%s'", server.version)
	}
	
	if server.tools == nil {
		t.Error("Tools map should be initialized")
	}
	
	if server.toolMetadata == nil {
		t.Error("Tool metadata map should be initialized")
	}
}

func TestRegisterTool(t *testing.T) {
	server := NewServer("test-server", "1.0.0")
	
	// Define a simple test tool
	testHandler := func(args map[string]interface{}) (*CallToolResult, error) {
		return &CallToolResult{
			Content: []Content{{Type: "text", Text: "test"}},
		}, nil
	}
	
	schema := InputSchema{
		Type: "object",
		Properties: map[string]interface{}{
			"param": map[string]interface{}{
				"type": "string",
			},
		},
		Required: []string{"param"},
	}
	
	server.RegisterTool("test_tool", "A test tool", schema, testHandler)
	
	// Verify tool handler is registered
	if _, exists := server.tools["test_tool"]; !exists {
		t.Error("Tool handler should be registered")
	}
	
	// Verify tool metadata is registered
	metadata, exists := server.toolMetadata["test_tool"]
	if !exists {
		t.Error("Tool metadata should be registered")
	}
	
	if metadata.Name != "test_tool" {
		t.Errorf("Expected tool name 'test_tool', got '%s'", metadata.Name)
	}
	
	if metadata.Description != "A test tool" {
		t.Errorf("Expected description 'A test tool', got '%s'", metadata.Description)
	}
}

func TestHandleInitialize(t *testing.T) {
	server := NewServer("test-server", "1.0.0")
	
	var outputBuffer bytes.Buffer
	server.writer = &outputBuffer
	
	req := &Request{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "initialize",
		Params: map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities":    map[string]interface{}{},
			"clientInfo": map[string]interface{}{
				"name":    "test-client",
				"version": "1.0.0",
			},
		},
	}
	
	server.handleInitialize(req)
	
	var response Response
	err := json.Unmarshal(outputBuffer.Bytes()[:len(outputBuffer.Bytes())-1], &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}
	
	if response.Jsonrpc != "2.0" {
		t.Errorf("Expected jsonrpc '2.0', got '%s'", response.Jsonrpc)
	}
	
	if response.Error != nil {
		t.Errorf("Expected no error, got: %v", response.Error)
	}
	
	// Verify result contains protocol version
	result, ok := response.Result.(map[string]interface{})
	if !ok {
		t.Fatal("Result should be a map")
	}
	
	if result["protocolVersion"] != "2024-11-05" {
		t.Errorf("Expected protocol version '2024-11-05', got '%v'", result["protocolVersion"])
	}
}

func TestHandleListTools(t *testing.T) {
	server := NewServer("test-server", "1.0.0")
	
	// Register test tools
	testHandler := func(args map[string]interface{}) (*CallToolResult, error) {
		return &CallToolResult{
			Content: []Content{{Type: "text", Text: "test"}},
		}, nil
	}
	
	server.RegisterTool("tool1", "First tool", InputSchema{Type: "object"}, testHandler)
	server.RegisterTool("tool2", "Second tool", InputSchema{Type: "object"}, testHandler)
	
	var outputBuffer bytes.Buffer
	server.writer = &outputBuffer
	
	req := &Request{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "tools/list",
	}
	
	server.handleListTools(req)
	
	var response Response
	err := json.Unmarshal(outputBuffer.Bytes()[:len(outputBuffer.Bytes())-1], &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}
	
	if response.Error != nil {
		t.Errorf("Expected no error, got: %v", response.Error)
	}
	
	// Verify result contains tools
	resultBytes, _ := json.Marshal(response.Result)
	var result ListToolsResult
	json.Unmarshal(resultBytes, &result)
	
	if len(result.Tools) != 2 {
		t.Errorf("Expected 2 tools, got %d", len(result.Tools))
	}
}

func TestHandleCallTool(t *testing.T) {
	server := NewServer("test-server", "1.0.0")
	
	// Register a test tool
	testHandler := func(args map[string]interface{}) (*CallToolResult, error) {
		message, _ := args["message"].(string)
		return &CallToolResult{
			Content: []Content{{Type: "text", Text: "Echo: " + message}},
		}, nil
	}
	
	server.RegisterTool("echo", "Echo tool", InputSchema{
		Type: "object",
		Properties: map[string]interface{}{
			"message": map[string]interface{}{"type": "string"},
		},
	}, testHandler)
	
	var outputBuffer bytes.Buffer
	server.writer = &outputBuffer
	
	req := &Request{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "tools/call",
		Params: map[string]interface{}{
			"name": "echo",
			"arguments": map[string]interface{}{
				"message": "test message",
			},
		},
	}
	
	server.handleCallTool(req)
	
	var response Response
	err := json.Unmarshal(outputBuffer.Bytes()[:len(outputBuffer.Bytes())-1], &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}
	
	if response.Error != nil {
		t.Errorf("Expected no error, got: %v", response.Error)
	}
}

func TestHandleCallToolNotFound(t *testing.T) {
	server := NewServer("test-server", "1.0.0")
	
	var outputBuffer bytes.Buffer
	server.writer = &outputBuffer
	
	req := &Request{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "tools/call",
		Params: map[string]interface{}{
			"name":      "nonexistent",
			"arguments": map[string]interface{}{},
		},
	}
	
	server.handleCallTool(req)
	
	var response Response
	err := json.Unmarshal(outputBuffer.Bytes()[:len(outputBuffer.Bytes())-1], &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}
	
	if response.Error == nil {
		t.Error("Expected error for nonexistent tool")
	}
	
	if response.Error.Code != -32602 {
		t.Errorf("Expected error code -32602, got %d", response.Error.Code)
	}
}

func TestSendResponse(t *testing.T) {
	server := NewServer("test-server", "1.0.0")
	
	var outputBuffer bytes.Buffer
	server.writer = &outputBuffer
	
	result := map[string]string{"status": "ok"}
	server.sendResponse(123, result)
	
	var response Response
	err := json.Unmarshal(outputBuffer.Bytes()[:len(outputBuffer.Bytes())-1], &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}
	
	if response.ID != 123.0 {
		t.Errorf("Expected ID 123, got %v", response.ID)
	}
	
	if response.Jsonrpc != "2.0" {
		t.Errorf("Expected jsonrpc '2.0', got '%s'", response.Jsonrpc)
	}
}

func TestSendError(t *testing.T) {
	server := NewServer("test-server", "1.0.0")
	
	var outputBuffer bytes.Buffer
	server.writer = &outputBuffer
	
	server.sendError(456, -32600, "Invalid Request", "Test error data")
	
	var response Response
	err := json.Unmarshal(outputBuffer.Bytes()[:len(outputBuffer.Bytes())-1], &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}
	
	if response.ID != 456.0 {
		t.Errorf("Expected ID 456, got %v", response.ID)
	}
	
	if response.Error == nil {
		t.Fatal("Expected error in response")
	}
	
	if response.Error.Code != -32600 {
		t.Errorf("Expected error code -32600, got %d", response.Error.Code)
	}
	
	if response.Error.Message != "Invalid Request" {
		t.Errorf("Expected error message 'Invalid Request', got '%s'", response.Error.Message)
	}
}
