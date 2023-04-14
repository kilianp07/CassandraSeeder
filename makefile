# Go parameters
GO := go
GOFLAGS := 
GOTEST := $(GO) test
BINARY_NAME := CassandraSeeder
BUILD_DIR := ./build
SRC_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Default target
.PHONY: all
all: build

# Build target
.PHONY: build
build: $(BUILD_DIR)/$(BINARY_NAME)

$(BUILD_DIR)/$(BINARY_NAME): $(SRC_FILES)
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)

# Run target
.PHONY: run
run:
	$(GO) run $(GOFLAGS) main.go


# Clean target
.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build    - Build the binary"
	@echo "  run      - Run the application"
	@echo "  clean    - Clean build artifacts"
	@echo "  help     - Show this help message"
