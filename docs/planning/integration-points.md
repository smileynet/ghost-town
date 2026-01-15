# Integration Points Planning

## Overview

This document outlines integration points between ghost-town and external tools and systems, including git, Go tooling, and AI agent frameworks.

## Integration Categories

### 1. Git Integration

#### Primary Integration Points

**Git Hooks as Propulsion Mechanism**
- **Purpose:** Git hooks act as "pistons" triggering ghost-town workflows
- **Pattern:** Gas Town Universal Propulsion Principle (GUPP) - "If there is work on your Hook, YOU MUST RUN IT"
- **Implementation:**
  - Pre-commit hooks: Run linters, tests before commit
  - Post-commit hooks: Update beads database with commit info
  - Post-merge hooks: Trigger deployment workflows
  - Pre-push hooks: Final validation before pushing

**Location:** `.git/hooks/`

**Example Pre-commit Hook:**
```bash
#!/bin/bash
# .git/hooks/pre-commit

# Run quality gates
ghost-town check --pre-commit

# Update beads with staged changes
ghost-town beads sync --staged
```

**Git Worktrees**
- **Purpose:** Isolated workspaces for parallel agent work
- **Use Cases:**
  - Polecat agents work in isolated worktrees
  - Crew members maintain persistent clones
  - Cross-rig work via worktrees
- **Implementation:**
  ```bash
  # Create worktree for polecat
  git worktree add .ghost-town/polecats/task-123 -b polecat/task-123

  # Cleanup after completion
  git worktree remove .ghost-town/polecats/task-123
  git branch -D polecat/task-123
  ```

**Git Attributes**
- **Purpose:** Project-specific configuration
- **Location:** `.gitattributes`
- **Configuration:**
  ```
  *.go text eol=lf
  .beads/ export-ignore
  .ghost-town/ export-ignore
  ```

**Git Configuration**
- **Purpose:** Ghost-town specific settings
- **Location:** `.git/config`
- **Settings:**
  ```ini
  [ghost-town]
      enabled = true
      auto-sync = true
      branch-prefix = ghost-town/
  ```

#### Git Operations Integration

**Branch Management**
- Auto-create feature branches from beads
- Track branch to issue mapping
- Clean up stale branches

**Commit Attribution**
- Format: `ghost-town/crew/<agent>` or `ghost-town/polecats/<name>`
- Link commits to beads via commit messages
- Support co-authorship for multi-agent work

**Merge Coordination**
- Detect merge conflicts
- Trigger conflict resolution workflows
- Record merge history in beads

**Ref Management**
- Auto-create MRs from completed polecat work
- Link MRs to beads
- Track MR status in beads

#### Integration Requirements

**Git Commands to Wrap:**
```bash
ghost-town commit          # Wrapper for git commit with beads update
ghost-town push            # Wrapper for git push with sync
ghost-town merge            # Wrapper for git merge with conflict handling
ghost-town branch           # Wrapper for git branch with issue linking
ghost-town worktree         # Wrapper for git worktree with isolation
```

**Git Hooks to Install:**
- `pre-commit`: Run quality gates, update beads
- `post-commit`: Record commit in beads
- `post-merge`: Trigger workflows
- `pre-push`: Final validation, sync beads

**Git Status Integration:**
- Show bead status alongside git status
- Color-coded indicators (○ ◐ ●)
- Merge bead metadata into git status output

---

### 2. Go Tooling Integration

#### Core Tooling (In Makefile)

**Format Check**
```makefile
fmt:
	@echo "Formatting code..."
	gofmt -s -w .
	goimports -w .
```

**Static Analysis**
```makefile
lint:
	@echo "Running linters..."
	golangci-lint run --timeout 5m
```

**Vet**
```makefile
vet:
	@echo "Running go vet..."
	go vet ./...
```

**Testing**
```makefile
test:
	@echo "Running tests..."
	go test -v -race -cover ./...

test-coverage:
	@echo "Running coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
```

