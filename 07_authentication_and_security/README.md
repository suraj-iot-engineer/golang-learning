# 07. Authentication & Security

Security is non-negotiable. This module demonstrates how to secure a Go application using JWT (JSON Web Tokens) and best practices.

## 📚 Topics Covered

1.  **JWT Authentication**: Issuing and verifying tokens.
2.  **Password Hashing**: Using `bcrypt` to store passwords securely.
3.  **Middleware Protection**: Guarding routes with Auth middleware.
4.  **Secure Headers**: Preventing XSS, CSRF, etc.
5.  **Environment Variables**: Handling secrets.

## 🛠️ Key Libraries

- `github.com/golang-jwt/jwt/v5`: JWT standard implementation.
- `golang.org/x/crypto/bcrypt`: Secure password hashing.

## 🚀 Running the Auth Service

```bash
go run main.go
```

### Test Flow
1. **POST /signup**: Create a user.
2. **POST /login**: Get a JWT token.
3. **GET /protected**: Access with `Authorization: Bearer <token>`.
