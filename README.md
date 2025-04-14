# Standard Golang Project Template

A robust and well-structured Golang project template following clean architecture principles, designed for building scalable and maintainable applications.

## 🌟 Key Features

### 1. Clean Architecture Implementation

- **Handlers Layer**: HTTP request/response handling
- **Services Layer**: Business logic implementation
- **Repositories Layer**: Data access management
- **Models Layer**: Data structure definitions
- **DTOs Layer**: Data transfer object management

### 2. Modern Tech Stack

- **Framework**: Fiber v2 (High-performance web framework)
- **Database**: PostgreSQL with Liquibase for migrations
- **Cache**: Redis 7.4.1
- **ORM**: GORM
- **Authentication**: JWT-based authentication
- **Validation**: Go Playground Validator v10

### 3. Developer Experience

- **Docker Support**: Complete containerization setup
- **Postman Collection**: Ready-to-use API documentation
- **Environment Management**: Flexible configuration with caarlos0/env
- **Health Checks**: Built-in system monitoring
- **JWT**: Implementation using golang-jwt/jwt/v5
- **Mocks**: Testing utilities

### 4. Security Features

- JWT Authentication with refresh tokens
- Password hashing
- Rate limiting protection
- CORS security
- Environment-based configuration

### 5. Data Management

- Built-in support for Thai administrative data
  - Provinces
  - Districts
  - Sub-districts
- JSON data structure for easy maintenance

## 🚀 Getting Started

### Prerequisites

- Go 1.23.2 or higher
- Docker and Docker Compose

### Installation

1. Clone the repository:

```bash
git clone https://github.com/wisaitas/standard-golang.git
cd standard-golang
```

2. Install dependencies:

```bash
go mod tidy
```

3. Set up environment variables:

```bash
cp deployment/env/api.env.example deployment/env/api.env
```

4. Start the application:

```bash
docker compose up -d --build
```

The application will be available at:

- API: http://localhost:8082
- PostgreSQL: localhost:8080
- Redis: localhost:8081

### API Documentation

Import the Postman collection from `postman-collection` to get started with the API endpoints.

## 📁 Project Structure

```

├── cmd/ # Application entry points
├── deployment/ # Deployment configurations
│ ├── docker-images/ # Dockerfile definitions
│ └── env/ # Environment configurations
├── internal/ # Private application code
│ ├── configs/ # Application configurations
│ ├── constants/ # Constant definitions
│ ├── dtos/ # Data transfer objects
│ ├── env/ # Environment variable handling
│ ├── handlers/ # HTTP request handlers
│ ├── initial/ # Application initialization
│ ├── middlewares/ # HTTP middleware components
│ ├── models/ # Data models
│ ├── mocks/ # Mock objects for testing
│ ├── repositories/ # Data access layer
│ ├── routes/ # API route definitions
│ ├── services/ # Business logic implementation
│ ├── utils/ # Utility functions
│ └── validates/ # Request validation logic
├── pkg/ # Public libraries/packages
│ ├── bcrypt.go # Password encryption utilities
│ ├── error.go # Error handling utilities
│ ├── jwt.go # JWT authentication utilities
│ ├── model.go # Model-related utilities
│ ├── query.go # Query building utilities
│ ├── redis.go # Redis client utilities
│ ├── repository.go # Repository pattern utilities
│ ├── response.go # HTTP response utilities
│ ├── transaction.go # Database transaction utilities
│ └── validator.go # Validation utilities
└── postman-collection/ # API documentation
```

## 📝 Configuration

The application can be configured through environment variables in `deployment/env/api.env`:

```env
PORT=8082
DB_HOST=standard_db
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=postgres
DB_PORT=5432
JWT_SECRET=secret
REDIS_HOST=standard_redis
REDIS_PORT=6379
```

## 🔗 Technology Stack

- **Go**: v1.23.2
- **Web Framework**: Fiber v2.52.6
- **ORM**: GORM v1.25.12
- **Database**: PostgreSQL with GORM postgres driver v1.5.11
- **Cache**: Redis v9.7.0
- **Authentication**: JWT v5.2.1
- **Validation**: Go Playground Validator v10.24.0
- **Environment**: caarlos0/env/v11
- **Testing**: stretchr/testify v1.10.0

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📞 Support

For support, please open an issue in the GitHub repository or contact the maintainers.
