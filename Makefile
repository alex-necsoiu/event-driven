PROTO_DIR=api/proto
GEN_DIR=$(PROTO_DIR)/gen

PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)

.PHONY: proto
proto:
	protoc -I=$(PROTO_DIR) --go_out=$(GEN_DIR) --go-grpc_out=$(GEN_DIR) $(PROTO_DIR)/*.proto

.PHONY: clean
clean:
	rm -rf $(GEN_DIR)/*
	rm -f coverage.out

.PHONY: test
test:
	go test ./... -v

.PHONY: test-short
test-short:
	go test ./... -v -short

.PHONY: cover
cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
	@echo "Coverage summary:"
	@go tool cover -func=coverage.out | grep total

.PHONY: cover-check
cover-check:
	go test ./... -coverprofile=coverage.out
	@echo "Checking coverage threshold (80%)..."
	@coverage=$$(go tool cover -func=coverage.out | grep total | awk '{gsub(/%/, "", $$3); print $$3}'); \
	if [ $$(echo "$$coverage < 80" | bc -l) -eq 1 ]; then \
		echo "❌ Coverage too low: $$coverage% (minimum: 80%)"; \
		exit 1; \
	else \
		echo "✅ Coverage: $$coverage% (minimum: 80%)"; \
	fi

.PHONY: cover-clean
cover-clean:
	rm -f coverage.out coverage.html

.PHONY: build
build:
	go build ./cmd/...

.PHONY: build-all
build-all:
	go build ./cmd/gateway
	go build ./cmd/user
	go build ./cmd/order
	go build ./cmd/notification

.PHONY: docker-build
docker-build:
	docker-compose build

.PHONY: docker-up
docker-up:
	docker-compose up -d

.PHONY: docker-down
docker-down:
	docker-compose down

.PHONY: docker-logs
docker-logs:
	docker-compose logs -f

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  proto        - Generate protobuf code"
	@echo "  clean        - Clean generated files"
	@echo "  test         - Run all tests with verbose output"
	@echo "  test-short   - Run tests with -short flag"
	@echo "  cover        - Run tests with coverage and generate HTML report"
	@echo "  cover-check  - Run tests and check coverage threshold (80%)"
	@echo "  cover-clean  - Clean coverage files"
	@echo "  build        - Build all services"
	@echo "  build-all    - Build each service individually"
	@echo "  docker-build - Build Docker images"
	@echo "  docker-up    - Start all services"
	@echo "  docker-down  - Stop all services"
	@echo "  docker-logs  - Show Docker logs"
	@echo "  help         - Show this help message" 