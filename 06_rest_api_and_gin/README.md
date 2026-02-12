# 06. REST API with Gin

Building robust RESTful APIs is a core backend skill. We use **Gin**, a high-performance HTTP web framework.

## 📚 Topics Covered

1.  **Gin Framework**: Setting up a router and handlers.
2.  **REST Principles**: Resources, HTTP Verbs (GET, POST, PUT, DELETE), Status Codes.
3.  **Request Binding**: Parsing JSON bodies automatically.
4.  **Middleware**: Logging, Auth, Recovery.
5.  **Validation**: Using `validator` tags.

## 🛠️ Key Commands

```bash
# Install Gin
go get -u github.com/gin-gonic/gin
```

## 🚀 Running the API

```bash
go run main.go
```
The server will start on `http://localhost:8080`.

### Test Endpoints
- **GET** `/ping`: Health check
- **GET** `/albums`: List all albums
- **POST** `/albums`: Create a new album
