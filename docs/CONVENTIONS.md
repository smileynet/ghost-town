# Documentation Conventions

This document establishes conventions for documentation in the ghost-town project.

## Document Structure

### Root-Level Documents

- **README.md**: Project overview, quick start, and usage
- **CONTRIBUTING.md**: Contribution guidelines and workflow
- **DEVELOPMENT.md**: Development environment setup
- **LICENSE**: License information

### docs/ Directory

- **ARCHITECTURE.md**: System architecture and design decisions
- **API.md**: Public API reference
- **CONVENTIONS.md**: This file - documentation and code conventions
- **PROTO_EXECUTION.md**: Protomolecule pattern documentation
- **research/**: Research findings and analysis

## Writing Style

### General Guidelines

- Use clear, concise language
- Write in second person ("you", "your") for tutorials
- Write in present tense for descriptions
- Avoid jargon when possible
- Provide examples for complex concepts

### Code Examples

- Use fenced code blocks with syntax highlighting
- Include comments in examples
- Show both correct and incorrect usage when relevant
- Test all code examples

**Example:**

\`\`\`go
// Good
config, err := loadConfig("config.yaml")
if err != nil {
    log.Fatal(err)
}

// Bad
config := loadConfig("config.yaml")
\`\`\`

### Headings

- Use ATX-style headings (`#`, `##`, `###`)
- Skip heading levels (don't jump from `#` to `###`)
- Keep headings short and descriptive
- Use sentence case (not Title Case)

## Code Documentation

### Package Comments

Every package should have a doc comment describing its purpose:

```go
// Package logger provides logging interfaces and implementations
// for the ghost-town application.
package logger
```

### Exported Functions and Types

All exported functions and types must have doc comments:

```go
// Info logs an informational message to the configured output.
func (l *DefaultLogger) Info(msg string) {
    // implementation
}
```

### Function Documentation Structure

```go
// FunctionName does something specific and useful.
//
// It takes the following parameters:
//   - param1: Description of param1
//   - param2: Description of param2
//
// It returns:
//   - result1: Description of result1
//   - error: Description of error conditions
//
// Example:
//
//	result, err := FunctionName("input")
//	if err != nil {
//	    // handle error
//	}
func FunctionName(param1 string, param2 int) (result1 string, err error) {
    // implementation
}
```

## README Structure

Every README.md should follow this structure:

1. **Title and Description**: What is this?
2. **Features**: What can it do?
3. **Installation**: How to install it
4. **Quick Start**: Get running in 5 minutes
5. **Usage**: How to use it
6. **Configuration**: How to configure it
7. **Examples**: Code examples
8. **Contributing**: How to contribute
9. **License**: License information

## Diagrams and Visuals

When including diagrams:
- Use ASCII art or Mermaid.js for simple diagrams
- Keep diagrams simple and readable
- Explain complex diagrams in accompanying text
- Include legends when necessary

**Mermaid Example:**

\`\`\`mermaid
graph LR
    A[Input] --> B[Process]
    B --> C[Output]
\`\`\`

## Version-Specific Documentation

- Keep documentation up-to-date with the current version
- Add version notes for breaking changes
- Document deprecations clearly
- Provide migration guides for major changes

## API Documentation

- Document all public APIs in docs/API.md
- Include function signatures
- Provide usage examples
- Document error conditions
- Note performance characteristics

## Research Documentation

Research documents should include:
- **Objective**: What was being researched?
- **Methodology**: How was it researched?
- **Findings**: What was discovered?
- **Recommendations**: What should we do?
- **References**: Sources and links

## Review Process

Before committing documentation:
1. Check for typos and grammar errors
2. Test all code examples
3. Verify all links work
4. Ensure structure is consistent
5. Update table of contents if needed

## Keeping Documentation Current

- Update docs when code changes
- Review documentation during code review
- Remove outdated information
- Document decisions and rationale
- Keep changelog for user-facing changes

## Tools and Resources

- **Markdownlint**: `npm install -g markdownlint-cli`
- **Typos**: `cargo install typos-cli`
- **Link Checker**: `npm install -g markdown-link-check`

## Common Mistakes to Avoid

1. **Outdated Examples**: Always test code examples
2. **Assumptions**: Don't assume prior knowledge
3. **Vague Instructions**: Be specific and clear
4. **Missing Context**: Provide background information
5. **Unexplained Jargon**: Define technical terms

## Documentation Templates

### Command Documentation

\`\`\`
### command-name

Brief description of what the command does.

**Usage:**

\`\`\`bash
ghost-town command-name [options] [arguments]
\`\`\`

**Options:**

- \`-f, --file\`: Specify input file
- \`-v, --verbose\`: Enable verbose output

**Examples:**

\`\`\`bash
# Basic usage
ghost-town command-name

# With options
ghost-town command-name -f input.txt -v
\`\`\`

**See Also:**
- Related command 1
- Related command 2
\`\`\`

### Feature Documentation

\`\`\`
## Feature Name

Brief description of the feature.

**Purpose:** Why this feature exists

**When to Use:** Best use cases

**How It Works:** Technical explanation

**Example:**

\`\`\`go
// Code example
\`\`\`

**Limitations:** Any known constraints
\`\`\`
