.PHONY: proto clean

# Define all proto files
PROTO_FILES := \
	proto/contracts/v2/hello.proto \
	proto/contracts/v2/service.proto

# Generate proto
proto:
	protoc --proto_path=proto --proto_path=googleapis \
	       --go_out=proto/contracts/v2 --go_opt=paths=source_relative \
	       --go-grpc_out=proto/contracts/v2 --go-grpc_opt=paths=source_relative \
	       --grpc-gateway_out=proto/contracts/v2 --grpc-gateway_opt=paths=source_relative \
	       $(PROTO_FILES)

# Clean generated files
clean:
	rm -rf proto/contracts/v2/*

# Default goal
.DEFAULT_GOAL := proto
