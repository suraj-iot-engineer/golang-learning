# 09. Backend Architecture

Writing code is easy; architecting maintainable systems is hard. This module covers industry-standard architectural patterns.

## 📚 Topics Covered

1.  **Clean Architecture (Uncle Bob)**: Separating concerns into layers.
2.  **Hexagonal Architecture (Ports & Adapters)**: Decoupling the core business logic from external dependencies.
3.  **Domain-Driven Design (DDD)**: Modeling software after real-world business domains.
4.  **Dependency Injection**: Managing dependencies explicitly.

## 🚀 Example
The example showcases a simple order processing system using **Hexagonal Architecture**.

- **Core**: Business logic (Order Service)
- **Ports**: Interfaces (Repository, specialized interfaces)
- **Adapters**: Implementations (In-memory DB, CLI handler)

```bash
go run main.go
```
