#!/bin/bash
# Simple MCP Server Test Script
# Tests the MCP server by sending JSON-RPC messages

set -e

echo "======================================"
echo "MCP Server Test Suite"
echo "======================================"
echo ""

# Build server if needed
if [ ! -f "bin/mcp-server" ]; then
    echo "Building server..."
    make build
    echo ""
fi

# Create test file
TEST_FILE="/tmp/mcp_test.txt"
echo "Sample content for MCP server test" > "$TEST_FILE"

echo "Running tests..."
echo ""

# Test 1: Initialize
echo "Test 1: Initialize"
RESPONSE=$(echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}}}' | timeout 2s ./bin/mcp-server 2>/dev/null | head -1)
if echo "$RESPONSE" | grep -q '"protocolVersion":"2024-11-05"'; then
    echo "✓ Initialize: PASSED"
else
    echo "✗ Initialize: FAILED"
    echo "  Response: $RESPONSE"
fi
echo ""

# Test 2: List Tools
echo "Test 2: List Tools"
RESPONSE=$(echo '{"jsonrpc":"2.0","id":2,"method":"tools/list"}' | timeout 2s ./bin/mcp-server 2>/dev/null | head -1)
if echo "$RESPONSE" | grep -q '"name":"echo"' && echo "$RESPONSE" | grep -q '"name":"get_time"' && echo "$RESPONSE" | grep -q '"name":"read_file"'; then
    echo "✓ List Tools: PASSED (found echo, get_time, read_file)"
else
    echo "✗ List Tools: FAILED"
    echo "  Response: $RESPONSE"
fi
echo ""

# Test 3: Echo Tool
echo "Test 3: Call Echo Tool"
RESPONSE=$(echo '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"echo","arguments":{"message":"Test message"}}}' | timeout 2s ./bin/mcp-server 2>/dev/null | head -1)
if echo "$RESPONSE" | grep -q '"text":"Echo: Test message"'; then
    echo "✓ Echo Tool: PASSED"
else
    echo "✗ Echo Tool: FAILED"
    echo "  Response: $RESPONSE"
fi
echo ""

# Test 4: Get Time Tool
echo "Test 4: Call Get Time Tool"
RESPONSE=$(echo '{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"get_time","arguments":{}}}' | timeout 2s ./bin/mcp-server 2>/dev/null | head -1)
if echo "$RESPONSE" | grep -q '"text":"Current server time:'; then
    echo "✓ Get Time Tool: PASSED"
else
    echo "✗ Get Time Tool: FAILED"
    echo "  Response: $RESPONSE"
fi
echo ""

# Test 5: Read File Tool
echo "Test 5: Call Read File Tool"
RESPONSE=$(echo "{\"jsonrpc\":\"2.0\",\"id\":5,\"method\":\"tools/call\",\"params\":{\"name\":\"read_file\",\"arguments\":{\"path\":\"$TEST_FILE\"}}}" | timeout 2s ./bin/mcp-server 2>/dev/null | head -1)
if echo "$RESPONSE" | grep -q '"text":"Sample content'; then
    echo "✓ Read File Tool: PASSED"
else
    echo "✗ Read File Tool: FAILED"
    echo "  Response: $RESPONSE"
fi
echo ""

# Cleanup
rm -f "$TEST_FILE"

echo "======================================"
echo "Test Suite Complete"
echo "======================================"