**Build**
```makefile
build:
	@echo "Building ghost-town..."
	go build -o bin/ghost-town ./cmd/ghost-town

build-all:
	@echo "Building for all platforms..."
	GOOS=linux GOARCH=amd64 go build -o dist/ghost-town-linux-amd64 ./cmd/ghost-town
	GOOS=darwin GOARCH=arm64 go build -o dist/ghost-town-darwin-arm64 ./cmd/ghost-town
	GOOS=windows GOARCH=amd64 go build -o dist/ghost-town-windows-amd64.exe ./cmd/ghost-town
```

#### Advanced Tooling

**Hot Reload (air)**
- **Purpose:** Live development with auto-rebuild
- **Configuration:** `.air.toml`
- **Integration:**
  ```toml
  root = "."
  tmp_dir = "tmp"

  [build]
    cmd = "go build -o ./tmp/main ./cmd/ghost-town"
    bin = "tmp/main"
    delay = 0
    exclude_dir = ["tmp", "vendor"]

  [log]
    time = true
  ```
- **Makefile target:**
  ```makefile
  dev:
      air
  ```

**Mock Generation (go-mock)**
- **Purpose:** Generate mocks for testing
- **Integration:**
  ```makefile
  mock-gen:
      @echo "Generating mocks..."
      mockgen -source=internal/agent/interface.go -destination=internal/agent/mock/mock.go
  ```

**Dependency Management**
```makefile
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

deps-update:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy
```

#### Linter Configuration

**.golangci.yml**
```yaml
linters:
  enable:
    - gofmt
    - govet
    - staticcheck
    - gosimple
    - unused
    - errcheck
    - gosec

linters-settings:
  govet:
    enable-all: true
  errcheck:
    check-type-assertions: true

run:
  timeout: 5m
  tests: true
```

#### CI/CD Integration

**GitHub Actions**
```yaml
name: CI

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - name: Install linter
        run: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
      - name: Download dependencies
        run: go mod download
      - name: Run linter
        run: golangci-lint run
      - name: Run tests
        run: go test -v -race -cover ./...
      - name: Upload coverage
        uses: codecov/codecov-action@v3
```

#### Integration Requirements

**Ghost-town Commands:**
- `ghost-town check` - Run quality gates
- `ghost-town test` - Run tests with beads tracking
- `ghost-town lint` - Run linters with reporting
- `ghost-town build` - Build with version info
- `ghost-town dev` - Start dev server

**Environment Variables:**
- `GHOST_TOWN_DEBUG` - Enable debug logging
- `GHOST_TOWN_CONFIG` - Path to config file
- `GHOST_TOWN_BEADS_PATH` - Path to beads database
- `GHOST_TOWN_GIT_BIN` - Path to git binary (override)

**Configuration File:**
- Location: `~/.config/ghost-town/config.yaml`
- Settings:
  ```yaml
  git:
    auto_hooks: true
    worktree_prefix: ".ghost-town/"
  go:
    version: "1.21"
    linters: ["golangci-lint", "go vet"]
    test_coverage_threshold: 80
  build:
    ldflags: "-X main.Version={{.Version}}"
  ```

---

### 3. AI Agent Framework Integration

#### Primary Framework: Claude Code

**Integration Points:**

**1. Command Interface**
- **Purpose:** Execute ghost-town commands from Claude Code
- **Implementation:**
  - Claude Code can run shell commands
  - Ghost-town provides CLI interface
  - Transparent integration via standard shell

**2. Session Management**
- **Purpose:** Handoff between Claude Code sessions
- **Implementation:**
  ```bash
  # Handoff at end of session
  ghost-town handoff --session-id <id> --notes "progress summary"

  # Restore at start of session
  ghost-town handoff --restore --session-id <id>
  ```

**3. Context Management**
- **Purpose:** Ghost-town manages token usage and context
- **Integration:**
  - Auto-compact at 60% threshold
  - Self-prompt after 90s idle
  - Preserved in `~/.config/ghost-town/context/`

