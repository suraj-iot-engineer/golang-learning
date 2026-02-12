# 02. Modules & Workspace

Managing dependencies is a critical skill for any Go engineer. This module covers `go mod`, semantic import versioning, and the powerful `go.work` for multi-module development.

## 📚 Topics Covered

1.  **Go Modules (`go.mod`)**: Initializing and managing dependencies.
2.  **Versioning**: Understanding semver (v1.0.0, v2.0.0).
3.  **Go Workspaces (`go.work`)**: Developing across multiple local modules simultaneously without `replace` directives.
4.  **Vendor Directory**: Copying dependencies locally for offline builds.

## 🛠️ Key Commands

```bash
# Initialize a new module
go mod init github.com/username/project

# Add a dependency
go get github.com/fatih/color

# Clean up unused dependencies
go mod tidy

# Verify module integrity
go mod verify
```

## 🚀 Running the Example

This example uses the external `color` package to demonstrate dependency management.

```bash
go run main.go
```
