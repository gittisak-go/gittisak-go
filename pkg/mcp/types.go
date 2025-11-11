package mcp

// JSON-RPC 2.0 Request structure
type Request struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

// JSON-RPC 2.0 Response structure
type Response struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}

// JSON-RPC 2.0 Error structure
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// MCP Initialize Request parameters
type InitializeParams struct {
	ProtocolVersion string       `json:"protocolVersion"`
	Capabilities    Capabilities `json:"capabilities"`
	ClientInfo      ClientInfo   `json:"clientInfo"`
}

// MCP Initialize Result
type InitializeResult struct {
	ProtocolVersion string            `json:"protocolVersion"`
	Capabilities    ServerCapabilities `json:"capabilities"`
	ServerInfo      ServerInfo        `json:"serverInfo"`
}

// Client capabilities
type Capabilities struct {
	Roots    *RootsCapability    `json:"roots,omitempty"`
	Sampling *SamplingCapability `json:"sampling,omitempty"`
}

// Server capabilities
type ServerCapabilities struct {
	Tools     *ToolsCapability     `json:"tools,omitempty"`
	Resources *ResourcesCapability `json:"resources,omitempty"`
	Prompts   *PromptsCapability   `json:"prompts,omitempty"`
}

// Tool capabilities
type ToolsCapability struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

// Resource capabilities
type ResourcesCapability struct {
	Subscribe   bool `json:"subscribe,omitempty"`
	ListChanged bool `json:"listChanged,omitempty"`
}

// Prompts capabilities
type PromptsCapability struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

// Roots capability
type RootsCapability struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

// Sampling capability
type SamplingCapability struct{}

// Client info
type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Server info
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Tool definition
type Tool struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	InputSchema InputSchema `json:"inputSchema"`
}

// Input schema for tools
type InputSchema struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Required   []string               `json:"required,omitempty"`
}

// List tools result
type ListToolsResult struct {
	Tools []Tool `json:"tools"`
}

// Call tool params
type CallToolParams struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments,omitempty"`
}

// Call tool result
type CallToolResult struct {
	Content []Content `json:"content"`
	IsError bool      `json:"isError,omitempty"`
}

// Content structure for tool results
type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
