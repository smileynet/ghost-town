# Protomolecule: Project Task Execution

Reusable workflow pattern for autonomous task execution using subagents with validation loops.

## Purpose

Execute tasks through subagents that retrieve context, complete work, and validate results before closing beads. Ensures quality and completion before marking work as done.

## Workflow Pattern

```
Primary Agent Loop:
  1. Claim bead (task) with `bd update <id> --status in_progress`
  2. Launch Worker Subagent
  3. Worker Subagent:
     - Retrieves bead context with `bd show <id>`
     - Executes work based on bead requirements
     - Notifies primary agent when complete
  4. Primary Agent launches Validator Subagent
  5. Validator Subagent:
     - Reviews completed work against bead acceptance criteria
     - Checks completion status
     - Returns validation result
  6. IF validation fails:
     - Return to step 2 (re-launch Worker Subagent)
  7. ELSE validation passes:
     - Complete bead with `bd close <id>`
     - Commit changes to git
```

## Bead Structure

Tasks should include:

- **description**: What needs to be done
- **acceptance**: Specific criteria for completion
- **notes**: Context or constraints

```bash
bd create "Task title" \
  --type task \
  --description "What the task does" \
  --acceptance "Specific completion criteria" \
  --notes "Additional context or constraints"
```

## Subagent Protocol

### Worker Subagent

**Prompt Template:**
```
You are a worker subagent executing task <bead-id>.

Task Context:
```
Use `bd show <bead-id>` to get full context
```

Requirements:
- Retrieve bead details using `bd show <bead-id>`
- Complete work according to bead description
- Follow any constraints in bead notes
- Return completion status when finished

Output: Report what you accomplished and any issues encountered.
```

### Validator Subagent

**Prompt Template:**
```
You are a validator subagent checking work for task <bead-id>.

Task Context:
```
Use `bd show <bead-id>` to get requirements and acceptance criteria
```

Validation Checklist:
- Review bead acceptance criteria
- Verify work meets all requirements
- Check for completion
- Identify any gaps or issues

Output: 
- PASS if work is complete and meets acceptance criteria
- FAIL with specific gaps if incomplete
```

## Implementation Commands

### Primary Agent Workflow

```bash
# Step 1: Claim work
bd update <bead-id> --status in_progress

# Step 2: Launch worker subagent (repeat until complete)
task subagent_type=general \
  description="Execute task <bead-id>" \
  prompt="You are a worker subagent. Use \`bd show <bead-id>\` to get context, complete the work, and report completion."

# Step 3: Launch validator subagent
task subagent_type=general \
  description="Validate task <bead-id>" \
  prompt="You are a validator. Use \`bd show <bead-id>\` to check acceptance criteria. Return PASS or FAIL with details."

# Step 4: Check validation result and loop or continue
# If PASS → complete bead
# If FAIL → go to Step 2

# Step 5: Complete bead when validated
bd close <bead-id> --reason "Validated and complete"

# Step 6: Commit changes
git add .
git commit -m "Complete <task-name> (<bead-id>)"
```

## Validation Criteria

Work is considered complete when:

1. All acceptance criteria in bead are met
2. Code changes are committed (if applicable)
3. Tests pass (if applicable)
4. Documentation is updated (if required)
5. No gaps or issues identified by validator

## Error Handling

### Worker Fails to Complete

- Log error to bead notes with `bd update <bead-id> --notes "Error: ..."`
- Return failure status to primary agent
- Primary agent decides whether to retry or escalate

### Validation Fails

- Validator returns specific gaps/issues
- Primary agent relaunches worker with feedback
- Loop continues until validation passes

### Context Loss

- If bead context is unclear, worker should query `bd show <bead-id>`
- If acceptance criteria are ambiguous, worker should ask for clarification
- Use `bd update <bead-id>` to add discovered requirements if needed

## Example Usage

```bash
# Primary agent claims a research task
bd update bd-ihr --status in_progress

# Launch worker
task subagent_type=general \
  description="Research kiro-cli components" \
  prompt="Use \`bd show bd-ihr\` to get context. Research kiro-cli components and document findings."

# Worker completes and reports back
# Output: "Researched hooks, steering agents, skills. Created docs/kiro-cli-components.md"

# Launch validator
task subagent_type=general \
  description="Validate kiro-cli research" \
  prompt="Use \`bd show bd-ihr\` to check acceptance criteria. Verify research is complete."

# Validator returns PASS

# Complete bead
bd close bd-ihr --reason "Research validated and complete"

# Commit changes
git add docs/kiro-cli-components.md
git commit -m "Complete kiro-cli research (bd-ihr)"
```

## Integration with Beads Workflow

This protomolecule integrates with existing beads workflow:

- Use `bd ready` to find available tasks
- Use `bd blocked` to check dependencies
- Use `bd close` to mark completion
- Use `bd sync` at session end

## Advantages

1. **Quality Assurance**: Validation ensures work meets criteria before closing
2. **Autonomy**: Subagents work independently with clear protocols
3. **Repeatability**: Failed validation triggers retry without manual intervention
4. **Traceability**: Each step is logged to bead status and notes
5. **Accountability**: Validation provides clear completion signal

## Related Patterns

- **Sequential Pipeline**: For multi-step workflows within a single task
- **Parallel Fanout**: For tasks that can be validated independently
- **Dynamic Bonding**: For tasks that spawn sub-tasks during execution

## See Also

- [MOLECULES.md](MOLECULES.md) - Core molecules concepts
- [CLI_REFERENCE.md](CLI_REFERENCE.md) - Full command reference
- [CLAUDE.md](../CLAUDE.md) - Agent instructions
