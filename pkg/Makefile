# Build and environment variables
BINARY_NAME := opengist-cli
CMD_DIR := ./cmd/opengist-cli
BUILD_DIR := ./bin

.PHONY: all build build-all clean test run help

## build: Compiles the binary for the native host OS/ARCH
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

## build-all: Cross-compiles 5 binaries (Linux AMD64/ARM64, macOS ARM64/AMD64, Windows AMD64)
build-all:
	@mkdir -p $(BUILD_DIR)
	# Linux (AMD64)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(CMD_DIR)
	# Linux (ARM64)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(CMD_DIR)
	# macOS (Apple Silicon - ARM64)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(CMD_DIR)
	# macOS (Intel - AMD64)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(CMD_DIR)
	# Windows (AMD64)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(CMD_DIR)
	# Windows (ARM64)
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe $(CMD_DIR)
