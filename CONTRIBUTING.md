# Contributing to Ghost-Town

Thank you for your interest in contributing to ghost-town!

## Development Setup

Follow the [DEVELOPMENT.md](../DEVELOPMENT.md) guide to set up your development environment.

## Workflow

1. **Check for work**: Run `bd ready` to see available tasks
2. **Claim a task**: `bd update <id> --status in_progress`
3. **Create a branch**: `git checkout -b feature/your-feature-name`
4. **Make changes**: Write code, tests, and documentation
5. **Run tests**: `make test`
6. **Lint code**: `make lint`
7. **Commit changes**: Follow [commit message guidelines](#commit-messages)
8. **Push branch**: `git push origin feature/your-feature-name`
9. **Complete task**: `bd close <id> --reason "Description"`
10. **Create pull request**: `gh pr create`

## Code Style

- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Format code with `gofmt -s -w .`
- Run `go vet ./...` before committing
- Write tests for all new functionality
- Keep functions focused and under 50 lines when possible

## Testing

- Write unit tests for all new functions
- Aim for >80% code coverage
- Use table-driven tests for multiple test cases
- Add integration tests for complex workflows

```go
func TestFunction(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "basic case",
			input: "test",
			want:  "expected",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Function(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Function() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Function() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

## Commit Messages

Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

**Examples:**
```
feat(cli): add version command
fix(config): handle missing config file gracefully
docs(readme): update installation instructions
```

## Documentation

- Update README.md for user-facing changes
- Add inline comments for complex logic
- Document exported functions and types
- Keep docs/ARCHITECTURE.md updated for structural changes

## Issue Tracking

We use **beads** for issue tracking:
- `bd ready` - Show available work
- `bd create --title="..."` - Create new issue
- `bd close <id>` - Mark complete
- `bd sync` - Sync to git

## Pull Requests

1. Keep PRs focused on a single issue
2. Include tests for new functionality
3. Update documentation as needed
4. Ensure all CI checks pass
5. Request review from at least one maintainer

## Getting Help

- Create an issue for bugs or feature requests
- Join discussions for questions
- Check existing documentation first

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
