# MCP Server Usage Examples

This document provides practical examples of using the Gittisak Go MCP Server with various clients.

## Example 1: Using with Claude Desktop

### Step 1: Build the Server
```bash
cd /path/to/gittisak-go
make build
```

### Step 2: Configure Claude Desktop

Edit your Claude Desktop config file:
- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`
- **Linux**: `~/.config/Claude/claude_desktop_config.json`

Add this configuration:
```json
{
  "mcpServers": {
    "gittisak-go": {
      "command": "/absolute/path/to/gittisak-go/bin/mcp-server",
      "args": [],
      "env": {}
    }
  }
}
```

### Step 3: Restart Claude Desktop

After restarting, you can ask Claude:
- "What tools do you have available?"
- "Echo back the message 'Hello MCP!'"
- "What time is it on the server?"
- "Read the file at /path/to/myfile.txt"

## Example 2: Testing with Command Line

You can test the server directly from the command line:

### Initialize the connection:
```bash
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"cli-test","version":"1.0"}}}' | ./bin/mcp-server
```

### List available tools:
```bash
echo '{"jsonrpc":"2.0","id":2,"method":"tools/list"}' | ./bin/mcp-server
```

### Call the echo tool:
```bash
echo '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"echo","arguments":{"message":"Hello!"}}}' | ./bin/mcp-server
```

## Example 3: Using with VSCode

If you have an MCP extension for VSCode:

1. Open VSCode settings (JSON)
2. Add:
```json
{
  "mcp.servers": {
    "gittisak-go": {
      "command": "/absolute/path/to/gittisak-go/bin/mcp-server",
      "args": []
    }
  }
}
```
3. Reload VSCode

## Example 4: Automated Testing

Run the test suite:
```bash
./test.sh
```

Expected output:
```
======================================
MCP Server Test Suite
======================================

Running tests...

Test 1: Initialize
✓ Initialize: PASSED

Test 2: List Tools
✓ List Tools: PASSED (found echo, get_time, read_file)

Test 3: Call Echo Tool
✓ Echo Tool: PASSED

Test 4: Call Get Time Tool
✓ Get Time Tool: PASSED

Test 5: Call Read File Tool
✓ Read File Tool: PASSED

======================================
Test Suite Complete
======================================
```

## Example 5: Creating a Custom Tool

Add a new tool to `pkg/tools/tools.go`:

```go
// CalculateTool performs basic calculations
func CalculateTool(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
    operation, _ := arguments["operation"].(string)
    a, _ := arguments["a"].(float64)
    b, _ := arguments["b"].(float64)
    
    var result float64
    switch operation {
    case "add":
        result = a + b
    case "subtract":
        result = a - b
    case "multiply":
        result = a * b
    case "divide":
        if b == 0 {
            return &mcp.CallToolResult{
                Content: []mcp.Content{{Type: "text", Text: "Error: Division by zero"}},
                IsError: true,
            }, nil
        }
        result = a / b
    default:
        return &mcp.CallToolResult{
            Content: []mcp.Content{{Type: "text", Text: "Error: Invalid operation"}},
            IsError: true,
        }, nil
    }
    
    return &mcp.CallToolResult{
        Content: []mcp.Content{
            {Type: "text", Text: fmt.Sprintf("Result: %.2f", result)},
        },
    }, nil
}
```

Register it in `cmd/mcp-server/main.go`:

```go
server.RegisterTool(
    "calculate",
    "Performs basic arithmetic calculations",
    mcp.InputSchema{
        Type: "object",
        Properties: map[string]interface{}{
            "operation": map[string]interface{}{
                "type":        "string",
                "description": "Operation to perform: add, subtract, multiply, divide",
            },
            "a": map[string]interface{}{
                "type":        "number",
                "description": "First number",
            },
            "b": map[string]interface{}{
                "type":        "number",
                "description": "Second number",
            },
        },
        Required: []string{"operation", "a", "b"},
    },
    tools.CalculateTool,
)
```

Rebuild and test:
```bash
make build
echo '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"calculate","arguments":{"operation":"add","a":5,"b":3}}}' | ./bin/mcp-server
```

## Troubleshooting Examples

### Issue: Server binary not found
```bash
# Solution: Build the server first
make build
```

### Issue: Permission denied
```bash
# Solution: Make binary executable
chmod +x bin/mcp-server
```

### Issue: Claude Desktop doesn't see the server
```bash
# Solution: Use absolute path in config
# Instead of: "./bin/mcp-server"
# Use: "/home/username/gittisak-go/bin/mcp-server"
```

## Next Steps

- Read [README-MCP.md](README-MCP.md) for comprehensive documentation
- Explore the [MCP Specification](https://modelcontextprotocol.io)
- Join the [MCP Community](https://modelcontextprotocol.io/community)
