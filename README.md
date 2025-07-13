# HTTP/1.1 Server Implementation in Go
A fully-featured HTTP/1.1 server built from using standard libraries in Golang, implementing core web server functionality without external frameworks.

[![progress-banner](https://app.codecrafters.io/courses/http-server/overview)



## 🚀 Features

- **HTTP/1.1 Protocol Support**: Full compliance with HTTP/1.1 specification
- **Persistent Connections**: Keep-alive connections for performance
- **GZIP Compression**: Automatic response compression when supported by client
- **File Serving**: Static file serving with proper MIME types
- **RESTful Routing**: Clean URL routing with method-based handlers
- **Concurrent Connections**: Goroutine-based request handling

## 🏗️ Architecture

The server follows this architecture:

```
app/
├── internal/
│   ├── handlers/     # Request handlers for different endpoints
│   ├── models/       # Data structures and HTTP parsing
│   ├── router/       # URL routing and request matching
│   └── server/       # Core server implementation
└── main.go          # Application entry point
```

### Key Modules

- **Server**: TCP connection management and request lifecycle
- **Router**: Pattern matching and handler dispatch
- **Parser**: HTTP request parsing with proper header handling
- **Handlers**: Endpoint-specific logic
- **Models**: Type definitions and response builders

## 📋 Built Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/` | Root endpoint |
| `GET` | `/echo/{text}` | Echo back the provided text |
| `GET` | `/user-agent` | Return client's User-Agent header |
| `GET` | `/files/{filename}` | Serve static files |
| `POST` | `/files/{filename}` | Upload files to server |

