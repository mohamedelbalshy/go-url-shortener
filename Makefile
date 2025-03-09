# Variables
APP_NAME := go-url-shortener
DOCKER_IMAGE := $(APP_NAME):latest
BINARY_DIR := ./bin
SWAGGER_DIR := ./docs
MAIN_FILE := ./main.go  

# Default target
.PHONY: all
all: build

# Build the Go application
.PHONY: build
build:
	@echo "Building the application..."
	@go build -o $(BINARY_DIR)/$(APP_NAME) $(MAIN_FILE)

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	@$(BINARY_DIR)/$(APP_NAME)

# Generate Swagger documentation
.PHONY: swag
swag:
	@echo "Generating Swagger documentation..."
	@swag init -g $(MAIN_FILE) -o $(SWAGGER_DIR)

# Build the Docker image
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .

# Run Docker Compose
.PHONY: docker-compose-up
docker-compose-up:
	@echo "Starting services with Docker Compose..."
	@docker-compose up --build

# Stop Docker Compose
.PHONY: docker-compose-down
docker-compose-down:
	@echo "Stopping services with Docker Compose..."
	@docker-compose down

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@go clean
	@rm -rf $(BINARY_DIR)
	@rm -rf $(SWAGGER_DIR)
