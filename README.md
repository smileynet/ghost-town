# ghost-town

A kiro-cli tool based on Gas Town architecture.

## Overview

ghost-town is an AI-assisted development tool that combines the best practices from Gas Town with modern Go development tooling. It provides task management, autonomous agent coordination, and development workflow automation.

## Project Status

🚧 **Early Development** - Project setup and research phase

## Features (Planned)

- Task management with beads (git-backed graph issue tracker)
- Autonomous agent coordination with worker/validator patterns
- Auto-compact mechanism for context management
- Self-prompt continuation for long-running sessions
- Go tooling integration (linters, formatters, testing frameworks)
- Skills and hooks for extensibility

## Project Structure

- `.beads/` - Task tracking and project management
- `.claude/` - Opencode AI agent configuration
- `cmd/` - CLI commands
- `internal/` - Internal packages
- `pkg/` - Public APIs
- `docs/` - Documentation
- `tests/` - Test suites

## Getting Started

This project is currently in the setup and research phase. 

### Prerequisites

- Go 1.21+
- bd (beads) for task tracking
- opencode for AI agent coordination

### Installation

```bash
# Clone repository
git clone https://github.com/steveyegge/ghost-town.git
cd ghost-town

# Install dependencies
go mod download

# Build
go build ./cmd/ghost-town

# Install
go install ./cmd/ghost-town
```

## Development

### Task Management

This project uses **beads** for issue tracking:

```bash
# See available work
bd ready

# Create new task
bd create "Implement feature" --type task

# Track progress
bd update <id> --status in_progress

# Complete task
bd close <id>
```

See [AGENTS.md](AGENTS.md) for full workflow documentation.

### Agent Execution

Tasks use the worker/validator protomolecule pattern:

1. Worker subagent executes work
2. Validator subagent reviews against acceptance criteria
3. Loop until validation passes
4. Complete bead and commit changes

See [docs/PROTO_EXECUTION.md](docs/PROTO_EXECUTION.md) for details.

## Current Work

See [beads project view](https://github.com/steveyegge/ghost-town/issues) for active work.

## License

[MIT License](LICENSE)
