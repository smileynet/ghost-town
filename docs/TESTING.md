# Testing Conventions

This document outlines testing conventions and best practices for ghost-town development.

## Framework

**Primary Testing Framework:** testify

- Assertions: `github.com/stretchr/testify/assert`
- Requirements: `github.com/stretchr/testify/require`
- Mocks: `github.com/stretchr/testify/mock`
- Test Suites: `github.com/stretchr/testify/suite`

## Installation

```bash
go get github.com/stretchr/testify
```

## File Organization

```
internal/
  package/
    package_test.go        # Unit tests for package
  testing/
    helpers.go             # Shared testing utilities
tests/
  integration_test.go      # Integration tests
```

## Test Structure

### Basic Unit Test

```go
package package

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestFunction(t *testing.T) {
    result := Function()
    assert.Equal(t, "expected", result)
    require.NotNil(t, result)
}
```

### Table-Driven Tests

```go
func TestCalculate(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"zero", 0, 0},
        {"positive", 5, 10},
        {"negative", -5, -10},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Calculate(tt.input)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

### Helper Functions

Use helper functions from `internal/testing`:

```go
import (
    "github.com/smileynet/ghost-town/internal/testing"
)

func TestWithTempDir(t *testing.T) {
    dir := testing.TempDir(t)
    // dir is automatically cleaned up after test
}
```

Available helpers:
- `TempDir(t)` - Create temp directory
- `TempFile(t, content)` - Create temp file
- `AssertEqual(t, expected, actual)` - Assert equality
- `RequireEqual(t, expected, actual)` - Require equality
- `AssertNoError(t, err)` - Assert no error
- `RequireNoError(t, err)` - Require no error
- `SkipIntegration(t)` - Skip if SKIP_INTEGRATION=true
- `SkipShort(t)` - Skip in short mode
- `RunTableTest(t, tests, testFunc)` - Run table-driven tests

## Assertions

### Use `assert` for non-fatal failures

```go
import "github.com/stretchr/testify/assert"

func TestFunction(t *testing.T) {
    result := Function()
    assert.Equal(t, "expected", result)
    assert.NotNil(t, result)
    assert.Contains(t, result, "substring")
}
```

### Use `require` for fatal failures

```go
import "github.com/stretchr/testify/require"

func TestFunction(t *testing.T) {
    file := os.Open("file.txt")
    require.NoError(t, err) // Stop test on error
    defer file.Close()

    content := ReadFile(file)
    require.NotNil(t, content) // Stop test if nil
}
```

## Mocking

### Generate Mocks

```bash
go install github.com/golang/mock/mockgen@latest
mockgen -source=internal/interface.go -destination=internal/mock/mock_interface.go
```

### Use Mocks

```go
import (
    "github.com/stretchr/testify/mock"
    "github.com/smileyney/ghost-town/internal/mock"
)

func TestWithMock(t *testing.T) {
    m := new(mock.MyInterface)
    m.On("Method", arg1, arg2).Return(result, nil)

    result, err := UseInterface(m)
    require.NoError(t, err)

    m.AssertExpectations(t)
}
```

## Integration Tests

Integration tests go in `tests/` directory:

```go
package integration_test

import (
    "os"
    "os/exec"
    "testing"
    "github.com/smileynet/ghost-town/internal/testing"
)

func TestBuildBinary(t *testing.T) {
    testing.SkipIntegration(t)

    cmd := exec.Command("go", "build", "-o", "/tmp/ghost-town-test", "./cmd/ghost-town")
    output, err := cmd.CombinedOutput()
    require.NoError(t, err, "Failed to build: %s", output)

    defer os.Remove("/tmp/ghost-town-test")
}
```

### Running Integration Tests

```bash
# Run all tests
make test

# Run only unit tests (skip integration)
SKIP_INTEGRATION=true make test

# Run only integration tests
make test-integration

# Run short tests
make test-short
```

## Test Coverage

### Generate Coverage Report

```bash
make test-coverage
```

This generates `coverage.out` and `coverage.html`.

### Coverage Threshold

- **Minimum target:** 80%
- **Good target:** 90%
- **Excellent target:** 95%

### Checking Coverage

```bash
go test -cover ./...
```

## Best Practices

1. **Use t.Helper()** - Mark helper functions
2. **Table-driven tests** - For multiple test cases
3. **Describe failures** - Use descriptive error messages
4. **Clean up resources** - Use defer or testing helpers
5. **Test edge cases** - Nil, empty, negative, overflow
6. **Keep tests fast** - Unit tests should run in <100ms
7. **Skip integration tests** - Use `SKIP_INTEGRATION=true` for CI
8. **Use require for setup** - Fail fast on setup errors
9. **Use assert for verification** - Continue test on assertion failures

## Naming Conventions

- Test functions: `Test<FunctionName>`
- Benchmark functions: `Benchmark<FunctionName>`
- Example functions: `Example<FunctionName>`
- Table test cases: Descriptive names like {"zero", "positive", "negative"}

## Makefile Targets

```makefile
test: ## Run tests
	go test -v -race -cover ./...

test-short: ## Run short tests
	go test -short -v ./...

test-coverage: ## Run tests with coverage
	go test -cover -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

test-integration: ## Run integration tests
	go test -v ./tests/
```

## CI/CD Integration

```yaml
# .github/workflows/ci.yml
- name: Run tests
  run: make test

- name: Check coverage
  run: make test-coverage

- name: Upload coverage
  uses: codecov/codecov-action@v3
  with:
    file: ./coverage.out
```

---

**Last Updated:** 2026-01-15
**Task:** bd-7q7ek
**Framework:** testify v1.11.1
