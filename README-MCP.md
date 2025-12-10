# Gittisak Go MCP Server

A Model Context Protocol (MCP) server implementation in Go that enables AI applications like Claude Desktop, VSCode, Perplexity, and Figma to connect to local tools and resources.

## Features

- **stdio Transport**: Standard MCP communication via stdin/stdout
- **Built-in Tools**:
  - `echo`: Echo back messages
  - `get_time`: Get current server time
  - `read_file`: Read file contents from the filesystem
- **Extensible**: Easy to add custom tools
- **Zero Dependencies**: Pure Go implementation

## Prerequisites

- Go 1.20 or higher
- Compatible MCP client (Claude Desktop, VSCode with MCP extension, etc.)

## Installation

### Building from Source

```bash
# Clone the repository
git clone https://github.com/gittisak-go/gittisak-go.git
cd gittisak-go

# Build the server
go build -o bin/mcp-server ./cmd/mcp-server

# The binary will be available at bin/mcp-server
```

### Quick Start

```bash
# Run the server (for testing)
./bin/mcp-server

# The server will start and listen on stdin/stdout
# You can interact with it by sending JSON-RPC messages
```

## Configuration

### Claude Desktop

1. Locate your Claude Desktop config file:
   - **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
   - **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`
   - **Linux**: `~/.config/Claude/claude_desktop_config.json`

2. Add the server configuration:

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

3. Restart Claude Desktop

### VSCode

If you're using VSCode with an MCP extension:

1. Open VSCode settings (JSON)
2. Add the MCP server configuration:

```json
{
  "mcp.servers": {
    "gittisak-go": {
      "command": "/absolute/path/to/gittisak-go/bin/mcp-server",
      "args": [],
      "env": {}
    }
  }
}
```

### Other MCP Clients

For other MCP clients (Perplexity, Figma plugins, etc.), refer to their specific documentation for MCP server configuration. Generally, you'll need to provide:

- **Command**: Path to the `mcp-server` binary
- **Args**: Empty array `[]`
- **Transport**: stdio (default for most clients)

## Available Tools

### echo

Echoes back the input message.

**Parameters:**
- `message` (string, required): The message to echo

**Example:**
```json
{
  "name": "echo",
  "arguments": {
    "message": "Hello, MCP!"
  }
}
```

### get_time

Returns the current server time in RFC3339 format.

**Parameters:** None

**Example:**
```json
{
  "name": "get_time",
  "arguments": {}
}
```

### read_file

Reads and returns the content of a file from the local filesystem.

**Parameters:**
- `path` (string, required): Absolute or relative path to the file

**Example:**
```json
{
  "name": "read_file",
  "arguments": {
    "path": "/path/to/file.txt"
  }
}
```

## Development

### Project Structure

```
gittisak-go/
├── cmd/
│   └── mcp-server/       # Main application entry point
│       └── main.go
├── pkg/
│   ├── mcp/              # MCP protocol implementation
│   │   ├── types.go      # Protocol types and structures
│   │   └── server.go     # Server implementation
│   └── tools/            # Tool implementations
│       └── tools.go
├── examples/             # Configuration examples
│   ├── claude-desktop-config.json
│   └── vscode-config.json
├── bin/                  # Built binaries (created after build)
├── go.mod
└── README-MCP.md
```

### Adding Custom Tools

1. Create your tool handler function in `pkg/tools/tools.go`:

```go
func MyCustomTool(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
    // Your tool implementation
    return &mcp.CallToolResult{
        Content: []mcp.Content{
            {
                Type: "text",
                Text: "Result from my custom tool",
            },
        },
    }, nil
}
```

2. Register the tool in `cmd/mcp-server/main.go`:

```go
server.RegisterTool(
    "my_custom_tool",
    "Description of my custom tool",
    mcp.InputSchema{
        Type: "object",
        Properties: map[string]interface{}{
            "param1": map[string]interface{}{
                "type":        "string",
                "description": "Description of param1",
            },
        },
        Required: []string{"param1"},
    },
    tools.MyCustomTool,
)
```

3. Rebuild the server:

```bash
go build -o bin/mcp-server ./cmd/mcp-server
```

### Testing

You can test the MCP server using the MCP Inspector or by sending JSON-RPC messages manually:

```bash
# Example initialize request
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}}}' | ./bin/mcp-server

# Example list tools request
echo '{"jsonrpc":"2.0","id":2,"method":"tools/list"}' | ./bin/mcp-server

