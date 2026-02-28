package tools

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestEchoTool(t *testing.T) {
	tests := []struct {
		name        string
		arguments   map[string]interface{}
		expectError bool
		expectText  string
	}{
		{
			name: "Valid message",
			arguments: map[string]interface{}{
				"message": "Hello, World!",
			},
			expectError: false,
			expectText:  "Echo: Hello, World!",
		},
		{
			name: "Empty message",
			arguments: map[string]interface{}{
				"message": "",
			},
			expectError: false,
			expectText:  "Echo: ",
		},
		{
			name:        "Missing message argument",
			arguments:   map[string]interface{}{},
			expectError: true,
			expectText:  "Error: 'message' argument must be a string",
		},
		{
			name: "Invalid message type",
			arguments: map[string]interface{}{
				"message": 123,
			},
			expectError: true,
			expectText:  "Error: 'message' argument must be a string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := EchoTool(tt.arguments)
			
			if err != nil {
				t.Errorf("EchoTool returned error: %v", err)
			}
			
			if result.IsError != tt.expectError {
				t.Errorf("Expected IsError=%v, got %v", tt.expectError, result.IsError)
			}
			
			if len(result.Content) != 1 {
				t.Fatalf("Expected 1 content item, got %d", len(result.Content))
			}
			
			if result.Content[0].Type != "text" {
				t.Errorf("Expected content type 'text', got '%s'", result.Content[0].Type)
			}
			
			if result.Content[0].Text != tt.expectText {
				t.Errorf("Expected text '%s', got '%s'", tt.expectText, result.Content[0].Text)
			}
		})
	}
}

func TestGetTimeTool(t *testing.T) {
	beforeCall := time.Now()
	
	result, err := GetTimeTool(map[string]interface{}{})
	
	if err != nil {
		t.Errorf("GetTimeTool returned error: %v", err)
	}
	
	if result.IsError {
		t.Error("GetTimeTool should not return error result")
	}
	
	if len(result.Content) != 1 {
		t.Fatalf("Expected 1 content item, got %d", len(result.Content))
	}
	
	if result.Content[0].Type != "text" {
		t.Errorf("Expected content type 'text', got '%s'", result.Content[0].Type)
	}
	
	text := result.Content[0].Text
	if !strings.HasPrefix(text, "Current server time:") {
		t.Errorf("Expected text to start with 'Current server time:', got '%s'", text)
	}
	
	// Extract and parse the timestamp
	timeStr := strings.TrimPrefix(text, "Current server time: ")
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		t.Errorf("Failed to parse returned time: %v", err)
	}
	
	// Verify the time is reasonable (within 1 second of the call)
	if parsedTime.Before(beforeCall.Add(-time.Second)) || parsedTime.After(time.Now().Add(time.Second)) {
		t.Errorf("Returned time %v is not close to expected time %v", parsedTime, beforeCall)
	}
}

func TestReadFileTool(t *testing.T) {
	// Create a temporary test file
	testContent := "This is test content for ReadFileTool"
	tmpFile, err := os.CreateTemp("", "mcp-test-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	
	if _, err := tmpFile.WriteString(testContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()
	
	tests := []struct {
		name        string
		arguments   map[string]interface{}
		expectError bool
		expectText  string
	}{
		{
			name: "Valid file path",
			arguments: map[string]interface{}{
				"path": tmpFile.Name(),
			},
			expectError: false,
			expectText:  testContent,
		},
		{
			name: "Nonexistent file",
			arguments: map[string]interface{}{
				"path": "/nonexistent/file.txt",
			},
			expectError: true,
			expectText:  "Error reading file:",
		},
		{
			name:        "Missing path argument",
			arguments:   map[string]interface{}{},
			expectError: true,
			expectText:  "Error: 'path' argument must be a string",
		},
		{
			name: "Invalid path type",
			arguments: map[string]interface{}{
				"path": 123,
			},
			expectError: true,
			expectText:  "Error: 'path' argument must be a string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ReadFileTool(tt.arguments)
			
			if err != nil {
				t.Errorf("ReadFileTool returned error: %v", err)
			}
			
			if result.IsError != tt.expectError {
				t.Errorf("Expected IsError=%v, got %v", tt.expectError, result.IsError)
			}
			
			if len(result.Content) != 1 {
				t.Fatalf("Expected 1 content item, got %d", len(result.Content))
			}
			
			if result.Content[0].Type != "text" {
				t.Errorf("Expected content type 'text', got '%s'", result.Content[0].Type)
			}
			
			if tt.expectError {
				// For error cases, check if the text starts with the expected prefix
				if !strings.HasPrefix(result.Content[0].Text, tt.expectText) {
					t.Errorf("Expected text to start with '%s', got '%s'", tt.expectText, result.Content[0].Text)
				}
			} else {
				// For success cases, check exact match
				if result.Content[0].Text != tt.expectText {
					t.Errorf("Expected text '%s', got '%s'", tt.expectText, result.Content[0].Text)
				}
			}
		})
	}
}

func TestReadFileToolWithDirectory(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "mcp-test-dir-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)
	
	result, err := ReadFileTool(map[string]interface{}{
		"path": tmpDir,
	})
	
	if err != nil {
		t.Errorf("ReadFileTool returned error: %v", err)
	}
	
	if !result.IsError {
		t.Error("Reading a directory should return an error result")
	}
	
	if len(result.Content) != 1 {
		t.Fatalf("Expected 1 content item, got %d", len(result.Content))
	}
	
	if !strings.Contains(result.Content[0].Text, "Error reading file:") {
		t.Errorf("Expected error message about reading file, got '%s'", result.Content[0].Text)
	}
}

// Benchmark tests
func BenchmarkEchoTool(b *testing.B) {
	args := map[string]interface{}{
		"message": "Benchmark test message",
	}
	
	for i := 0; i < b.N; i++ {
		_, _ = EchoTool(args)
	}
}

func BenchmarkGetTimeTool(b *testing.B) {
	args := map[string]interface{}{}
	
	for i := 0; i < b.N; i++ {
		_, _ = GetTimeTool(args)
	}
}

func BenchmarkReadFileTool(b *testing.B) {
	// Create a temporary test file
	testContent := "Benchmark test content"
	tmpFile, err := os.CreateTemp("", "mcp-benchmark-*.txt")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	
	if _, err := tmpFile.WriteString(testContent); err != nil {
		b.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()
	
	args := map[string]interface{}{
		"path": tmpFile.Name(),
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ReadFileTool(args)
	}
}
