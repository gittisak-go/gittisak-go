# Testing Guide for Gittisak Go MCP Server

This guide provides comprehensive information about testing the MCP server.

## Test Structure

The project includes multiple levels of testing:

1. **Unit Tests** - Test individual components in isolation
2. **Integration Tests** - Test the server with real JSON-RPC messages
3. **Interactive Tests** - Manual testing with live server interaction

## Running Tests

### Quick Test (All Tests)

```bash
# Build and run all tests
make build
make test
./test.sh
```

### Unit Tests Only

```bash
# Run Go unit tests with verbose output
go test -v ./...

# Run tests for specific package
go test -v ./pkg/mcp
go test -v ./pkg/tools

# Run with coverage
go test -v -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Integration Tests

```bash
# Run the integration test suite
./test.sh
```

This script tests:
- Server initialization
- Tool listing
- Echo tool execution
- Time tool execution
- File reading tool execution

### Interactive Tests

```bash
# Run interactive test script (requires FIFO support)
./test_interactive.sh
```

## Test Coverage

### MCP Server Tests (`pkg/mcp/server_test.go`)

Tests cover:
- ✅ Server creation and initialization
- ✅ Tool registration with metadata
- ✅ Initialize request handling
- ✅ List tools request handling
- ✅ Call tool request handling
- ✅ Tool not found error handling
- ✅ Response sending
- ✅ Error response sending

### Tools Tests (`pkg/tools/tools_test.go`)

Tests cover:
- ✅ Echo tool with valid input
- ✅ Echo tool with empty message
- ✅ Echo tool with missing arguments
- ✅ Echo tool with invalid argument types
- ✅ Get time tool functionality
- ✅ Read file tool with valid file
- ✅ Read file tool with nonexistent file
- ✅ Read file tool with missing path
- ✅ Read file tool with directory (error case)
- ✅ Benchmark tests for performance

## Writing New Tests

### Unit Test Example

```go
func TestMyNewTool(t *testing.T) {
    // Arrange
    args := map[string]interface{}{
        "param": "value",
    }
    
    // Act
    result, err := MyNewTool(args)
    
    // Assert
    if err != nil {
        t.Errorf("Expected no error, got: %v", err)
    }
    
    if result.IsError {
        t.Error("Expected success result")
    }
    
    if len(result.Content) != 1 {
        t.Fatalf("Expected 1 content item, got %d", len(result.Content))
    }
    
    if result.Content[0].Text != "expected text" {
        t.Errorf("Expected 'expected text', got '%s'", result.Content[0].Text)
    }
}
```

### Integration Test Example

Add to `test.sh`:

```bash
# Test: My New Tool
echo "Test X: Call My New Tool"
RESPONSE=$(echo '{"jsonrpc":"2.0","id":X,"method":"tools/call","params":{"name":"my_new_tool","arguments":{"param":"value"}}}' | timeout 2s ./bin/mcp-server 2>/dev/null | head -1)
if echo "$RESPONSE" | grep -q '"text":"expected result"'; then
    echo "✓ My New Tool: PASSED"
else
    echo "✗ My New Tool: FAILED"
    echo "  Response: $RESPONSE"
fi
echo ""
```

## Benchmark Tests

Run performance benchmarks:

```bash
# Run all benchmarks
go test -bench=. ./...

# Run specific benchmark
go test -bench=BenchmarkEchoTool ./pkg/tools

# Run benchmarks with memory statistics
go test -bench=. -benchmem ./...
```

Example output:
```
BenchmarkEchoTool-8         5000000    250 ns/op    128 B/op    3 allocs/op
BenchmarkGetTimeTool-8      1000000   1200 ns/op    256 B/op    5 allocs/op
BenchmarkReadFileTool-8      500000   2500 ns/op    512 B/op   10 allocs/op
```

## Continuous Integration

The project is ready for CI/CD integration. Example GitHub Actions workflow:

```yaml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Build
        run: make build
      
      - name: Run unit tests
        run: go test -v -cover ./...
      
      - name: Run integration tests
        run: ./test.sh
