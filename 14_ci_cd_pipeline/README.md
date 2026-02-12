# 14. CI/CD Pipeline

Automate your build, test, and deployment process using GitHub Actions.

## 📚 Topics Covered

1.  **Continuous Integration (CI)**: Running tests and linters on every push.
2.  **Continuous Deployment (CD)**: Building Docker images and deploying.
3.  **GitHub Actions**: Workflows, Jobs, and Steps.
4.  **Linting**: Using `golangci-lint` for code quality.

## 🛠️ Workflows

This module includes a sample `.github/workflows/go.yml` that:
- Sets up Go.
- Runs `go test ./...`.
- Runs `golangci-lint`.
- Builds the binary.

## 🚀 Usage

Copy the workflow file to your root `.github/workflows/` directory to enable it for the entire repository.
