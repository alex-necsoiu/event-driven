# Start from the official Golang image
FROM golang:1.23.3-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire module (from module root)
COPY . .

# Build the specific service
RUN go build -o /order ./cmd/order

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /order .
COPY --from=builder /app/configs/order.env.example .env
EXPOSE 50052
CMD ["./order"] 

RUN go install github.com/alex-necsoiu/event-driven/cmd/order@v1.0.0    