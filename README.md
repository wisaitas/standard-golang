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
- **Database**: PostgreSQL 17
- **Cache**: Redis 7.4.1
- **ORM**: GORM
- **Authentication**: JWT-based authentication
- **Validation**: Go Playground Validator

### 3. Developer Experience
- **Docker Support**: Complete containerization setup
- **Postman Collection**: Ready-to-use API documentation
- **Environment Management**: Flexible configuration system
- **Health Checks**: Built-in system monitoring
- **Rate Limiting**: Request throttling support
- **CORS**: Cross-origin resource sharing enabled

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
- Make (optional)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/wisatas/standard-golang.git
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
Import the Postman collection from `postman-collection/Standard Golang.postman_collection.json` to get started with the API endpoints.

## 📁 Project Structure

```
.
├── cmd/                    # Application entry point
├── data/                   # JSON data files
├── deployment/             # Deployment configurations
│   ├── docker-images/     # Dockerfile definitions
│   └── env/               # Environment configurations
├── internal/              # Internal application code
│   ├── handlers/         # HTTP request handlers
│   ├── services/         # Business logic
│   ├── repositories/     # Data access layer
│   ├── models/          # Data models
│   └── dto/             # Data transfer objects
└── postman-collection/   # API documentation
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

## 📝 Testing

Run tests using:
```bash
go test ./...
```

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