# Aturan Makefile
.PHONY: proto clean

# Generate proto
proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	       --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	       proto/*.proto

# Bersihkan hasil generate
clean:
	rm -rf pb/*

# Aturan default
.DEFAULT_GOAL := proto
