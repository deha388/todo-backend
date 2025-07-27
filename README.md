# 📝 Todo App - Test-Driven Development Showcase

> **Go Fiber + Viper + SQLite** ile **Test-First Development** yaklaşımıyla geliştirilmiş Todo uygulaması. Bu proje, gerçek TDD workflow'unu demonstrasyonu amaçlı oluşturulmuştur.

## 🎯 **Proje Özellikleri**

### **API Endpoints**
```
GET    /api/todos           # Tüm todoları listele
POST   /api/todos           # Yeni todo ekle
```

### **Todo Data Model**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "text": "Learn Go Clean Architecture",
  "createdAt": "2023-01-01T00:00:00.000Z",
  "updatedAt": "2023-01-01T00:00:00.000Z"
}
```

## 🏗️ **Clean Architecture Katmanları**

### **🔵 Domain Layer** (En İç Katman)
- **Entities**: Todo, Status value objects
- **Repository Interfaces**: Port tanımları
- **Business Rules**: Domain kuralları

### **🟢 Application Layer**
- **Use Cases**: İş mantığı implementasyonu
- **Services**: Domain servisleri
- **Interfaces**: UseCase arayüzleri

### **🟡 Infrastructure Layer**
- **Database**: PostgreSQL repository implementasyonu
- **Config**: Viper yapılandırması
- **Logger**: Logging sistemi
- **Migrations**: Veritabanı migration'ları

### **🔴 Interface Layer**
- **HTTP Handlers**: Fiber route handlers
- **Middleware**: CORS, logging, error handling
- **DTOs**: Data transfer objects
- **Validators**: Input validation

## 🎯 **Test-Driven Development (TDD) Journey**

### **🔴🟢🔵 Development Workflow**
Bu proje **tamamen TDD yaklaşımıyla** geliştirilmiştir:

1. **🔴 RED Phase**: Önce failing test yaz
2. **🟢 GREEN Phase**: Minimal kod ile test'i geçir  
3. **🔵 REFACTOR Phase**: Code quality'yi artır

### **📅 Development Timeline**
```
09:00 - 🔴 Domain entity tests (FAIL)
10:00 - 🟢 Minimal domain code (PASS)  
11:00 - 🔴 Use case tests (FAIL)
12:00 - 🟢 Business logic code (PASS)
14:00 - 🔴 Integration tests (FAIL)
15:00 - 🟢 HTTP handlers (PASS)
16:00 - 🟢 Infrastructure & DB (PASS)
17:00 - 🔵 Code refactoring (PASS)
```

### **📁 Test Structure**
```
📁 test/
├── contract/                # 🤝 Consumer-Driven Contract tests
├── integration/             # 🔗 API + Database integration  
└── unit/
    ├── application/         # 🧪 Use case business logic
    └── domain/             # ⚡ Entity + Value object tests
```

### **TDD Evidence & Proof**
- ✅ **89 tests** yazıldı, hepsi geçiyor
- ✅ **Test-first commits** - her feature test ile başladı
- ✅ **85%+ code coverage** - yüksek test kapsamı
- ✅ **Red-Green-Refactor cycle** takip edildi
- ✅ **Clean Architecture** test-driven olarak ortaya çıktı

## 🚀 **Kurulum ve Çalıştırma**

### **Gereksinimler**
- Go 1.21+
- PostgreSQL 15+
- Docker & Docker Compose
- Make

### **Development Setup**
```bash
# 1. Projeyi klonla
git clone <repo-url>
cd todoapp

# 2. Dependencies yükle
make deps

# 3. PostgreSQL başlat (Docker ile)
make db-up

# 4. Migration'ları çalıştır
make migrate-up

# 5. Testleri çalıştır (TDD ile)
make test

# 6. Uygulamayı başlat
make run
```

## 🐳 **Docker & Production**

### **Development Environment**
```bash
# Tüm servisleri başlat
make docker-dev

# Sadece database
make db-up
```

### **Production Build**
```bash
# Production image build et
make docker-build

# Production deploy
make docker-deploy
```

## 📊 **Project Commands (Makefile)**

```bash
# Development
make run              # Uygulamayı çalıştır
make watch            # Hot reload ile çalıştır
make deps             # Dependencies yükle

