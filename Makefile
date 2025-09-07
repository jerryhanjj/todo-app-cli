# Todo App CLI Makefile

# 变量定义
BINARY_NAME=todo
MAIN_PATH=cmd/main.go
BUILD_DIR=build

# 默认目标
.PHONY: all
all: build

# 构建可执行文件
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: ./$(BINARY_NAME)"

# 构建到build目录
.PHONY: build-dir
build-dir:
	@echo "Building $(BINARY_NAME) to $(BUILD_DIR)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# 运行应用（开发模式）
.PHONY: run
run:
	@go run $(MAIN_PATH) $(ARGS)

# 运行测试
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...

# 清理构建文件
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete"

# 安装到系统PATH
.PHONY: install
install: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	@sudo mv $(BINARY_NAME) /usr/local/bin/
	@echo "Installation complete"

# 格式化代码
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# 检查代码
.PHONY: vet
vet:
	@echo "Vetting code..."
	@go vet ./...

# 显示帮助
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build      - Build the application"
	@echo "  build-dir  - Build to build/ directory"
	@echo "  run        - Run in development mode"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build files"
	@echo "  install    - Install to system PATH"
	@echo "  fmt        - Format code"
	@echo "  vet        - Vet code"
	@echo "  help       - Show this help"
	@echo ""
	@echo "Examples:"
	@echo "  make build"
	@echo "  make run ARGS='list'"
	@echo "  make run ARGS='add \"New Task\"'"
