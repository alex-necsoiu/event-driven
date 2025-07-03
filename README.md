# 🚀 Event-Driven Microservices System

Go Version License: MIT

A modern, scalable microservices architecture built with Go, featuring event-driven communication, gRPC APIs, REST gateways, and clean architecture principles. This system demonstrates best practices for building distributed applications with loose coupling and high scalability.

## ✨ Features

* **Event-Driven Architecture**: Asynchronous communication between services using NATS messaging
* **gRPC + REST APIs**: Internal service-to-service communication via gRPC, external APIs via REST
* **Clean Architecture**: Layered design with clear separation of concerns
* **Microservices**: Modular services (User, Order, Notification, Gateway) with independent deployment
* **Docker & Docker Compose**: Containerized development and deployment
* **PostgreSQL Integration**: Reliable data persistence with automatic migrations
* **Structured Logging**: Comprehensive logging for monitoring and debugging
* **Environment Configuration**: Flexible configuration management per service
* **CI/CD Ready**: GitHub Actions workflow structure for automated testing and deployment

## 🚀 Getting Started

Follow these instructions to get the project up and running on your local machine.

### Prerequisites

* Go version 1.23 or newer
* Docker and Docker Compose
* Protocol Buffers compiler (`protoc`)
* PostgreSQL (or use Docker Compose)

### Installation & Execution

1. **Clone the repository:**
   ```bash
   git clone https://github.com/alex-necsoiu/event-driven.git
   cd event-driven
   ```

2. **Install Protocol Buffers tools:**
   ```bash
   # macOS
   brew install protobuf
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   
   # Linux
   sudo apt-get install protobuf-compiler
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

3. **Generate Protocol Buffers:**
   ```bash
   make proto
   ```

4. **Start the entire system:**
   ```bash
   docker-compose up
   ```

### ✅ Running Tests

To ensure everything is working as expected, run the test suite:

```bash
go test ./...
```

## 🏗️ Architecture

The application follows a clean, layered architecture with event-driven communication, promoting loose coupling and high scalability.

```
.
├── api/                    # Protocol Buffers definitions and generated code
│   └── proto/
│       ├── gen/           # Generated Go code from .proto files
│       ├── event.proto    # Event message definitions
│       ├── user.proto     # User service gRPC definitions
│       ├── order.proto    # Order service gRPC definitions
│       └── notification.proto # Notification service definitions
├── cmd/                   # Application entry points
│   ├── gateway/           # REST API gateway service
│   ├── user/              # User management service
│   ├── order/             # Order processing service
│   └── notification/      # Event notification service
├── internal/              # Private application code
│   ├── gateway/           # Gateway service implementation
│   ├── user/              # User service business logic
│   ├── order/             # Order service business logic
│   └── notification/      # Notification service business logic
├── pkg/                   # Shared, reusable packages
│   └── messaging/         # Event messaging utilities
├── configs/               # Configuration files
├── docker/                # Docker-related files
├── test/                  # Test files and mocks
└── docker-compose.yml     # Multi-service orchestration
```

### Service Architecture

Each service follows the same architectural pattern:

* **`cmd/<service>/`**: Entry point with main function and service initialization
* **`internal/<service>/`**: Business logic, handlers, repositories, and services
* **`config.go`**: Configuration loading and validation
* **`handler.go`**: gRPC/HTTP request handlers
* **`service.go`**: Core business logic and event publishing
* **`repository.go`**: Data access layer (PostgreSQL)

### Communication Flow

```
┌─────────────┐    REST     ┌─────────────┐    gRPC     ┌─────────────┐
│   Client    │ ──────────► │   Gateway   │ ──────────► │    User     │
│             │             │             │             │   Service   │
└─────────────┘             └─────────────┘             └─────────────┘
                                    │                           │
                                    │ Events                    │ Events
                                    ▼                           ▼
                            ┌─────────────┐             ┌─────────────┐
                            │    NATS     │ ◄────────── │   Order     │
                            │  Messaging  │             │  Service    │
                            └─────────────┘             └─────────────┘
                                    │
                                    │ Events
                                    ▼
                            ┌─────────────┐
                            │Notification│
                            │  Service    │
                            └─────────────┘
