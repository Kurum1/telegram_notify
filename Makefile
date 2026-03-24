.PHONY: build build-windows clean test help

BINARY_NAME=telegram_notify
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.version=${VERSION}"

help:
	@echo "可用的命令:"
	@echo "  make build         - 编译当前平台"
	@echo "  make build-windows - 编译 Windows 可执行文件"
	@echo "  make test          - 运行测试"
	@echo "  make clean         - 清理编译产物"

build:
	go build ${LDFLAGS} -o ${BINARY_NAME}

build-windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY_NAME}.exe

test:
	go test -v ./...

clean:
	rm -f ${BINARY_NAME} ${BINARY_NAME}.exe
