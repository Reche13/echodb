APP_NAME := echodb
BUILD_DIR ?= .
VERSION := $(shell cat VERSION)
GOOS     ?= $(shell go env GOOS)
GOARCH   ?= $(shell go env GOARCH)

LDFLAGS := -X 'github.com/reche13/echodb/internal/info.Version=$(VERSION)'

DOCKER_IMAGE := $(APP_NAME)
DOCKER_TAG   := $(VERSION)
DOCKER_PORT  := 6380

.PHONY: build run start test docker-build docker-run


build:
	@echo "Building $(APP_NAME) for $(GOOS)/$(GOARCH)"
	@echo "version $(VERSION)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) \
	go build \
	-ldflags "$(LDFLAGS)" \
	-o $(BUILD_DIR)/$(APP_NAME)

build-prod:
	@echo "Building PROD $(APP_NAME) for $(GOOS)/$(GOARCH)"
	@echo "version $(VERSION)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) \
	CGO_ENABLED=0 \
	go build \
	-trimpath \
	-ldflags "$(LDFLAGS) -s -w" \
	-o $(BUILD_DIR)/$(APP_NAME)

run:
	go run main.go

start: build
	@./$(APP_NAME)

test:
	go test -v -cover ./...

docker-build:
	@echo "Building Docker image $(DOCKER_IMAGE):$(DOCKER_TAG)"
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

docker-run:
	@echo "Running Docker container $(DOCKER_IMAGE) on port $(DOCKER_PORT)"
	docker run --rm -p $(DOCKER_PORT):$(DOCKER_PORT) $(DOCKER_IMAGE):$(DOCKER_TAG)