# Copilot: Commit Message Generation Skill

## Purpose
Enable Copilot/ChatGPT to generate commit messages that strictly follow the Conventional Commits style as documented in `copilot-commit-skill.md`.

## Input
- **change_summary** (string): A summary of the changes, or a list of staged files/changes. This is the only required input.

## Output
- **commit_message** (string): The generated commit message, following these rules:
  - First line: `type(scope?): short summary` (max 50 chars)
  - Allowed types: build, chore, ci, docs, feat, fix, perf, refactor, revert, style, test
  - Scope is optional, lowercase, hyphenated if present
  - Optional body: lines <= 72 chars, further details if needed
  - No explanations or extra output—only the commit message text

## Language
- Prompt the model in Chinese or English as needed, but output must always be a valid Conventional Commit message in English.

## Example
**Input:**
change_summary: "Add JWT refresh endpoint to authentication module"

**Output:**
commit_message: "feat(auth): add JWT refresh endpoint"

---
**Input:**
change_summary: "Update CI to use Go 1.22, fix flaky test in user package"

**Output:**
commit_message: "ci: update to Go 1.22\ntest(user): fix flaky test"

## Usage
- Call this skill with a change summary or staged file list. The output is a commit message ready for use with `git commit -m`.
- For multi-part changes, generate a single message with multiple lines, each following the Conventional Commits style.

## Reference
See `.github/skills/copilot-commit-skill.md` for full commit style rules and prompt templates.
