.PHONY: build run clean install tidy test

APP_NAME := lidless
BUILD_DIR := ./build
CMD_PATH := ./cmd/lidless

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_PATH)

run: build
	$(BUILD_DIR)/$(APP_NAME)

clean:
	rm -rf $(BUILD_DIR)
	go clean

full-build-install-user: tidy clean build install-user
full-build-install-system: tidy clean build install-system 

install-user: build
	cp $(BUILD_DIR)/$(APP_NAME) ~/bin/$(APP_NAME)

install-system: build
	cp $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/


tidy:
	go mod tidy

test:
	go test -v ./...
