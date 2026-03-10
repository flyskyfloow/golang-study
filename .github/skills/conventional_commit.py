#!/usr/bin/env python3
"""
conventional_commit.py

Generates a Conventional Commits style commit message from a change summary.
Usage:
    python3 conventional_commit.py "your change summary here"

Example:
    python3 conventional_commit.py "fix bug in user login logic"

Outputs a valid Conventional Commits commit message.
"""
import sys
import re

def infer_type_and_scope(summary):
    # Simple heuristics for type
    summary_lower = summary.lower()
    if summary_lower.startswith(("fix", "bugfix", "bug fix")):
        ctype = "fix"
    elif summary_lower.startswith(("feat", "feature", "add")):
        ctype = "feat"
    elif summary_lower.startswith(("docs", "doc", "documentation")):
        ctype = "docs"
    elif summary_lower.startswith(("refactor",)):
        ctype = "refactor"
    elif summary_lower.startswith(("test", "tests", "testing")):
        ctype = "test"
    elif summary_lower.startswith(("chore", "ci", "build")):
        ctype = "chore"
    else:
        ctype = "chore"
    # Try to extract scope in parentheses, e.g. "fix(login): ..."
    match = re.match(r"(\w+)\(([^)]+)\):?\s*(.*)", summary)
    if match:
        ctype = match.group(1)
        scope = match.group(2)
        rest = match.group(3)
        return ctype, scope, rest
    # Try to extract scope from first word if in brackets
    match = re.match(r"(\w+)\s*\[([\w-]+)\]\s*:?\s*(.*)", summary)
    if match:
        ctype = match.group(1)
        scope = match.group(2)
        rest = match.group(3)
        return ctype, scope, rest
    # Otherwise, no scope
    return ctype, None, summary

def generate_commit_message(summary):
    ctype, scope, rest = infer_type_and_scope(summary)
    if scope:
        header = f"{ctype}({scope}): {rest}"
    else:
        header = f"{ctype}: {rest}"
    # Conventional Commits: header max 72 chars, wrap if needed
    if len(header) > 72:
        header = header[:69] + "..."
    return header

def main():
    if len(sys.argv) < 2:
        print("Usage: python3 conventional_commit.py \"your change summary here\"")
        sys.exit(1)
    summary = sys.argv[1]
    commit_msg = generate_commit_message(summary)
    print(commit_msg)

if __name__ == "__main__":
    main()
