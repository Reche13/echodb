APP_NAME := echodb
BIN := ./bin
MAIN_FILE := ./cmd/main.go

.PHONY: build run start


build:
	@mkdir -p $(BIN)
	go build -o $(BIN)/$(APP_NAME) $(MAIN_FILE)

run:
	go run $(MAIN_FILE)

start: build
	@$(BIN)/$(APP_NAME)