# Start from the official Golang image
FROM golang:1.23.3-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY api/proto/gen api/proto/gen
COPY cmd/notification cmd/notification
COPY internal/notification internal/notification
COPY configs/notification.env.example configs/notification.env.example
RUN go build -o /notification ./cmd/notification

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /notification .
COPY --from=builder /app/configs/notification.env.example .env
CMD ["./notification"]
RUN go install github.com/alex-necsoiu/event-driven/cmd/order@v1.0.0 