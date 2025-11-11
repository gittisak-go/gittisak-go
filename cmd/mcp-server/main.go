package main

import (
	"log"

	"github.com/gittisak-go/gittisak-go/pkg/mcp"
	"github.com/gittisak-go/gittisak-go/pkg/tools"
)

func main() {
	// Create MCP server
	server := mcp.NewServer("gittisak-go-mcp-server", "1.0.0")

	// Register tools
	server.RegisterTool(
		"echo",
		"Echoes back the input text",
		mcp.InputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"message": map[string]interface{}{
					"type":        "string",
					"description": "The message to echo back",
				},
			},
			Required: []string{"message"},
		},
		tools.EchoTool,
	)

	server.RegisterTool(
		"get_time",
		"Returns the current server time",
		mcp.InputSchema{
			Type:       "object",
			Properties: map[string]interface{}{},
		},
		tools.GetTimeTool,
	)

	server.RegisterTool(
		"read_file",
		"Reads the content of a file",
		mcp.InputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"path": map[string]interface{}{
					"type":        "string",
					"description": "The path to the file to read",
				},
			},
			Required: []string{"path"},
		},
		tools.ReadFileTool,
	)

	// Start server
	log.Println("MCP Server starting...")
	if err := server.Start(); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
