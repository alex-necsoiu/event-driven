# Start from the official Golang image
FROM golang:1.23.3-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY api/proto/gen api/proto/gen
COPY cmd/user cmd/user
COPY internal/user internal/user
COPY configs/user.env.example configs/user.env.example
# Debug: list generated files
RUN ls -lR api/proto/gen
RUN go build -o /user ./cmd/user

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /user .
COPY --from=builder /app/configs/user.env.example .env
EXPOSE 50051
CMD ["./user"] 