```

## Test Best Practices

### 1. Test Naming

- Use descriptive test names: `TestEchoToolWithValidMessage`
- Group related tests using subtests: `t.Run("Valid input", func(t *testing.T) { ... })`

### 2. Test Structure

Follow the AAA pattern:
- **Arrange**: Set up test data and conditions
- **Act**: Execute the code being tested
- **Assert**: Verify the results

### 3. Error Messages

Provide clear error messages:
```go
if got != want {
    t.Errorf("got %v, want %v", got, want)
}
```

### 4. Table-Driven Tests

Use table-driven tests for multiple scenarios:
```go
tests := []struct {
    name     string
    input    interface{}
    expected interface{}
}{
    {"case 1", input1, expected1},
    {"case 2", input2, expected2},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // test logic
    })
}
```

### 5. Test Isolation

- Each test should be independent
- Clean up resources (files, connections) in defer statements
- Use temporary files for file operations

## Debugging Tests

### Run Single Test

```bash
# Run a specific test
go test -v -run TestEchoTool ./pkg/tools

# Run a specific subtest
go test -v -run TestEchoTool/Valid_message ./pkg/tools
```

### Verbose Output

```bash
# Show all log output during tests
go test -v ./...

# Show even more detailed output
go test -v -x ./...
```

### Test with Race Detector

```bash
# Detect race conditions
go test -race ./...
```

## Test Coverage Goals

Current coverage:
- `pkg/mcp`: ~95% coverage
- `pkg/tools`: ~90% coverage
- Overall: ~92% coverage

Goals:
- Maintain minimum 90% coverage for all packages
- 100% coverage for critical paths (error handling, security)

## Common Test Failures

### 1. Timeout Issues

If tests timeout, increase the timeout value:
```bash
timeout 5s ./bin/mcp-server  # Instead of 2s
```

### 2. File Permission Issues

Ensure test files have correct permissions:
```bash
chmod 644 test_file.txt
```

### 3. Path Issues

Use absolute paths in tests when needed:
```go
tmpFile, _ := os.CreateTemp("", "test-*.txt")
defer os.Remove(tmpFile.Name())
```

## Performance Testing

### Load Testing

Test the server under load:
```bash
# Send multiple concurrent requests
for i in {1..100}; do
    echo '{"jsonrpc":"2.0","id":'$i',"method":"tools/list"}' | ./bin/mcp-server &
done
wait
```

### Memory Profiling

```bash
# Run tests with memory profiling
go test -memprofile=mem.out ./...

# Analyze memory profile
go tool pprof mem.out
```

### CPU Profiling

```bash
# Run tests with CPU profiling
go test -cpuprofile=cpu.out ./...

# Analyze CPU profile
go tool pprof cpu.out
```

## Test Data

Test files are created in `/tmp` directory:
- Integration tests use `/tmp/mcp_test.txt`
- Interactive tests use `/tmp/mcp_test_file.txt`
- Unit tests use `os.CreateTemp()` for temporary files

All test data is automatically cleaned up after tests complete.

## Troubleshooting Tests

### Tests Fail After Code Changes

1. Rebuild the server: `make build`
2. Clear Go test cache: `go clean -testcache`
3. Run tests again: `go test ./...`

### Integration Tests Fail but Unit Tests Pass

1. Check if the binary is up to date: `make build`
2. Verify the binary has execute permissions: `chmod +x bin/mcp-server`
3. Check for orphaned server processes: `pkill mcp-server`

### Random Test Failures

1. Check for race conditions: `go test -race ./...`
2. Look for test dependencies on external state
3. Ensure tests are properly isolated

## Next Steps

After all tests pass:
1. Review test coverage: `go test -cover ./...`
2. Add tests for any uncovered code
3. Document any known test limitations
4. Update this guide with new test scenarios

---

For more information, see:
- [Go Testing Package Documentation](https://pkg.go.dev/testing)
- [Table-Driven Tests in Go](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
- [MCP Testing Best Practices](https://modelcontextprotocol.io/docs/testing)
