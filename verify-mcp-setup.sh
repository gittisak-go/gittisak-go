#!/bin/bash
# MCP Setup Verification Script
# สคริปต์ตรวจสอบการตั้งค่า MCP
#
# This script helps users verify their MCP setup and find where their project is located.
# สคริปต์นี้ช่วยผู้ใช้ตรวจสอบการตั้งค่า MCP และหาตำแหน่งโครงการ

set -e

echo "=================================================="
echo "MCP Setup Verification / ตรวจสอบการตั้งค่า MCP"
echo "=================================================="
echo ""

# Get script directory
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

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
RESPONSE=$(echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"verify","version":"1.0"}}}' | timeout 2s ./bin/mcp-server 2>/dev/null | head -1)
if echo "$RESPONSE" | grep -q '"protocolVersion"'; then
    echo "   ✓ Server responds correctly"
else
    echo "   ✗ Server test failed"
    echo "   Response: $RESPONSE"
fi
echo ""

# Check Claude Desktop config
echo "🔍 Checking Claude Desktop configuration..."
CONFIG_FOUND=false
CONFIG_PATH=""

# macOS
if [ -f "$HOME/Library/Application Support/Claude/claude_desktop_config.json" ]; then
    CONFIG_PATH="$HOME/Library/Application Support/Claude/claude_desktop_config.json"
    CONFIG_FOUND=true
fi

# Linux
if [ -f "$HOME/.config/Claude/claude_desktop_config.json" ]; then
    CONFIG_PATH="$HOME/.config/Claude/claude_desktop_config.json"
    CONFIG_FOUND=true
fi

if [ "$CONFIG_FOUND" = true ]; then
    echo "   ✓ Config found at: $CONFIG_PATH"
    
    # Check if our server is configured
    if grep -q "gittisak-go" "$CONFIG_PATH" 2>/dev/null; then
        echo "   ✓ gittisak-go server is configured"
    else
        echo "   ⚠ gittisak-go server NOT found in config"
        echo ""
        echo "   Add this to your config:"
        echo '   {'
        echo '     "mcpServers": {'
        echo '       "gittisak-go": {'
        echo "         \"command\": \"$SCRIPT_DIR/bin/mcp-server\","
        echo '         "args": [],'
        echo '         "env": {}'
        echo '       }'
        echo '     }'
        echo '   }'
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
echo '{'
echo '  "mcpServers": {'
echo '    "gittisak-go": {'
echo "      \"command\": \"$SCRIPT_DIR/bin/mcp-server\","
echo '      "args": [],'
echo '      "env": {}'
echo '    }'
echo '  }'
echo '}'
echo ""
echo "For VSCode (settings.json):"
echo '{'
echo '  "mcp": {'
echo '    "servers": {'
echo '      "gittisak-go": {'
echo "        \"command\": \"$SCRIPT_DIR/bin/mcp-server\","
echo '        "args": [],'
echo '        "env": {}'
echo '      }'
echo '    }'
echo '  }'
echo '}'
echo ""
echo "=================================================="
echo "Need help? / ต้องการความช่วยเหลือ?"
echo "  - GitHub Issues: https://github.com/gittisak-go/gittisak-go/issues"
echo "  - Documentation: README-MCP.md"
echo "=================================================="
