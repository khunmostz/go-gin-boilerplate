# Go Gin Boilerplate

[![Go Version](https://img.shields.io/badge/Go-1.23+-blue.svg)](https://golang.org)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.10.1-green.svg)](https://gin-gonic.com)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

🚀 **Go Gin Boilerplate** เป็น template สำหรับการสร้างแอปพลิเคชัน API backend ด้วย Go และ Gin framework ที่มาพร้อมกับ Clean Architecture, Docker support, และ features ที่จำเป็นสำหรับการพัฒนาแอปพลิเคชันสมัยใหม่

## ✨ Features

- 🔥 **Gin Framework** - Fast HTTP web framework
- 🏗️ **Clean Architecture** - Domain, Repository, Service, Handler layers
- 🐳 **Docker Support** - Development และ Production ready
- 🗄️ **Multi Database** - PostgreSQL และ MongoDB support
- ⚡ **Redis Cache** - In-memory caching
- 🔐 **JWT Authentication** - Secure authentication system
- 📚 **Swagger Documentation** - Auto-generated API docs
- 🧪 **Testing Ready** - Unit tests และ mocking support
- ⚙️ **Configuration Management** - YAML-based configs
- 🔄 **Hot Reload** - Development mode with live reload
- 🛠️ **Make Commands** - Easy development workflow

## 📁 Project Structure

```
go-gin-boilerplate/
├── cmd/                    # Application entry points
│   └── main.go            # Main application
├── config/                # Configuration files
│   ├── config.go          # Config loader
│   ├── dev/               # Development configs
│   ├── prod/              # Production configs
│   └── example/           # Example configs
├── internal/              # Internal application code
│   ├── cache/             # Redis cache implementation
│   ├── db/                # Database connections
│   ├── domain/            # Domain entities
│   ├── handler/           # HTTP handlers
│   │   └── api/           # API route handlers
│   ├── middleware/        # HTTP middlewares
│   ├── port/              # Interfaces/Ports
│   ├── repository/        # Data access layer
│   ├── service/           # Business logic layer
│   ├── tests/             # Unit tests
│   └── utils/             # Utility functions
├── docs/                  # Swagger documentation
├── docker-compose.*.yaml  # Docker compose files
├── Dockerfile.*           # Docker configurations
└── Makefile              # Development commands
```

## 🚀 Quick Start

### Prerequisites

- Go 1.23+
- Docker & Docker Compose
- Make (optional)

### 1. Clone และ Setup

```bash
# Clone repository
git clone <your-repo-url>
cd go-gin-boilerplate

# Copy example config
cp config/example/config.example.yaml config/dev/config.dev.yaml

# Edit config file ตามต้องการ
vim config/dev/config.dev.yaml
```

### 2. Development Mode

```bash
# Start databases และ services
make dev

# หรือใช้ docker compose โดยตรง
docker compose -f docker-compose.dev.yaml -f docker-compose.override.yaml up --build
```

### 3. Production Mode

```bash
# Start production environment
make prod

# หรือ
docker compose -f docker-compose.prod.yaml up -d --build
```

## 🔧 Configuration

สร้างไฟล์ config ใน `config/dev/config.dev.yaml`:

```yaml
app:
  env: development

server:
  port: "8080"

database:
  type: postgresql  # หรือ mongodb
  postgres:
    host: localhost
    port: "5432"
    username: postgres
    password: password
    dbname: your_db
    sslmode: disable
    timezone: Asia/Bangkok
  mongodb:
    uri: ""
    host: localhost
    port: "27017"
    username: ""
    password: ""
    dbname: your_db

jwt:
  secret: "your-jwt-secret-key"
  access_duration: "15m"
  refresh_duration: "7d"

redis:
  host: localhost
  port: "6379"
  password: ""
```

## 📖 API Documentation

หลังจากรันแอป เข้าไปดู Swagger documentation ได้ที่:

```
http://localhost:8080/swagger/index.html
```

### API Endpoints

- `GET /api/v1/health` - Health check
- `POST /api/v1/auth/login` - User login
- `GET /api/v1/foo` - Get foo items
- `POST /api/v1/foo` - Create foo item
- `GET /api/v1/bar` - Get bar items (Authentication required)
- `POST /api/v1/bar` - Create bar item (Authentication required)

## 🛠️ Development Commands

### Make Commands

```bash
make help              # Show all available commands
make dev               # Start development environment
make dev-detached      # Start development in background
make down              # Stop all services
make clean             # Remove containers, volumes, images
make logs              # Show development logs
make build             # Build Go application
make test              # Run tests
make db-dev            # Start only databases
```

### Manual Commands

```bash
# Build application
go build -o bin/main cmd/main.go

# Run tests
go test ./...

# Generate swagger docs
swag init -g cmd/main.go -o docs

# Run application locally
go run cmd/main.go
```

## 🧪 Testing

```bash
# Run all tests
make test

# Run specific package tests
go test ./internal/service/...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## 🗄️ Database

### PostgreSQL

Default configuration สำหรับ development:
- Host: localhost:5432
- Username: postgres
- Password: password
- Database: chatbot_dev

### MongoDB

Default configuration สำหรับ development:
- Host: localhost:27017
- Database: chatbot_dev

### Redis

Default configuration สำหรับ development:
- Host: localhost:6379

## 🔐 Authentication

ใช้ JWT (JSON Web Token) สำหรับ authentication:

1. Login ผ่าน `/api/v1/auth/login`
2. ใช้ access token ใน Authorization header: `Bearer <token>`
3. Token จะหมดอายุตาม config (`access_duration`)

## 🐳 Docker

### Development

```bash
# Start all services
docker compose -f docker-compose.dev.yaml -f docker-compose.override.yaml up

# Start only databases
docker compose -f docker-compose.dev.yaml up pgsql redis mongodb

# View logs
docker compose logs -f app
```

### Production

```bash
# Start production
docker compose -f docker-compose.prod.yaml up -d

# Scale application
docker compose -f docker-compose.prod.yaml up -d --scale app=3
```

## 📝 Creating New Features

### 1. สร้าง Domain Entity

```go
// internal/domain/user.go
package domain

type User struct {
    ID    string `json:"id" bson:"_id" gorm:"primaryKey"`
    Name  string `json:"name" bson:"name" gorm:"column:name"`
    Email string `json:"email" bson:"email" gorm:"column:email"`
}
```

### 2. สร้าง Repository Interface

```go
// internal/port/user.go
package port

import "your-app/internal/domain"

type UserRepository interface {
    Create(user *domain.User) error
    GetByID(id string) (*domain.User, error)
    GetAll() ([]*domain.User, error)
    Update(user *domain.User) error
    Delete(id string) error
}
```

### 3. Implement Repository

```go
// internal/repository/user.go
package repository

import (
    "your-app/internal/db"
    "your-app/internal/domain"
    "your-app/internal/port"
)

type userRepository struct {
    db         db.BaseRepository
    collection string
}

func NewUserRepository(db db.BaseRepository, collection string) port.UserRepository {
    return &userRepository{
        db:         db,
        collection: collection,
    }
}

func (r *userRepository) Create(user *domain.User) error {
    return r.db.Create(r.collection, user)
}

// ... implement other methods
```

### 4. สร้าง Service

```go
// internal/service/user.go
package service

import (
    "your-app/internal/domain"
    "your-app/internal/port"
)

type UserService struct {
    userRepo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) *UserService {
    return &UserService{
        userRepo: userRepo,
    }
}

func (s *UserService) CreateUser(user *domain.User) error {
    // Business logic here
    return s.userRepo.Create(user)
}
```

### 5. สร้าง Handler

```go
// internal/handler/user.go
package handler

import (
    "net/http"
    "your-app/internal/service"
    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{
        userService: userService,
    }
}

// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.User true "User object"
// @Success 201 {object} domain.User
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
    // Handler logic here
}
```

### 6. Register Routes

```go
// internal/handler/api/user.go
package api

import (
    "your-app/internal/handler"
    "github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler) {
    router.POST("/", userHandler.CreateUser)
    router.GET("/:id", userHandler.GetUser)
    router.GET("/", userHandler.GetUsers)
    router.PUT("/:id", userHandler.UpdateUser)
    router.DELETE("/:id", userHandler.DeleteUser)
}
```

### 7. Add to main.go

```go
// cmd/main.go
userRouter := apiRouter.Group("/users")
{
    userRepo := repository.NewUserRepository(baseRepo, "users")
    userSvc := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userSvc)
    api.RegisterUserRoutes(userRouter, userHandler)
}
```

## 🔄 Using as Template

### สำหรับโปรเจคใหม่:

1. **Fork หรือ Clone repository นี้**
2. **เปลี่ยนชื่อ module ใน go.mod**
   ```go
   module your-new-project-name
   ```
3. **Update import paths ทั้งหมด**
4. **แก้ไข config files**
5. **เปลี่ยน database names และ credentials**
6. **แก้ไข Swagger info ใน main.go**
7. **เพิ่ม features ตามต้องการ**

### Script สำหรับการ setup ใหม่:

```bash
#!/bin/bash
# setup-new-project.sh

OLD_MODULE="go-gin-boilerplate"
NEW_MODULE="$1"

if [ -z "$NEW_MODULE" ]; then
    echo "Usage: ./setup-new-project.sh <new-module-name>"
    exit 1
fi

# Update go.mod
sed -i "s/$OLD_MODULE/$NEW_MODULE/g" go.mod

# Update all Go files
find . -name "*.go" -type f -exec sed -i "s/$OLD_MODULE/$NEW_MODULE/g" {} +

echo "Project setup complete for: $NEW_MODULE"
echo "Don't forget to:"
echo "1. Update config files"
echo "2. Change database names"
echo "3. Update Swagger info"
echo "4. Initialize git repository"
```

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [MongoDB Go Driver](https://go.mongodb.org/mongo-driver/)
- [Redis Go Client](https://github.com/go-redis/redis)
- [Viper](https://github.com/spf13/viper)
- [Swagger](https://swagger.io/)

---

Made with ❤️ for Go developers