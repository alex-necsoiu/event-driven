# Start from the official Golang image
FROM golang:1.23.3-alpine as builder

# Install git for fetching from GitHub
RUN apk add --no-cache git

WORKDIR /app

# Create go.mod for the service
RUN go mod init order-service

# Add the required dependencies
RUN go get github.com/alex-necsoiu/event-driven@v1.0.10
RUN go get github.com/joho/godotenv@v1.5.1
RUN go get github.com/lib/pq@v1.10.9
RUN go get github.com/nats-io/nats.go@v1.43.0
RUN go get google.golang.org/grpc@v1.73.0
RUN go get google.golang.org/protobuf@v1.36.6

# Copy config files
COPY configs/order.env.example /app/configs/order.env.example

# Build the service
RUN go build -o /order github.com/alex-necsoiu/event-driven/cmd/order

FROM alpine:latest
WORKDIR /root/

# Copy the binary
COPY --from=builder /order .

# Copy environment file from the root context
COPY --from=builder /app/configs/order.env.example .env

EXPOSE 50052
CMD ["./order"]    