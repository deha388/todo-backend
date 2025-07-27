# Todo Backend API

A **Clean Architecture** Go backend service with comprehensive **TDD implementation** and **CDC contract compliance** for the Todo application.

## 🏗️ Architecture

This project follows Clean Architecture principles with:

- **Domain Layer**: Business entities and interfaces
- **Application Layer**: Use cases and business logic
- **Infrastructure Layer**: Database connections and external services  
- **Interface Layer**: HTTP handlers and DTOs

## 🧪 Testing Strategy

### **Test-Driven Development (TDD)**
- ✅ **Unit Tests**: Business logic testing with mocks
- ✅ **Integration Tests**: API endpoint testing with real database
- ✅ **Contract Tests**: CDC provider tests ensuring frontend compatibility

### **Test Coverage**
```bash
# Run all tests
go test -v ./test/...

# Run specific test suites
go test -v ./test/unit/...        # Unit tests
go test -v ./test/integration/... # Integration tests
go test -v ./test/contract/...    # CDC contract tests
```

## 🚀 Local Development

### **Prerequisites**
- Go 1.21+
- SQLite (included with Go)

### **Quick Start**
```bash
# Clone repository
git clone https://github.com/deha388/todo-backend.git
cd todo-backend

# Install dependencies
go mod download

# Run tests
go test -v ./test/...

# Run application
go run cmd/main.go
```

### **API Endpoints**
- `GET /health` - Health check
- `GET /api/todos` - List all todos
- `POST /api/todos` - Create new todo

### **Example Usage**
```bash
# Create a todo
curl -X POST http://localhost:8083/api/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Learn Go"}'

# Get all todos
curl http://localhost:8083/api/todos
```

## 🐳 Docker Development

### **Build Docker Image**
```bash
# Build multi-stage Docker image
docker build -t todo-backend:latest .

# Run container
docker run -p 8083:8083 todo-backend:latest
```

### **Docker Compose (Optional)**
```yaml
# docker-compose.yml
version: '3.8'
services:
  backend:
    build: .
    ports:
      - "8083:8083"
    environment:
      - GO_ENV=development
    volumes:
      - ./data:/app/data  # SQLite persistence
```

## 📋 Contract Compliance

This backend implements **Consumer Driven Contracts (CDC)** to ensure compatibility with the frontend:

### **Frontend Contract Requirements**
- **GET /api/todos**: Returns plain array `[]` (not wrapped object)
- **POST /api/todos**: Returns `{id, text, createdAt}` (no `updatedAt`)
- **Time Format**: UTC with `.000Z` suffix (`2024-01-01T10:00:00.000Z`)
- **Content-Type**: `application/json`

### **Contract Testing**
```bash
# Run CDC provider tests
go test -v ./test/contract/

# Expected output:
# ✅ TestGetAllTodos_NoTodosExist
# ✅ TestGetAllTodos_TodosExist  
# ✅ TestCreateTodo_BackendReady
```

## 🚀 Production Deployment

### **Automated Deployment**
This project uses **GitHub Actions** for automated deployment to Kubernetes:

1. **Push to main branch** triggers deployment pipeline
2. **Tests run** (unit + integration + contract)
3. **Docker image built** with multi-stage Dockerfile
4. **Deployed to K8s cluster** with health checks

### **Deployment Configuration**
- **Service Name**: `todo-backend-service`
- **Port**: `8083`
- **Health Check**: `/health` endpoint
- **Resources**: 128Mi-512Mi RAM, 100m-300m CPU

### **Environment Variables**
```bash
PORT=8083
GO_ENV=production
```

## 📁 Project Structure

```
todo-backend/
├── cmd/main.go                          # Application entry point
├── internal/
│   ├── application/usecases/            # Business logic layer
│   ├── domain/
│   │   ├── entities/                    # Business entities
│   │   └── repositories/                # Repository interfaces
│   ├── infrastructure/
│   │   ├── config/                      # Configuration management
│   │   └── database/                    # SQLite implementation
│   └── interfaces/
│       ├── dto/                         # Data transfer objects
│       ├── handlers/                    # HTTP handlers
│       └── routes/                      # Route definitions
├── test/
│   ├── unit/                           # Unit tests
│   ├── integration/                    # Integration tests
│   └── contract/                       # CDC provider tests
├── configs/config.yaml                 # Configuration file
├── Dockerfile                          # Multi-stage Docker build
├── .github/workflows/                  # GitHub Actions pipeline
└── README.md
```

## 🔧 Development Tools

### **Makefile Commands**
```bash
make test       # Run all tests
make build      # Build application
make run        # Run application
make docker     # Build Docker image
make clean      # Clean build artifacts
```

### **Database**
- **Type**: SQLite (file-based)
- **Auto-Migration**: GORM handles schema migration
- **Location**: `todo.db` (auto-created)

## 🏷️ Version Information

- **Go Version**: 1.21
- **Web Framework**: Fiber v2
- **Database**: SQLite with GORM
- **Testing**: Testify + Suite
- **Architecture**: Clean Architecture
- **Development**: TDD with CDC compliance

## 🤝 Development Workflow

1. **Red Phase**: Write failing test
2. **Green Phase**: Write minimal code to pass
3. **Refactor Phase**: Improve code quality
4. **Repeat**: Continue TDD cycle

### **Contract Compliance Workflow**
1. Frontend defines Pact contract
2. Backend implements CDC provider tests
3. Contract tests ensure API compatibility
4. Both sides can develop independently

## 📞 Support

- **Repository**: https://github.com/deha388/todo-backend
- **Issues**: Create GitHub issue
- **Architecture**: Clean Architecture with TDD
- **Testing**: Comprehensive test coverage (unit + integration + contract) 