#!/bin/bash
# Interactive MCP Server Test Script
# This script demonstrates how to interact with the MCP server

set -e

echo "MCP Server Interactive Test"
echo "============================"
echo ""

# Check if server binary exists
if [ ! -f "bin/mcp-server" ]; then
    echo "Building server..."
    make build
    echo ""
fi

# Create a test file for read_file tool
TEST_FILE="/tmp/mcp_test_file.txt"
echo "This is a test file for the MCP server read_file tool." > "$TEST_FILE"
echo "Created test file: $TEST_FILE"
echo ""

# Start the server in the background
echo "Starting MCP server..."
mkfifo /tmp/mcp_in /tmp/mcp_out 2>/dev/null || true
./bin/mcp-server < /tmp/mcp_in > /tmp/mcp_out &
SERVER_PID=$!
echo "Server started with PID: $SERVER_PID"
echo ""

# Give the server a moment to start
sleep 0.5

# Function to send request and get response
send_request() {
    local request="$1"
    local description="$2"
    
    echo "Test: $description"
    echo "Request: $request"
    echo "$request" > /tmp/mcp_in
    sleep 0.2
    
    # Read response (non-blocking)
    timeout 1s head -n 1 /tmp/mcp_out 2>/dev/null || echo "No response"
    echo ""
}

# Test 1: Initialize
send_request \
    '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test-script","version":"1.0.0"}}}' \
    "Initialize Connection"

# Test 2: List Tools
send_request \
    '{"jsonrpc":"2.0","id":2,"method":"tools/list"}' \
    "List Available Tools"

# Test 3: Echo Tool
send_request \
    '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"echo","arguments":{"message":"Hello from test script!"}}}' \
    "Call Echo Tool"

# Test 4: Get Time Tool
send_request \
    '{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"get_time","arguments":{}}}' \
    "Call Get Time Tool"

# Test 5: Read File Tool
send_request \
    "{\"jsonrpc\":\"2.0\",\"id\":5,\"method\":\"tools/call\",\"params\":{\"name\":\"read_file\",\"arguments\":{\"path\":\"$TEST_FILE\"}}}" \
    "Call Read File Tool"

# Cleanup
echo "Cleaning up..."
kill $SERVER_PID 2>/dev/null || true
rm -f /tmp/mcp_in /tmp/mcp_out "$TEST_FILE"
echo "Tests completed!"
