# 15. Capstone Projects

Put everything together. This module contains specifications for large-scale projects to test your skills.

## 🏆 Project 1: Cloud-Native Microservices Backend

**Goal**: Build a distributed e-commerce backend.

- **Services**: Product Service, Order Service, User Service.
- **Tech Stack**: Go, gRPC, PostgreSQL, Redis, Docker, Kubernetes.
- **Requirements**:
    - Centralized Logging (ELK or similar).
    - Distributed Tracing (Jaeger).
    - CI/CD Pipeline.

## 🏆 Project 2: Real-time IoT Data Platform

**Goal**: Ingest and visualize data from 10,000+ sensors.

- **Tech Stack**: Go, MQTT, TimescaleDB, Grafana.
- **Requirements**:
    - High-throughput ingestion (use Worker Pools).
    - Real-time WebSocket dashboard.
    - Alerting system (Send email/SMS on high temp).

## 🚀 Getting Started

The `impl/` directory contains a skeleton for Project 2 (IoT Platform).

```bash
cd impl/iot-platform
go run main.go
```
