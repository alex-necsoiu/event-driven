# Start from the official Golang image
FROM golang:1.23.3-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire module (from module root)
COPY . .

# Build the specific service
# Debug: list generated files
RUN ls -lR api/proto/gen
RUN go build -o /user ./cmd/user

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /user .
COPY --from=builder /app/configs/user.env.example .env
EXPOSE 50051
CMD ["./user"] 