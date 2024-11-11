# Set the directory for your Proto files
PROTO_DIR := proto

# Install necessary tools for protobuf compilation
install-tools:
	@echo "Installing necessary tools..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "Tools installed successfully."

# Compile proto files
generate-proto:
	@echo "Generating gRPC code..."
	protoc -I=$(PROTO_DIR) \
		--go_out=. \
		--go-grpc_out=. \
		$(PROTO_DIR)/**/*.proto

# Download and tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	go mod tidy

# Build the application
build:
	go build -o main ./cmd/main.go

run-app:
	go run ./cmd/

# Run the entire pipeline: install tools, generate protobuf, tidy, and build
all: install-tools generate-proto tidy run-app
