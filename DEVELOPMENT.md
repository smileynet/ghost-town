# Development Guide

This guide covers setting up the development environment for ghost-town, a kiro-cli tool based on Gas Town architecture.

## Prerequisites

- **Go**: Version 1.21 or later
- **Git**: For version control
- **GitHub CLI (`gh`)**: For GitHub operations
- **Beads (`bd`)**: Issue tracking and workflow management

## Installation

### 1. Install Go

Download Go from [https://go.dev/dl/](https://go.dev/dl/) or use your system package manager.

**Verify installation:**
```bash
go version
```

**Set up Go environment (add to `~/.bashrc` or `~/.zshrc`):**
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### 2. Install GitHub CLI

```bash
# macOS
brew install gh

# Linux
curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
sudo apt update
sudo apt install gh
```

**Authenticate:**
```bash
gh auth login
```

### 3. Install Beads (Issue Tracking)

Beads is a lightweight issue tracking system. Install from source:

```bash
git clone https://github.com/steveyegge/beads.git
cd beads
go install ./cmd/bd@latest
```

**Initialize in project:**
```bash
bd init bd  # Initialize with 'bd' prefix
```

## Project Setup

### Clone Repository

```bash
git clone https://github.com/smileynet/ghost-town.git
cd ghost-town
```

### Initialize Go Module

```bash
go mod init github.com/smileynet/ghost-town
```

### Install Dependencies

```bash
go mod tidy
```

## Development Workflow

### Check for Work

```bash
bd ready
```

### Claim a Task

```bash
bd update <id> --status in_progress
```

### Make Changes

Edit files, write code, run tests.

### Run Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for specific package
go test ./internal/...
```

### Build

```bash
# Build for current platform
go build ./cmd/ghost-town

# Build for multiple platforms
make build-all
```

### Complete Task

```bash
# Close task with reason
bd close <id> --reason "Completed task description"

# Close multiple tasks
bd close <id1> <id2> <id3>

# Sync to git
bd sync
```

### Commit and Push

```bash
git status
git add <files>
git commit -m "message"
git push
```

## Project Structure

```
ghost-town/
├── cmd/                  # CLI entry points
│   └── ghost-town/      # Main CLI application
├── internal/            # Private packages
├── pkg/                 # Public APIs
├── docs/                # Documentation
├── tests/               # Integration tests
├── .beads/              # Issue tracking
├── .claude/             # Opencode AI configuration
├── Makefile             # Build automation
└── go.mod              # Go module definition
```

## Testing

### Unit Tests

```bash
# Run all unit tests
go test -v ./...

# Run tests with race detection
go test -race ./...

# Run specific test
go test -v -run TestFunctionName ./internal/package
```

### Integration Tests

```bash
# Run integration tests
go test -v ./tests/...
```

### Linting

```bash
# Run gofmt
gofmt -l .

# Run golint
golint ./...

# Run staticcheck
staticcheck ./...
```

## Code Style

- Follow Go's standard formatting: `gofmt -w .`
- Write clear, descriptive variable names
- Add comments for exported functions
- Keep functions focused and small
- Write tests alongside code

## Making Releases

1. Update version in `internal/version/version.go`
2. Tag the commit:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```
3. Build release binaries:
   ```bash
   make build-all
   ```
4. Create GitHub release:
   ```bash
   gh release create v1.0.0 ./dist/*
   ```

## Troubleshooting

### Go module not found

```bash
go mod download
go mod tidy
```

### Tests failing

```bash
# Clean cache
go clean -testcache

# Run with verbose output
go test -v ./...
```

### Beads sync issues

```bash
# Check sync status
bd sync --status

# Force sync
bd sync --force
```

## Useful Commands

```bash
# Show project statistics
bd stats

# Show blocked issues
bd blocked

# Show all open issues
bd list --status open

# Create new issue
bd create --title="Feature X" --type=feature --priority=1
```

## Resources

- [Go Documentation](https://go.dev/doc/)
- [Beads Documentation](https://github.com/steveyegge/beads)
- [Gas Town Architecture](https://github.com/steveyegge/gas-town)
- [Effective Go](https://go.dev/doc/effective_go)