# Example call tool request
echo '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"echo","arguments":{"message":"Hello"}}}' | ./bin/mcp-server
```

## Protocol Details

This server implements the Model Context Protocol (MCP) version 2024-11-05, which is based on JSON-RPC 2.0.

### Supported Methods

- `initialize`: Initialize the MCP connection
- `tools/list`: List available tools
- `tools/call`: Execute a tool
- `notifications/initialized`: Client initialization notification

### Message Format

All messages follow JSON-RPC 2.0 format:

**Request:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "method_name",
  "params": {}
}
```

**Response:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {}
}
```

## Troubleshooting

### Cannot access MCP / Where is my project?
### ไม่สามารถเข้าถึง MCP ได้ / โครงการของฉันอยู่ที่ไหน?

If you cannot access MCP or are unsure where your project was created, follow these steps:

หากคุณไม่สามารถเข้าถึง MCP ได้ หรือไม่แน่ใจว่าสร้างโครงการไว้ที่ไหน ให้ทำตามขั้นตอนเหล่านี้:

#### Step 1: Verify Installation Location / ตรวจสอบตำแหน่งการติดตั้ง

Run the verification script to check your MCP setup:
```bash
# Navigate to the project directory
cd /path/to/gittisak-go

# Run the verification script
./verify-mcp-setup.sh
```

Or manually check these locations:
```bash
# Find where you cloned this repository
find ~ -name "gittisak-go" -type d 2>/dev/null

# Check if the MCP server binary exists
ls -la bin/mcp-server
```

#### Step 2: Check Client Configuration / ตรวจสอบการตั้งค่าไคลเอนต์

Check your MCP client configuration file for existing settings:

**Claude Desktop configuration locations:**
- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`
- **Linux**: `~/.config/Claude/claude_desktop_config.json`

**VSCode MCP configuration:**
- Check VSCode settings (JSON) for `mcp.servers` section

```bash
# On macOS, view your Claude Desktop config:
cat ~/Library/Application\ Support/Claude/claude_desktop_config.json

# On Linux, view your Claude Desktop config:
cat ~/.config/Claude/claude_desktop_config.json

# If the file doesn't exist, Claude Desktop may not be installed or configured
```

#### Step 3: Rebuild and Test / สร้างใหม่และทดสอบ

If you found your project but it's not working:
```bash
# Clean and rebuild
make clean
make build

# Test the server
./test.sh

# Verify the binary works
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}}}' | ./bin/mcp-server
```

#### Step 4: Verify Permissions / ตรวจสอบสิทธิ์

```bash
# Make sure the binary is executable
chmod +x bin/mcp-server

# Verify the path is correct in your config
# The path must be ABSOLUTE, not relative
echo "Full path to server: $(pwd)/bin/mcp-server"
```

### Server not appearing in Claude Desktop

1. Verify the path to the binary is absolute and correct
2. Check that the binary has execute permissions: `chmod +x bin/mcp-server`
3. Look at Claude Desktop logs (usually in the application data directory)
4. Restart Claude Desktop after making configuration changes

### "Permission Denied" errors

Make sure the binary is executable:
```bash
chmod +x bin/mcp-server
```

### Tool execution errors

Check that:
- All required parameters are provided
- File paths (for read_file) are accessible
- The server has appropriate permissions

### Common Issues / ปัญหาที่พบบ่อย

| Issue / ปัญหา | Solution / วิธีแก้ |
|---------------|-------------------|
| Cannot find project / หาโครงการไม่เจอ | Run `find ~ -name "gittisak-go" -type d` |
| Server not starting / เซิร์ฟเวอร์ไม่เริ่มทำงาน | Rebuild with `make clean && make build` |
| Config not found / หาการตั้งค่าไม่เจอ | Create config file at the correct location (see Step 2) |
| Path error / ข้อผิดพลาดเส้นทาง | Use absolute path: `$(pwd)/bin/mcp-server` |
| Permission error / ข้อผิดพลาดสิทธิ์ | Run `chmod +x bin/mcp-server` |

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details

## Resources

- [Model Context Protocol Specification](https://modelcontextprotocol.io)
- [MCP Documentation](https://modelcontextprotocol.io/docs)
- [Claude Desktop MCP Guide](https://support.claude.com/en/articles/10949351-getting-started-with-local-mcp-servers-on-claude-desktop)

## Support

For issues and questions:
- GitHub Issues: https://github.com/gittisak-go/gittisak-go/issues
- MCP Community: https://modelcontextprotocol.io/community
