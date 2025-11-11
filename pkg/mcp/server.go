package mcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Server represents an MCP server
type Server struct {
	name    string
	version string
	tools   map[string]ToolHandler
	reader  *bufio.Reader
	writer  io.Writer
}

// ToolHandler is a function that handles tool execution
type ToolHandler func(arguments map[string]interface{}) (*CallToolResult, error)

// NewServer creates a new MCP server
func NewServer(name, version string) *Server {
	return &Server{
		name:    name,
		version: version,
		tools:   make(map[string]ToolHandler),
		reader:  bufio.NewReader(os.Stdin),
		writer:  os.Stdout,
	}
}

// RegisterTool registers a tool handler
func (s *Server) RegisterTool(name, description string, inputSchema InputSchema, handler ToolHandler) {
	s.tools[name] = handler
}

// Start starts the MCP server and listens for requests
func (s *Server) Start() error {
	log.Printf("Starting MCP server: %s v%s\n", s.name, s.version)

	for {
		// Read request from stdin
		line, err := s.reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("error reading input: %w", err)
		}

		// Parse request
		var req Request
		if err := json.Unmarshal(line, &req); err != nil {
			s.sendError(nil, -32700, "Parse error", err.Error())
			continue
		}

		// Handle request
		s.handleRequest(&req)
	}
}

// handleRequest processes an incoming request
func (s *Server) handleRequest(req *Request) {
	switch req.Method {
	case "initialize":
		s.handleInitialize(req)
	case "tools/list":
		s.handleListTools(req)
	case "tools/call":
		s.handleCallTool(req)
	case "notifications/initialized":
		// Client notification that initialization is complete
		log.Println("Client initialized")
	default:
		s.sendError(req.ID, -32601, "Method not found", fmt.Sprintf("Unknown method: %s", req.Method))
	}
}

// handleInitialize handles the initialize request
func (s *Server) handleInitialize(req *Request) {
	result := InitializeResult{
		ProtocolVersion: "2024-11-05",
		Capabilities: ServerCapabilities{
			Tools: &ToolsCapability{
				ListChanged: false,
			},
		},
		ServerInfo: ServerInfo{
			Name:    s.name,
			Version: s.version,
		},
	}

	s.sendResponse(req.ID, result)
}

// handleListTools handles the tools/list request
func (s *Server) handleListTools(req *Request) {
	var tools []Tool
	for name := range s.tools {
		// Get tool metadata from registered tools
		tools = append(tools, s.getToolDefinition(name))
	}

	result := ListToolsResult{
		Tools: tools,
	}

	s.sendResponse(req.ID, result)
}

// handleCallTool handles the tools/call request
func (s *Server) handleCallTool(req *Request) {
	// Parse params
	var params CallToolParams
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		s.sendError(req.ID, -32602, "Invalid params", err.Error())
		return
	}

	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		s.sendError(req.ID, -32602, "Invalid params", err.Error())
		return
	}

	// Find tool handler
	handler, exists := s.tools[params.Name]
	if !exists {
		s.sendError(req.ID, -32602, "Tool not found", fmt.Sprintf("Tool '%s' not found", params.Name))
		return
	}

	// Execute tool
	result, err := handler(params.Arguments)
	if err != nil {
		s.sendError(req.ID, -32603, "Tool execution error", err.Error())
		return
	}

	s.sendResponse(req.ID, result)
}

// sendResponse sends a JSON-RPC response
func (s *Server) sendResponse(id interface{}, result interface{}) {
	resp := Response{
		Jsonrpc: "2.0",
		ID:      id,
		Result:  result,
	}

	s.sendJSON(resp)
}

// sendError sends a JSON-RPC error response
func (s *Server) sendError(id interface{}, code int, message string, data interface{}) {
	resp := Response{
		Jsonrpc: "2.0",
		ID:      id,
		Error: &Error{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}

	s.sendJSON(resp)
}

// sendJSON sends a JSON object to stdout
func (s *Server) sendJSON(v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		log.Printf("Error marshaling response: %v\n", err)
		return
	}

	_, err = fmt.Fprintf(s.writer, "%s\n", data)
	if err != nil {
		log.Printf("Error writing response: %v\n", err)
	}
}

// getToolDefinition returns the tool definition for a registered tool
func (s *Server) getToolDefinition(name string) Tool {
	// This is a simple implementation. In a real scenario, you'd store
	// tool metadata when registering tools
	switch name {
	case "echo":
		return Tool{
			Name:        "echo",
			Description: "Echoes back the input text",
			InputSchema: InputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"message": map[string]interface{}{
						"type":        "string",
						"description": "The message to echo back",
					},
				},
				Required: []string{"message"},
			},
		}
	case "get_time":
		return Tool{
			Name:        "get_time",
			Description: "Returns the current server time",
			InputSchema: InputSchema{
				Type:       "object",
				Properties: map[string]interface{}{},
			},
		}
	case "read_file":
		return Tool{
			Name:        "read_file",
			Description: "Reads the content of a file",
			InputSchema: InputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"path": map[string]interface{}{
						"type":        "string",
						"description": "The path to the file to read",
					},
				},
				Required: []string{"path"},
			},
		}
	default:
		return Tool{
			Name:        name,
			Description: "No description available",
			InputSchema: InputSchema{
				Type:       "object",
				Properties: map[string]interface{}{},
			},
		}
	}
}
