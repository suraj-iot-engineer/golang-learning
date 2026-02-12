# 🚀 Master Golang: From Zero to Cloud-Native Backend Architect

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg?style=for-the-badge)]()
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)](http://makeapullrequest.com)

> **The Ultimate Guide to Professional Go Engineering.**
> Learn backend development, microservices, cloud-native systems, and high-performance IoT architecture using idiomatic Go.

---

## � Course Curriclum

This repository is structured as a complete interactive course. Each module contains **theory**, **real-world code**, **exercises**, and **projects**.

### 🔹 Phase 1: Core Systems & Concurrency
| Module | Topic | Key Concepts |
| :--- | :--- | :--- |
| **[01 Fundamentals](./01_go_fundamentals)** | **Go Syntax & Types** | Zero values, Structs, Interfaces, Error Handling |
| **[02 Dependency Management](./02_modules_and_workspace)** | **Modules & Workspaces** | `go.mod`, Private Repos, Versioning, `go.work` |
| **[03 Concurrency](./03_concurrency_and_context)** | **Goroutines & Channels** | WaitGroups, Mutexes, Select, Context API, Worker Pools |

### 🔹 Phase 2: Modern Backend Engineering
| Module | Topic | Key Concepts |
| :--- | :--- | :--- |
| **[04 Testing](./04_testing_and_debugging)** | **Quality Assurance** | Table Driven Tests, Fuzzing, Benchmarking, Delve |
| **[05 Databases](./05_databases_sql_nosql)** | **Data Persistence** | `database/sql`, GORM, PostgreSQL, Redis, MongoDB, Migrations |
| **[06 REST APIs](./06_rest_api_and_gin)** | **API Development** | Gin Framework, Middleware, Validation, OpenAPI/Swagger |
| **[07 Security](./07_authentication_and_security)** | **Auth & Security** | JWT, OAuth2, Role-Based Access Control (RBAC), Hashing |

### � Phase 3: Advanced Architecture & Microservices
| Module | Topic | Key Concepts |
| :--- | :--- | :--- |
| **[08 Advanced Patterns](./08_advanced_go_patterns)** | **Expert Go** | Reflection, Generics, Functional Options, Middleware Pattern |
| **[09 Architecture](./09_backend_architecture)** | **System Design** | Clean Architecture, Hexagonal (Ports & Adapters), DDD |
| **[10 Microservices](./10_microservices)** | **Distributed Systems** | gRPC, Protobuf, Service Discovery, Circuit Breakers |

### 🔹 Phase 4: Cloud Native & IoT
| Module | Topic | Key Concepts |
| :--- | :--- | :--- |
| **[11 Cloud Native](./11_cloud_native_go)** | **Cloud SDKs** | AWS/GCP Integration, Lambda/Cloud Functions, Storage |
| **[12 IoT Integration](./12_iot_backend_projects)** | **IoT Backends** | MQTT (Paho), Real-time Sensor Ingestion, Data Pipelines |
| **[13 Containers](./13_docker_and_kubernetes)** | **Deployment** | Dockerfiles, Multi-stage Builds, Kubernetes Manifests, Helm |

### 🔹 Phase 5: DevOps & Capstones
| Module | Topic | Key Concepts |
| :--- | :--- | :--- |
| **[14 CI/CD](./14_ci_cd_pipeline)** | **Automation** | GitHub Actions, Linters (GolangCI-Lint), Release Engineering |
| **[15 Capstones](./15_capstone_projects)** | **Production Projects** | Full Microservices Backend, Scalable IoT Platform |

---

## �️ Getting Started

### Prerequisites
- **Go 1.22+** installed ([Download](https://go.dev/dl/))
- **Docker** & **Docker Compose**
- **VS Code** (Recommended) with Go Extension

### Installation
1. **Clone the repository**
   ```bash
   git clone https://github.com/your-username/golang-learning.git
   cd golang-learning
   ```

2. **Initialize Workspace** (Optional but recommended for multi-module dev)
   ```bash
   go work init
   go work use ./01_go_fundamentals ./02_modules_and_workspace # ... add usage as you go
   ```

3. **Run Examples**
   Navigate to any module and run:
   ```bash
   cd 01_go_fundamentals
   go run main.go
   ```

---

## 🏗️ Project Roadmap

- [ ] **Fundamentals**: Complete core syntax and error handling guides.
- [ ] **API Layer**: Build a production-ready REST API with Gin.
- [ ] **Microservices**: Implement gRPC communication between services.
- [ ] **IoT**: Create a real-time MQTT dashboard backend.
- [ ] **Deployment**: Deploy the full stack to Kubernetes.

---

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guide](docs/CONTRIBUTING.md) to get started.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

---

<p align="center">
  Connect with me: <a href="https://linkedin.com">LinkedIn</a> • <a href="https://twitter.com">Twitter</a>
</p>