```

## 🔧 Services Overview

### Gateway Service (`cmd/gateway/`)

**Purpose**: REST API gateway that routes requests to appropriate microservices.

**Features**:
* RESTful HTTP endpoints for external clients
* Request routing to internal gRPC services
* Request/response transformation
* Load balancing and service discovery ready

**Endpoints**:
* `POST /users` - Create new user
* `GET /users/{id}` - Get user by ID
* `POST /orders` - Create new order
* `GET /orders/{id}` - Get order by ID

### User Service (`cmd/user/`)

**Purpose**: Manages user accounts and authentication.

**Features**:
* User CRUD operations
* gRPC API for internal communication
* Event publishing for user lifecycle events
* PostgreSQL data persistence

**Events Published**:
* `UserCreated` - When a new user is registered
* `UserUpdated` - When user information is modified
* `UserDeleted` - When a user account is removed

### Order Service (`cmd/order/`)

**Purpose**: Handles order processing and management.

**Features**:
* Order CRUD operations
* Order status management
* gRPC API for internal communication
* Event publishing for order lifecycle events
* PostgreSQL data persistence

**Events Published**:
* `OrderCreated` - When a new order is placed
* `OrderStatusUpdated` - When order status changes
* `OrderCancelled` - When an order is cancelled

### Notification Service (`cmd/notification/`)

**Purpose**: Processes events and sends notifications.

**Features**:
* Event subscription and processing
* Notification delivery (email, SMS, push)
* Event filtering and routing
* Background processing

**Events Consumed**:
* All events from User and Order services
* Processes events asynchronously
* Sends appropriate notifications based on event type

## 🔧 Extensibility

The architecture makes it straightforward to add new functionality.

### Adding a New Service

1. **Create service structure**:
   ```bash
   mkdir -p cmd/newservice internal/newservice
   ```

2. **Define Protocol Buffers** in `api/proto/newservice.proto`:
   ```protobuf
   syntax = "proto3";
   package newservice;
   option go_package = "github.com/alex-necsoiu/event-driven/api/proto/gen";
   
   service NewService {
     rpc CreateItem(CreateItemRequest) returns (CreateItemResponse);
   }
   ```

3. **Generate code**:
   ```bash
   make proto
   ```

4. **Implement service logic** in `internal/newservice/`:
   * `service.go` - Business logic
   * `handler.go` - gRPC handlers
   * `repository.go` - Data access

5. **Add to Docker Compose** in `docker-compose.yml`

### Adding New Events

1. **Define event in `api/proto/event.proto`**:
   ```protobuf
   message NewEvent {
     string event_type = 1;
     bytes payload = 2;
     string timestamp = 3;
   }
   ```

2. **Add event types in `pkg/messaging/events.go`**:
   ```go
   const (
       EventTypeNewEvent = "NewEvent"
   )
   ```

3. **Publish events in services**:
   ```go
   event := &messaging.Event{
       Type:    messaging.EventTypeNewEvent,
       Payload: payload,
   }
   messaging.PublishEvent(event)
   ```

4. **Subscribe to events** in notification service:
   ```go
   messaging.SubscribeToEvent(messaging.EventTypeNewEvent, handler)
   ```

## 🐳 Docker & Deployment

### Local Development

```bash
# Build all services
docker-compose build

# Start the entire system
docker-compose up

# Start specific service
docker-compose up user-service

# View logs
docker-compose logs -f gateway
```

### Production Deployment

Each service has its own Dockerfile optimized for production:

```bash
# Build production images
docker build -t event-driven-gateway:latest cmd/gateway/
docker build -t event-driven-user:latest cmd/user/
docker build -t event-driven-order:latest cmd/order/
docker build -t event-driven-notification:latest cmd/notification/
```

## 🔧 Configuration

### Environment Variables

Each service loads configuration from environment variables:

**Gateway Service**:
```bash
GATEWAY_PORT=8080
USER_SERVICE_URL=localhost:50051
ORDER_SERVICE_URL=localhost:50052
```

**User Service**:
```bash
USER_PORT=50051
USER_DB_CONN=postgres://user:pass@localhost:5432/users?sslmode=disable
NATS_URL=nats://localhost:4222
```

**Order Service**:
```bash
ORDER_PORT=50052
ORDER_DB_CONN=postgres://user:pass@localhost:5432/orders?sslmode=disable
NATS_URL=nats://localhost:4222
```

**Notification Service**:
```bash
NOTIFICATION_PORT=50053
NATS_URL=nats://localhost:4222
```

### Database Setup

The system uses PostgreSQL for data persistence:

```bash
# Run migrations
psql -h localhost -U user -d users -f docker/migrate.sql
psql -h localhost -U user -d orders -f docker/migrate.sql
```

## 🧪 Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific service tests
go test ./internal/user/...
```

### Test Structure

```
test/
├── mocks/           # Mock implementations
├── integration/     # Integration tests
└── unit/           # Unit tests
```

## 📊 Monitoring & Logging

### Structured Logging

All services use structured logging for better observability:

```go
log.Printf("User created: %s", user.ID)
log.Printf("Order processed: %s, status: %s", order.ID, order.Status)
```

### Health Checks

Each service exposes health check endpoints:

```bash
# Gateway health check
curl http://localhost:8080/health

# gRPC health checks
grpc_health_probe -addr=localhost:50051
```

## 🚀 CI/CD Pipeline

The project includes GitHub Actions workflow structure:

```yaml
# .github/workflows/ci.yml
name: CI/CD Pipeline
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
      - run: go test ./...
      - run: go build ./...
```

## 📝 API Documentation

### REST API (Gateway)

**Create User**:
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

**Get User**:
```bash
curl http://localhost:8080/users/123
```

**Create Order**:
```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"user_id": "123", "items": [{"product_id": "456", "quantity": 2}]}'
```

### gRPC APIs

**User Service** (port 50051):
```protobuf
service UserService {
  rpc CreateUser(CreateUserRequest) returns (CreateUserResponse);
  rpc GetUser(GetUserRequest) returns (GetUserResponse);
  rpc UpdateUser(UpdateUserRequest) returns (UpdateUserResponse);
  rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse);
}
```

**Order Service** (port 50052):
```protobuf
service OrderService {
  rpc CreateOrder(CreateOrderRequest) returns (CreateOrderResponse);
  rpc GetOrder(GetOrderRequest) returns (GetOrderResponse);
  rpc UpdateOrderStatus(UpdateOrderStatusRequest) returns (UpdateOrderStatusResponse);
}
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License. See the LICENSE file for details.

## 🆘 Troubleshooting

### Common Issues

**Port conflicts**:
```bash
# Check what's using port 5432
lsof -i :5432
# Stop conflicting containers
docker stop $(docker ps -q)
```

**Proto generation errors**:
```bash
# Ensure protoc is installed
protoc --version
# Regenerate proto files
make proto
```

**Docker build failures**:
```bash
# Clean Docker cache
docker system prune -a
# Rebuild without cache
docker-compose build --no-cache
```

## 📚 Additional Resources

* [gRPC Documentation](https://grpc.io/docs/)
* [NATS Documentation](https://docs.nats.io/)
* [Protocol Buffers Guide](https://developers.google.com/protocol-buffers)
* [Docker Compose Reference](https://docs.docker.com/compose/)
* [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
