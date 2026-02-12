# 10. Microservices with gRPC

Microservices require efficient communication. this module introduces gRPC and Protocol Buffers.

## 📚 Topics Covered

1.  **gRPC vs REST**: Protocol Buffers vs JSON.
2.  **Protocol Buffers (.proto)**: Defining service contracts.
3.  **Code Generation**: Using `protoc` to generate Go code.
4.  **gRPC Server & Client**: Implementing the service.

## 🛠️ Prerequisites

You need the `protoc` compiler installed and the Go plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 🚀 Running the Example

The example includes a simple "Greeter" service.

```bash
# Start Server
go run server/main.go

# Run Client (in another terminal)
go run client/main.go
```