**4. Work Assignment**
- **Purpose:** Route work from beads to agents
- **Implementation:**
  ```bash
  # Get ready work
  ghost-town ready

  # Claim work
  ghost-town claim bd-123 --agent "claude-code"

  # Mark progress
  ghost-town progress bd-123 --status in_progress
  ```

#### Secondary Frameworks

**GitHub Copilot**
- **Integration:** Ghost-town commands in VS Code terminal
- **Support:** Same CLI interface as Claude Code

**OpenAI ChatGPT**
- **Integration:** Shell command execution
- **Support:** Same CLI interface

**Cursor IDE**
- **Integration:** Terminal commands
- **Support:** Same CLI interface

**Kiro Agents**
- **Purpose:** Multi-agent coordination
- **Implementation:**
  - Ghost-town as kiro CLI tool
  - Beads integration for tracking
  - Mail protocol for communication

#### Protomolecule Pattern

**Worker/Validator Loop:**
```
1. Worker agent spawns from ghost-town
2. Worker executes task from beads
3. Worker reports completion to beads
4. Validator spawns (or is primary agent)
5. Validator checks work quality
6. Validator updates beads with verdict
7. Loop repeats for next task
```

**Integration with Beads:**
```bash
# Spawn worker
ghost-town worker spawn bd-123 --agent "claude-code"

# Worker completes
ghost-town worker complete bd-123 --status success --output "results"

# Validator checks
ghost-town validator run bd-123 --agent "claude-code"

# Validator reports
ghost-town validator report bd-123 --status passed --notes "quality ok"
```

#### Message Protocol

**Mail Types:**
- `POLECAT_DONE` - Worker completed
- `MERGE_READY` - Work ready for merge
- `MERGED` - Successfully merged
- `MERGE_FAILED` - Merge failed
- `HELP` - Request assistance
- `HANDOFF` - Session continuity
- `WITNESS_PING` - Health check

**Mail Commands:**
```bash
# Send message
ghost-town mail send <addr> -s "Subject" -m "Body"

# Read inbox
ghost-town mail inbox

# Acknowledge message
ghost-town mail ack <msg-id>
```

#### Integration Requirements

**Ghost-town Commands:**
- `ghost-town ready` - List ready work
- `ghost-town claim` - Claim work
- `ghost-town handoff` - Session handoff
- `ghost-town mail` - Mail protocol
- `ghost-town worker` - Worker management
- `ghost-town validator` - Validation

**Agent Metadata:**
- Store agent identity in beads
- Track agent performance metrics
- Support multiple simultaneous agents
- Cross-session agent identity

**Session Context:**
- Maintain context across sessions
- Support handoff between sessions
- Restore previous state on reconnect
- Track session history

---

### 4. Beads Integration

#### Primary Integration

**Beads CLI Commands**
- **Location:** System PATH (`bd` command)
- **Usage:** Wrap or call directly
- **Integration:**
  ```bash
  # Create issue from ghost-town
  bd create --title "Task" --type=task --from=ghost-town

  # Sync beads with git
  bd sync

  # List ready work
  bd ready
  ```

**Beads Hooks**
- **Purpose:** Auto-inject ghost-town context
- **Implementation:**
  ```bash
  # .beads/hooks/post-create
  # Called after issue creation
  ghost-town bead created $BEAD_ID

  # .beads/hooks/pre-close
  # Called before issue close
  ghost-town bead closing $BEAD_ID
  ```

**Beads Database Integration**
- **Location:** `.beads/database.db` (SQLite)
- **Access:** Direct SQLite access or via `bd` CLI
- **Integration:**
  ```go
  // Direct access
  import "github.com/smiley/beads/internal/storage"

  store := storage.NewSQLiteStore(".beads/database.db")
  issue, err := store.GetIssue("bd-123")
  ```

#### Integration Requirements

