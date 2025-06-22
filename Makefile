# 项目变量
BINARY_NAME=lanCLIChat
VERSION=1.0.0
BUILD_TIME=$(shell date +%Y-%m-%d\ %H:%M:%S)
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DIR=build
SOURCE_DIR=./cmd/lanCLIChat

# Go 编译标志
LDFLAGS=-ldflags "-X 'main.Version=$(VERSION)' -X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GitCommit=$(GIT_COMMIT)'"

# 默认目标
.PHONY: all build clean run test deps version help

all: build

# 显示版本信息
version:
	@echo "项目: $(BINARY_NAME)"
	@echo "版本: $(VERSION)"
	@echo "构建时间: $(BUILD_TIME)"
	@echo "Git提交: $(GIT_COMMIT)"

# 构建目标
build:
	@echo "正在构建 $(BINARY_NAME) v$(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_DIR)
	@echo "构建完成: $(BUILD_DIR)/$(BINARY_NAME)"

# 构建多平台版本
build-all: build-linux build-windows build-darwin

build-linux:
	@echo "构建 Linux 版本..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(SOURCE_DIR)

build-windows:
	@echo "构建 Windows 版本..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(SOURCE_DIR)

build-darwin:
	@echo "构建 macOS 版本..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(SOURCE_DIR)

# 清理目标
clean:
	@echo "清理构建文件..."
	rm -rf $(BUILD_DIR)
	@echo "清理完成"

# 运行目标
run: build
	@echo "运行 $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# 开发模式运行（直接运行源码）
dev:
	@echo "开发模式运行..."
	go run $(SOURCE_DIR)/main.go

# 测试目标
test:
	@echo "运行测试..."
	go test -v ./...

# 测试覆盖率
test-coverage:
	@echo "运行测试覆盖率..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告生成: coverage.html"

# 安装依赖
deps:
	@echo "安装依赖..."
	go mod download
	go mod tidy

# 更新依赖
update-deps:
	@echo "更新依赖..."
	go get -u ./...
	go mod tidy

# 格式化代码
fmt:
	@echo "格式化代码..."
	go fmt ./...

# 代码检查
lint:
	@echo "代码检查..."
	golangci-lint run

# 安装到系统
install: build
	@echo "安装 $(BINARY_NAME) 到系统..."
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "安装完成"

# 卸载
uninstall:
	@echo "卸载 $(BINARY_NAME)..."
	sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "卸载完成"

# 打包发布
package: build-all
	@echo "打包发布版本..."
	@mkdir -p $(BUILD_DIR)/release
	tar -czf $(BUILD_DIR)/release/$(BINARY_NAME)-$(VERSION)-linux-amd64.tar.gz -C $(BUILD_DIR) $(BINARY_NAME)-linux-amd64
	tar -czf $(BUILD_DIR)/release/$(BINARY_NAME)-$(VERSION)-darwin-amd64.tar.gz -C $(BUILD_DIR) $(BINARY_NAME)-darwin-amd64
	zip -j $(BUILD_DIR)/release/$(BINARY_NAME)-$(VERSION)-windows-amd64.zip $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe
	@echo "打包完成: $(BUILD_DIR)/release/"

# 帮助信息
help:
	@echo "可用的命令:"
	@echo "  build        - 构建项目"
	@echo "  build-all    - 构建所有平台版本"
	@echo "  clean        - 清理构建文件"
	@echo "  run          - 构建并运行项目"
	@echo "  dev          - 开发模式运行"
	@echo "  test         - 运行测试"
	@echo "  test-coverage- 运行测试覆盖率"
	@echo "  deps         - 安装依赖"
	@echo "  update-deps  - 更新依赖"
	@echo "  fmt          - 格式化代码"
	@echo "  lint         - 代码检查"
	@echo "  install      - 安装到系统"
	@echo "  uninstall    - 从系统卸载"
	@echo "  package      - 打包发布版本"
	@echo "  version      - 显示版本信息"
	@echo "  help         - 显示此帮助信息" 