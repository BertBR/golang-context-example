### Golang Context Example

This repository serves as a practical illustration of leveraging Golang's `context` package to seamlessly integrate user information into API requests through an authentication middleware. By adopting this approach, you gain the ability to utilize user information for purposes such as tracking, context cancellation, and recovery.

#### How to Test:

1. Clone the repository.
2. Execute `go mod tidy`.
3. Run the application using `go run cmd/main.go`.

#### Request Example:

To exemplify the utilization of the implemented functionality, execute the following CURL command:

```bash
curl -X POST http://localhost:8080/users -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJ1dWlkIjoieHB0byJ9.MQtNhTd7J5uMZ86TdCZY2HXW1CZAnnWcdSnaLfjahKM"
```

#### Recommended Reading:

1. [Golang Context Package Documentation](https://pkg.go.dev/context)
2. [Defer, Panic, and Recover in Go](https://go.dev/blog/defer-panic-and-recover)