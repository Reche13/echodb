APP_NAME := echodb
GOOS     ?= $(shell go env GOOS)
GOARCH   ?= $(shell go env GOARCH)

BIN := ./bin
MAIN_FILE := main.go

.PHONY: build run start test


build:
	@echo "Building $(APP_NAME) for $(GOOS)/$(GOARCH)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) \
	go build \
	-o $(APP_NAME)

run:
	go run $(MAIN_FILE)

start: build
	@./$(APP_NAME)

test:
	go test -v -cover ./...