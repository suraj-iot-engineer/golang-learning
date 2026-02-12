# 13. Docker & Kubernetes

Containerization is standard for modern backend deployment. This module covers Dockerfiles, Docker Compose, and Kubernetes Manifests.

## 📚 Topics Covered

1.  **Docker**: Creating small, secure Go images (Multi-stage builds).
2.  **Docker Compose**: Orchestrating local dev environments.
3.  **Kubernetes**: Deploying to a cluster (Deployments, Services, Ingress).

## 🚀 Example
The example includes a Dockerfile for a Go API and K8s manifests.

```bash
# Build Docker Image
docker build -t my-go-app .

# Run with Compose
docker-compose up
```
