# Architecture

This document describes the architecture of ghost-town, a kiro-cli tool based on Gas Town architecture.

## Overview

Ghost-town is a command-line interface tool that implements the Gas Town architecture patterns for task execution and coordination.

## Design Principles

1. **Modularity**: Clean separation between CLI, business logic, and utilities
2. **Simplicity**: Easy to understand, test, and extend
3. **Performance**: Efficient resource usage and fast startup
4. **Compatibility**: Works with Gas Town components and workflows

## Project Structure

```
ghost-town/
├── cmd/                  # CLI entry points
│   └── ghost-town/      # Main CLI application
├── internal/            # Private packages (not importable by external projects)
│   ├── cli/            # CLI command logic
│   ├── config/         # Configuration management
│   └── version/        # Version information
├── pkg/                 # Public APIs (importable by external projects)
│   ├── errors/         # Common error types
│   └── logger/         # Logging interfaces
├── docs/                # Documentation
└── tests/               # Integration tests
```

## Core Components

### CLI Layer (`cmd/`)

The CLI layer handles:
- Command parsing and validation
- User interaction and feedback
- Command dispatching to internal logic

### Business Logic (`internal/`)

Private packages containing core functionality:
- `cli/`: Command execution logic
- `config/`: Configuration loading and validation
- `version/`: Version and build information

### Public APIs (`pkg/`)

Reusable packages for external integration:
- `errors/`: Standard error types and handling
- `logger/`: Logging interfaces and implementations

## Protomolecule Pattern

Ghost-town uses the **protomolecule** pattern for autonomous task execution, as documented in [docs/PROTO_EXECUTION.md](./PROTO_EXECUTION.md).

Key concepts:
- **Worker Agent**: Executes specific tasks
- **Validator Agent**: Validates and provides feedback
- **Iteration**: Loop until approval or max attempts

## Task Execution Flow

```
User Input
    ↓
CLI Parsing
    ↓
Command Dispatch
    ↓
Worker Agent Execution
    ↓
Validator Review
    ↓
[Approved] → Complete
   ↕ [Needs Changes]
    ↓
Worker Agent (Retry)
```

## Configuration

Configuration is loaded from:
1. Command-line flags (highest priority)
2. Environment variables
3. Configuration file
4. Defaults (lowest priority)

## Logging

Ghost-town uses structured logging with multiple levels:
- Debug: Detailed diagnostic information
- Info: General informational messages
- Error: Error conditions

## Error Handling

Errors are categorized using the `errors` package:
- `ErrNotFound`: Resource not found
- `ErrInvalid`: Invalid input or configuration
- Custom errors for specific scenarios

## Dependencies

### Go Libraries
- Standard library (no external dependencies initially)

### External Tools
- `bd`: Issue tracking (beads)
- `gh`: GitHub CLI for GitHub operations

## Extension Points

1. **New Commands**: Add to `internal/cli/`
2. **New Packages**: Add to `internal/` (private) or `pkg/` (public)
3. **Protomolecules**: Document in `docs/` as `*.formula.toml`

## Performance Considerations

- Lazy initialization of heavy resources
- Efficient string handling (avoid unnecessary allocations)
- Goroutine pools for concurrent operations
- Context cancellation for long-running operations

## Security Considerations

- Validate all user input
- Sanitize file paths
- Use environment variables for secrets (never commit)
- Implement proper error handling (don't leak sensitive info)

## Future Enhancements

- Plugin system for extensions
- REST API for remote operations
- Metrics and observability
- Enhanced testing framework
