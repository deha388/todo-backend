# 🔗 Backend Deployment Information

**For Frontend Team Integration**

## 📋 Service Details

### **Kubernetes Service**
- **Service Name**: `todo-backend-service`
- **Port**: `8083`
- **Type**: `ClusterIP` (internal only)
- **Namespace**: `default`

### **API Endpoints**
```
Base URL: http://todo-backend-service:8083

Health Check:
GET /health

Todo API:
GET  /api/todos          # List all todos
POST /api/todos          # Create new todo
```

## 🤝 Contract Compliance

### **GET /api/todos Response**
```json
// Empty state
[]

// With todos
[
  {
    "id": "uuid-string",
    "text": "todo text",
    "createdAt": "2024-01-01T10:00:00.000Z"
  }
]
```

### **POST /api/todos Request/Response**
```json
// Request
{
  "text": "New todo item"
}

// Response (Status: 201)
{
  "id": "uuid-string", 
  "text": "New todo item",
  "createdAt": "2024-01-01T10:00:00.000Z"
}
```

### **Important Notes**
- ✅ **Time Format**: UTC with `.000Z` suffix
- ✅ **No `updatedAt` field** in responses
- ✅ **Plain array** for GET (not wrapped object)
- ✅ **Content-Type**: `application/json`

## 🧪 Contract Testing Status

```bash
✅ CDC Provider Tests: ALL PASSING
✅ Unit Tests: 9/9 PASSING  
✅ Integration Tests: 6/6 PASSING
✅ Docker Build: SUCCESSFUL
✅ Contract Compliance: VERIFIED
```

## 🚀 Deployment Pipeline

### **GitHub Repository**
- **URL**: https://github.com/deha388/todo-backend
- **Branch**: `main`
- **Auto-Deploy**: ✅ On push to main

### **Deployment Trigger**
```bash
# Automatic deployment on:
git push origin main

# Manual trigger available via GitHub Actions UI
```

## 🔧 Local Development (Frontend Team)

### **If you need to run backend locally:**
```bash
# Clone repository
git clone https://github.com/deha388/todo-backend.git
cd todo-backend

# Run with Go
go run cmd/main.go

# Or with Docker
docker build -t todo-backend .
docker run -p 8083:8083 todo-backend

# Test endpoints
curl http://localhost:8083/health
curl http://localhost:8083/api/todos
```

## 📞 Support

- **Issues**: Create GitHub issue at https://github.com/deha388/todo-backend/issues
- **Contract Changes**: Backend will maintain backward compatibility
- **Deployment**: Automated via GitHub Actions

---

**Backend Team**: Ready for frontend integration! ✅ 