# Standard Golang Service Template

A robust, production-ready Golang service template built with clean architecture principles, designed for building scalable REST APIs with modern development practices.

## 🚀 Overview

This project serves as a standardized template for building Go-based microservices with comprehensive features including user management, authentication, and Thai administrative data support. It demonstrates best practices in project structure, dependency management, and containerized deployment.

## ✨ Key Features

### 🏗️ Clean Architecture

- **Layered Architecture**: Separation of concerns with distinct layers
  - **Handlers**: HTTP request/response management
  - **Services**: Business logic implementation
  - **Repositories**: Data access layer
  - **Entities**: Domain models and data structures
  - **API/DTOs**: Data transfer objects for external communication

### 🔧 Technology Stack

- **Framework**: Fiber v2.52.6 (High-performance Express-inspired web framework)
- **Database**: PostgreSQL with GORM v1.25.12 ORM
- **Migration**: Liquibase 4.31 for database versioning
- **Cache**: Redis 7.4.1 for high-performance caching
- **Authentication**: JWT with golang-jwt/jwt/v5
- **Validation**: Go Playground Validator v10.26.0
- **Configuration**: Viper for flexible configuration management
- **Documentation**: Swagger with swaggo/swag integration

### 🔐 Security & Authentication

- JWT-based authentication with secure token handling
- Password hashing with bcrypt
- Middleware-based authorization
- User session management
- Secure environment configuration

### 🗺️ Thai Administrative Data

Built-in support for Thai geographic data:

- **Provinces** (จังหวัด)
- **Districts** (อำเภอ/เขต)
- **Sub-districts** (ตำบล/แขวง)

### 🐳 DevOps & Deployment

- **Docker Compose**: Complete containerized development environment
- **Health Checks**: Built-in service monitoring
- **Graceful Shutdown**: Proper resource cleanup
- **Environment Management**: Flexible configuration per environment

## 🛠️ Getting Started

### Prerequisites

- **Go**: 1.23.2 or higher
- **Docker**: Latest version
- **Docker Compose**: Latest version

### Quick Start

1. **Clone the repository**

```bash
git clone https://github.com/wisaitas/standard-golang.git
cd standard-golang
```

2. **Install dependencies**

```bash
go mod tidy
```

3. **Start the infrastructure**

```bash
docker compose up -d
```

4. **Run the application locally**

```bash
go run cmd/standard-service/main.go
```

The application will be available at:

- **API Server**: http://localhost:8005
- **PostgreSQL**: localhost:9000
- **Redis**: localhost:9001

## 🔌 API Endpoints

### Authentication

- `POST /api/v1/auth/login` - User authentication
- `POST /api/v1/auth/refresh` - Token refresh
- `POST /api/v1/auth/logout` - User logout

### User Management

- `GET /api/v1/users` - List users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users/{id}` - Get user by ID
- `PUT /api/v1/users/{id}` - Update user
- `DELETE /api/v1/users/{id}` - Delete user

### Thai Administrative Data

- `GET /api/v1/provinces` - List provinces
- `GET /api/v1/districts` - List districts
- `GET /api/v1/sub-districts` - List sub-districts

## ⚙️ Configuration

Configure the application through environment variables in `deployment/env/api.env`:

```env
# Server Configuration
SERVER_ENV=dev
SERVER_PORT=8005
SERVER_MAX_FILE_SIZE=5
SERVER_JWT_SECRET=your-secret-key

# Database Configuration
DATABASE_HOST=db
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=root
DATABASE_NAME=postgres
DATABASE_DRIVER=postgres

# Redis Configuration
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=
```

## 🧪 Testing

The project includes comprehensive testing utilities:

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Generate HTML coverage report
make test-coverage-html

# Check coverage threshold (60% minimum)
make test-coverage-check

# Clean coverage files
make clean-coverage
```

## 🏛️ Architecture Principles

### Clean Architecture

- **Dependency Inversion**: High-level modules don't depend on low-level modules
- **Separation of Concerns**: Each layer has a single responsibility
- **Testability**: Easy to unit test with dependency injection
- **Maintainability**: Clear boundaries between business logic and infrastructure

### Repository Pattern

- Abstracts data access logic
- Enables easy testing with mock repositories
- Supports multiple data sources

### Service Layer

- Contains business logic
- Orchestrates between repositories and external services
- Maintains transaction boundaries

## 🔒 Security Features

- **JWT Authentication**: Stateless token-based authentication
- **Password Security**: bcrypt hashing with salt
- **Middleware Protection**: Route-level authorization
- **Environment Security**: Sensitive data in environment variables
- **Input Validation**: Comprehensive request validation

## 🐳 Docker Support

The project includes complete Docker support:

- **Multi-stage builds** for optimized production images
- **Health checks** for service reliability
- **Volume management** for data persistence
- **Service orchestration** with Docker Compose

## 📊 Monitoring & Health Checks

- Database connectivity monitoring
- Redis connectivity monitoring
- Graceful shutdown handling
- Resource cleanup on termination

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙋‍♂️ Support

For questions and support:

- Open an issue in the GitHub repository
- Contact the maintainers

---

**Built with ❤️ in Go**
