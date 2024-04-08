# Имя исполняемого файла
BINARY_NAME=xkcd

# Правило по умолчанию
all: build

# Правило для сборки приложения
build:
	@echo "Building..." && go build -o $(BINARY_NAME) cmd/xkcd/main.go

# Правило для запуска приложения
run: build
	@echo "Running..." && ./$(BINARY_NAME)

# Правило для тестирования
test:
	@echo "Testing..." && go test -v ./...

.PHONY: all build run test

