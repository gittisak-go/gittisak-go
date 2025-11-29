#!/bin/bash
# MCP Setup Verification Script
# สคริปต์ตรวจสอบการตั้งค่า MCP
#
# This script helps users verify their MCP setup and find where their project is located.
# สคริปต์นี้ช่วยผู้ใช้ตรวจสอบการตั้งค่า MCP และหาตำแหน่งโครงการ

set -e

# Get script directory
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

# Function to print Claude Desktop config
print_claude_config() {
    cat <<EOF
{
  "mcpServers": {
    "gittisak-go": {
      "command": "$SCRIPT_DIR/bin/mcp-server",
      "args": [],
      "env": {}
    }
  }
}
EOF
}

# Function to print VSCode config
print_vscode_config() {
    cat <<EOF
{
  "mcp": {
    "servers": {
      "gittisak-go": {
        "command": "$SCRIPT_DIR/bin/mcp-server",
        "args": [],
        "env": {}
      }
    }
  }
}
EOF
}

# Function to find Claude Desktop config
find_claude_config() {
    local config_paths=(
        "$HOME/Library/Application Support/Claude/claude_desktop_config.json"
        "$HOME/.config/Claude/claude_desktop_config.json"
    )
    
    for path in "${config_paths[@]}"; do
        if [ -f "$path" ]; then
            echo "$path"
            return 0
        fi
    done
    return 1
}

echo "=================================================="
echo "MCP Setup Verification / ตรวจสอบการตั้งค่า MCP"
echo "=================================================="
echo ""

echo "📁 Project Location / ตำแหน่งโครงการ:"
echo "   $SCRIPT_DIR"
echo ""

# Check if Go is installed
echo "🔍 Checking Go installation / ตรวจสอบการติดตั้ง Go..."
if command -v go &> /dev/null; then
    GO_VERSION=$(go version)
    echo "   ✓ Go is installed: $GO_VERSION"
else
    echo "   ✗ Go is NOT installed"
    echo "   Please install Go 1.20+ from https://go.dev/dl/"
    exit 1
fi
echo ""

# Check if binary exists
echo "🔍 Checking MCP server binary / ตรวจสอบไฟล์โปรแกรม MCP..."
if [ -f "bin/mcp-server" ]; then
    echo "   ✓ Binary exists at: $SCRIPT_DIR/bin/mcp-server"
    
    # Check if executable
    if [ -x "bin/mcp-server" ]; then
        echo "   ✓ Binary is executable"
    else
        echo "   ✗ Binary is NOT executable"
        echo "   Run: chmod +x bin/mcp-server"
    fi
else
    echo "   ✗ Binary not found"
    echo "   Building server... / กำลังสร้างเซิร์ฟเวอร์..."
    make build
    if [ -f "bin/mcp-server" ]; then
        echo "   ✓ Build successful!"
    else
        echo "   ✗ Build failed"
        exit 1
    fi
fi
echo ""

# Test the server
echo "🧪 Testing MCP server / ทดสอบเซิร์ฟเวอร์ MCP..."
# Create test request payload (must be single line for JSON-RPC)
TEST_REQUEST='{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"verify","version":"1.0"}}}'

RESPONSE=$(echo "$TEST_REQUEST" | timeout 2s ./bin/mcp-server 2>/dev/null | head -1)
if echo "$RESPONSE" | grep -q '"protocolVersion"'; then
    echo "   ✓ Server responds correctly"
else
    echo "   ✗ Server test failed"
    echo "   Response: $RESPONSE"
fi
echo ""

# Check Claude Desktop config
echo "🔍 Checking Claude Desktop configuration..."
if CONFIG_PATH=$(find_claude_config); then
    echo "   ✓ Config found at: $CONFIG_PATH"
    
    # Check if our server is configured
    if grep -q "gittisak-go" "$CONFIG_PATH" 2>/dev/null; then
        echo "   ✓ gittisak-go server is configured"
    else
        echo "   ⚠ gittisak-go server NOT found in config"
        echo ""
        echo "   Add this to your config:"
        print_claude_config | sed 's/^/   /'
    fi
else
    echo "   ⚠ Claude Desktop config not found"
    echo "   Expected locations:"
    echo "     - macOS: ~/Library/Application Support/Claude/claude_desktop_config.json"
    echo "     - Linux: ~/.config/Claude/claude_desktop_config.json"
    echo "     - Windows: %APPDATA%\\Claude\\claude_desktop_config.json"
fi
echo ""

# Summary
echo "=================================================="
echo "Summary / สรุป"
echo "=================================================="
echo ""
echo "📍 Your MCP server is located at:"
echo "   $SCRIPT_DIR/bin/mcp-server"
echo ""
echo "📋 Add this configuration to your MCP client:"
echo ""
echo "For Claude Desktop (claude_desktop_config.json):"
print_claude_config
echo ""
echo "For VSCode (settings.json):"
print_vscode_config
echo ""
echo "=================================================="
echo "Need help? / ต้องการความช่วยเหลือ?"
echo "  - GitHub Issues: https://github.com/gittisak-go/gittisak-go/issues"
echo "  - Documentation: README-MCP.md"
echo "=================================================="
