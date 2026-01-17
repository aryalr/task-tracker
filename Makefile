.PHONY: build run clean

# File path
BINARY_NAME=task-tracker-cli
MAIN_PATH=cmd/task-tracker/main.go

# Build
build:
	@echo "Building..."
	@mkdir -p build
	go build -o build/$(BINARY_NAME) $(MAIN_PATH)

# Run
run:
	go run $(MAIN_PATH)

# Clear build folder
clean:
	@echo "Cleaning..."
	rm -rf build/
