APP_NAME := echodb
GOOS     ?= $(shell go env GOOS)
GOARCH   ?= $(shell go env GOARCH)

DOCKER_IMAGE := $(APP_NAME)
DOCKER_TAG   := latest
DOCKER_PORT  := 6380

.PHONY: build run start test docker-build docker-run


build:
	@echo "Building $(APP_NAME) for $(GOOS)/$(GOARCH)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) \
	go build \
	-o $(APP_NAME)

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