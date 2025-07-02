# Start from the official Golang image
FROM golang:1.23.3-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY api/proto/gen api/proto/gen
COPY cmd/gateway cmd/gateway
COPY internal/gateway internal/gateway
COPY configs/gateway.env.example configs/gateway.env.example
RUN go build -o /gateway ./cmd/gateway

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /gateway .
COPY --from=builder /app/configs/gateway.env.example .env
EXPOSE 8080
CMD ["./gateway"] 