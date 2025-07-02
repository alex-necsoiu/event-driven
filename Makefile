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