
## Architecture Overview

This project follows clean architecture principles with clear separation of concerns:

1. **Handlers Layer**: Handles HTTP requests/responses
2. **Services Layer**: Contains business logic
3. **Repositories Layer**: Manages data access
4. **Models Layer**: Defines data structures
5. **DTOs Layer**: Manages data transfer objects

### Key Benefits of This Structure

1. **Separation of Concerns**
   - Clear boundaries between different layers
   - Easy to maintain and modify individual components
   - Better testability

2. **Dependency Injection**
   - Loose coupling between components
   - Easy to mock dependencies for testing
   - Better control over object creation

3. **Middleware Architecture**
   - Centralized middleware management
   - Easy to add/remove global middlewares
   - Configurable per-route middleware

4. **Configuration Management**
   - Environment-based configuration
   - Centralized configuration handling
   - Easy to switch between different environments

5. **Testing Support**
   - Dedicated mocks directory
   - Easy to write unit tests
   - Clear structure for test organization

## Features

- JWT Authentication
- Redis Session Management
- PostgreSQL Database
- Request Rate Limiting
- CORS Support
- Health Check Endpoints
- Structured Logging
- Request Validation
- Graceful Shutdown

## Getting Started

1. Clone the repository
2. Install dependencies: