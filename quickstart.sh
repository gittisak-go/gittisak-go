#!/bin/bash
# Quick start script for MCP server

set -e

echo "Gittisak Go MCP Server - Quick Start"
echo "===================================="
echo ""

# Build the server
echo "Building MCP server..."
make build

echo ""
echo "✓ Build complete!"
echo ""
echo "The MCP server is now ready to use."
echo ""
echo "Next steps:"
echo "1. Configure your MCP client (Claude Desktop, VSCode, etc.)"
echo "2. Add the following configuration:"
echo ""
echo "   {\"
echo "     \"command\": \"$(pwd)/bin/mcp-server\","
echo "     \"args\": [],"
echo "     \"env\": {}"
echo "   }"
echo ""
echo "3. For Claude Desktop, edit:"
echo "   - macOS: ~/Library/Application Support/Claude/claude_desktop_config.json"
echo "   - Windows: %APPDATA%\\Claude\\claude_desktop_config.json"
echo "   - Linux: ~/.config/Claude/claude_desktop_config.json"
echo ""
echo "4. Restart your MCP client"
echo ""
echo "For more details, see README-MCP.md"
