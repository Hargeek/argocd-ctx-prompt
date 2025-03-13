.PHONY: build install clean

# 设置变量
BINARY_NAME=argocd-ctx-prompt
GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin

# 默认目标
all: build

# 构建二进制文件
build:
	go build -o bin/$(BINARY_NAME) .

# 安装到 GOBIN 目录
install: build
	cp $(BINARY_NAME) $(GOBIN)/$(BINARY_NAME)

# 清理构建产物
clean:
	rm -f $(BINARY_NAME)

# 运行测试
test:
	go test -v ./... 