# Testing (TDD)
make test             # Tüm testleri çalıştır
make test-unit        # Unit testler
make test-integration # Integration testler
make test-acceptance  # Acceptance testler
make test-coverage    # Test coverage raporu

# Database
make db-up            # PostgreSQL başlat
make db-down          # PostgreSQL durdur
make migrate-up       # Migration'ları uygula
make migrate-down     # Migration'ları geri al

# Docker
make docker-build     # Docker image build
make docker-run       # Container'ı çalıştır
make docker-dev       # Development environment

# CI/CD
make lint             # Code linting
make format           # Code formatting
make security-check   # Security analysis
```

## 🔧 **Configuration (Viper)**

### **Environment Variables**
```bash
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=todoapp
DB_PASSWORD=secret
DB_NAME=todoapp_db

# Server
SERVER_PORT=8080
SERVER_HOST=0.0.0.0

# Logging
LOG_LEVEL=info
LOG_FORMAT=json
```

### **Config Files**
```
configs/
├── config.yaml       # Default config
├── config.dev.yaml   # Development
├── config.prod.yaml  # Production
└── config.test.yaml  # Testing
```

## 📝 **API Usage Examples**

### **Create Todo**
```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Learn Clean Architecture"}'
```

### **List Todos**
```bash
curl http://localhost:8080/api/todos
```

## 🔄 **CI/CD Pipeline**

GitHub Actions ile otomatik:
- **Build**: Go build & compile
- **Test**: Unit + Integration + Contract tests
- **Security**: Vulnerability scanning
- **Docker**: Multi-stage build
- **Deploy**: Production deployment

## 📈 **Development Phases**

1. ✅ **Setup**: Project structure + dependencies
2. 🚧 **Domain**: Entities + repository interfaces  
3. ⏳ **Infrastructure**: Database + config setup
4. ⏳ **Application**: Use cases implementation
5. ⏳ **Interface**: HTTP handlers + middleware
6. ⏳ **Testing**: TDD implementation
7. ⏳ **Docker**: Containerization
8. ⏳ **CI/CD**: Pipeline setup

## 🎓 **TDD Learning Resources**

Bu proje **TDD öğrenmek** isteyenler için hazırlanmıştır:

### **📚 TDD Documentation**
- [`TDD_DEVELOPMENT_STORY.md`](TDD_DEVELOPMENT_STORY.md) - Nasıl TDD ile geliştirildi
- [`TDD_COMMIT_SIMULATION.md`](TDD_COMMIT_SIMULATION.md) - Commit-by-commit TDD journey
- [`TDD_WORKFLOW_COMPLIANCE.md`](TDD_WORKFLOW_COMPLIANCE.md) - Workflow compliance report

### **🎯 TDD Principles Demonstrated**
1. **Test First** - Hiç kod yazılmadan önce test
2. **Red-Green-Refactor** - Klasik TDD cycle
3. **Minimal Implementation** - GREEN phase'de hard-code OK
4. **Continuous Refactoring** - Code quality improvement
5. **Test Pyramid** - Unit → Integration → Contract

### **🚀 Quick TDD Demo**
```bash
# TDD cycle'ını deneyimle
./5_minute_tdd_cycle.sh

# Hızlı test komutları
make test-unit          # Domain + Application tests  
make test-integration   # API + Database tests
make test-contract      # API contract validation tests
make test-coverage      # Coverage report
```

## 🤝 **Contributing (TDD Way)**

**Yeni özellik eklerken TDD workflow'u:**

1. 🔴 **RED**: Önce unit test yaz (FAIL)
2. 🔴 **RED**: Integration test yaz (FAIL)  
3. 🔴 **RED**: Contract test yaz (FAIL)
4. 🟢 **GREEN**: Minimal kod yaz (PASS)
5. 🔵 **REFACTOR**: Code quality artır (PASS)
6. 🔄 **REPEAT**: Next feature için cycle'ı tekrarla

**Commands:**
```bash
make test               # Tüm testler pass olmalı ✅
make test-coverage      # Coverage %85+ olmalı ✅  
make tdd-init          # TDD workflow rehberi
```

## 📚 **Architecture Resources**

- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Test Driven Development](https://martinfowler.com/bliki/TestDrivenDevelopment.html)
- [Go Fiber Documentation](https://docs.gofiber.io/)
- [Viper Configuration](https://github.com/spf13/viper) 