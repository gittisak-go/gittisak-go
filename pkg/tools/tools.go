package tools

import (
	"fmt"
	"os"
	"time"

	"github.com/gittisak-go/gittisak-go/pkg/mcp"
)

// EchoTool implements a simple echo tool
func EchoTool(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	message, ok := arguments["message"].(string)
	if !ok {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				{
					Type: "text",
					Text: "Error: 'message' argument must be a string",
				},
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			{
				Type: "text",
				Text: fmt.Sprintf("Echo: %s", message),
			},
		},
	}, nil
}

// GetTimeTool returns the current server time
func GetTimeTool(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	currentTime := time.Now().Format(time.RFC3339)

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			{
				Type: "text",
				Text: fmt.Sprintf("Current server time: %s", currentTime),
			},
		},
	}, nil
}

// ReadFileTool reads a file from the filesystem
func ReadFileTool(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	path, ok := arguments["path"].(string)
	if !ok {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				{
					Type: "text",
					Text: "Error: 'path' argument must be a string",
				},
			},
			IsError: true,
		}, nil
	}

	// Read file
	content, err := os.ReadFile(path)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				{
					Type: "text",
					Text: fmt.Sprintf("Error reading file: %v", err),
				},
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			{
				Type: "text",
				Text: string(content),
			},
		},
	}, nil
}