**Ghost-town Commands:**
- `ghost-town bead create` - Create bead
- `ghost-town bead show` - Show bead details
- `ghost-town bead update` - Update bead
- `ghost-town bead close` - Close bead
- `ghost-town bead sync` - Sync with git

**Configuration:**
```yaml
beads:
  path: ".beads"
  database: "database.db"
  auto_sync: true
  hooks:
    post_create: "~/.beads/hooks/post-create"
    pre_close: "~/.beads/hooks/pre-close"
```

---

### 5. Other Tool Integrations

#### Makefile Integration

**Ghost-town as Target:**
```makefile
.PHONY: ghost-town ghost-town-check ghost-town-test ghost-town-lint

ghost-town:
	@echo "Running ghost-town ready..."
	ghost-town ready

ghost-town-check:
	@echo "Running ghost-town checks..."
	ghost-town check --all

ghost-town-test:
	@echo "Running tests via ghost-town..."
	ghost-town test

ghost-town-lint:
	@echo "Running lints via ghost-town..."
	ghost-town lint
```

#### Editor Integration

**VS Code**
- **Extension:** Ghost-town commands in terminal
- **Configuration:** `.vscode/settings.json`
  ```json
  {
    "ghost-town.enabled": true,
    "ghost-town.autoSync": true,
    "ghost-town.statusIndicator": "status-bar"
  }
  ```
- **Commands:**
  - `Ghost Town: Ready Work`
  - `Ghost Town: Claim Task`
  - `Ghost Town: Sync Beads`

**Neovim**
- **Plugin:** Telescope integration
- **Commands:**
  ```lua
  :GhostTownReady
  :GhostTownClaim
  :GhostTownSync
  ```

#### Monitoring and Metrics

**Prometheus Integration**
- **Metrics:**
  - Agent spawn rate
  - Task completion rate
  - Merge success rate
  - Context compactions
- **Endpoint:** `/metrics`

**Logging**
- **Location:** `~/.ghost-town/logs/`
- **Format:** JSON structured logs
- **Levels:** DEBUG, INFO, WARN, ERROR
- **Rotation:** Daily rotation, 7-day retention

---

## Integration Priority

### Phase 1 (MVP - P0)
1. **Git Hooks** - Pre-commit, post-commit
2. **Beads CLI** - Ready, claim, close commands
3. **Go Tooling** - Makefile targets (fmt, test, lint, build)
4. **Claude Code** - CLI command execution
5. **Session Handoff** - Basic handoff/restore

### Phase 2 (Enhanced - P1)
1. **Git Worktrees** - Isolated agent workspaces
2. **Mail Protocol** - Basic message passing
3. **Protomolecule** - Worker/validator pattern
4. **Go Linter** - golangci-lint integration
5. **Hot Reload** - air for development

### Phase 3 (Advanced - P2)
1. **Full Mail Protocol** - All message types
2. **CI/CD** - GitHub Actions integration
3. **Editor Plugins** - VS Code, Neovim
4. **Metrics** - Prometheus, monitoring
5. **Multi-Agent** - Complex coordination

---

## Open Questions

1. **Git Hooks Location**
   - Should ghost-town install hooks globally or per-project?
   - How to handle existing hooks?

2. **Beads Access Pattern**
   - Use `bd` CLI or direct SQLite access?
   - What about concurrent access?

3. **Agent Identity**
   - How to track agents across sessions?
   - Support anonymous agents?

4. **Conflict Resolution**
   - How to handle merge conflicts automatically?
   - When to involve human?

5. **Context Management**
   - How to measure token usage?
   - What summarization approach?

---

## References

- Gas Town Architecture: `docs/research/gastown-architecture.md`
- Kiro-CLI Components: `docs/research/kiro-cli-components.md`
- Go Tooling: `docs/research/go-tooling.md`
- Research Summary: `docs/research/summary.md`

---

**Document Created:** 2026-01-15
**Task:** bd-65z.1
**Status:** In Progress
