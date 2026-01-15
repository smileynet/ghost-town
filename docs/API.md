# API Reference

This document provides reference information for ghost-town's public APIs.

## Public Packages

### pkg/errors

Common error types and handling utilities.

#### ErrNotFound

Error returned when a resource is not found.

```go
var ErrNotFound = errors.New("not found")
```

#### ErrInvalid

Error returned when input is invalid.

```go
var ErrInvalid = errors.New("invalid input")
```

### pkg/logger

Logging interfaces and implementations.

#### Logger Interface

```go
type Logger interface {
    Info(msg string)
    Debug(msg string)
    Error(msg string)
}
```

#### DefaultLogger

Simple console logger implementation.

```go
type DefaultLogger struct{}

func (l *DefaultLogger) Info(msg string)
func (l *DefaultLogger) Debug(msg string)
func (l *DefaultLogger) Error(msg string)
```

**Usage:**

```go
logger := &logger.DefaultLogger{}
logger.Info("Application started")
logger.Debug("Processing request")
logger.Error("Something went wrong")
```

## CLI Commands

### ghost-town

Main CLI entry point.

#### version

Display version and build information.

```bash
ghost-town version
```

**Output:**

```
Version: 1.0.0
Build: 2026-01-15T12:00:00Z
Commit: abc123def456
```

## Future APIs

This section is a placeholder for future API documentation.

### CLI Extensions
- Command: `ghost-town task <action>`
- Command: `ghost-town agent <action>`

### Library APIs
- Package: `pkg/task` - Task execution
- Package: `pkg/agent` - Agent coordination
- Package: `pkg/validator` - Validation logic
