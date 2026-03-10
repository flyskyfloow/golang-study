# Copilot Instructions for golang-study

## Build, Test, and Lint Commands

- **Run all tests:**
  ```bash
  go test ./...
  ```
- **Run a single test file:**
  ```bash
  go test ./internal/app/app_test.go
  ```
- **Run the application:**
  ```bash
  go run ./cmd/app
  ```
- **Dependencies:**
  - Go modules are used (see `go.mod`).
  - No custom Makefile or scripts are required for basic build/test.

## High-Level Architecture

- **Modern Go Project Layout:**
  - `cmd/app/`: Application entry point (main.go)
  - `internal/app/`: Internal business logic (not for external use)
  - `pkg/version/`: Publicly reusable code (e.g., version info)
  - `api/`, `build/`, `configs/`, `deployments/`, `docs/`, `scripts/`, `test/`: Reserved for future expansion (currently mostly placeholders)
- **Typical flow:**
  - `main.go` in `cmd/app/` calls into `internal/app/` for core logic.
  - Versioning is managed in `pkg/version/`.

## Key Conventions

- **Directory usage follows [Standard Go Project Layout](https://github.com/golang-standards/project-layout)**
- **Tests** are colocated with the code they test, using Go's standard `*_test.go` pattern.
- **No custom linting or CI configuration** is present; use standard Go tools.
- **All business logic intended for internal use should go in `internal/`**; only reusable libraries go in `pkg/`.

---

This file was generated to help Copilot and other AI tools understand the structure and conventions of this repository. If you want to adjust or expand these instructions, let me know!
