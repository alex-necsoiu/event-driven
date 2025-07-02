# Start from the official Golang image
FROM golang:1.23.3-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY api/proto/gen api/proto/gen
COPY cmd/order cmd/order
COPY internal/order internal/order
COPY configs/order.env.example configs/order.env.example
RUN go build -o /order ./cmd/order

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /order .
COPY --from=builder /app/configs/order.env.example .env
EXPOSE 50052
CMD ["./order"] 