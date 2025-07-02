# Start from the official Golang image
FROM golang:1.23.3-alpine as builder
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire module (from module root)
COPY . .

# Build the specific service
RUN go build -o /gateway ./cmd/gateway

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /gateway .
COPY --from=builder /app/configs/gateway.env.example .env
EXPOSE 8080
CMD ["./gateway"] 