.PHONY: build run clean install tidy test app app-amd64 install-app

APP_NAME := lidless
BUILD_DIR := ./build
CMD_PATH := ./cmd/lidless
BUNDLE_NAME := Lidless.app

build:
	@mkdir -p $(BUILD_DIR)
	MACOSX_DEPLOYMENT_TARGET=10.13 go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_PATH)

build-amd64:
	@mkdir -p $(BUILD_DIR)
	MACOSX_DEPLOYMENT_TARGET=10.13 CGO_ENABLED=1 GOARCH=amd64 CGO_CFLAGS="-arch x86_64" CGO_LDFLAGS="-arch x86_64" go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_PATH)

run: build
	$(BUILD_DIR)/$(APP_NAME)

clean:
	rm -rf $(BUILD_DIR)
	go clean

app: build
	./scripts/create-app-bundle.sh $(BUILD_DIR)/$(APP_NAME) $(BUILD_DIR)

app-amd64: build-amd64
	./scripts/create-app-bundle.sh $(BUILD_DIR)/$(APP_NAME) $(BUILD_DIR)


install-app: app
	@echo "Installing Lidless.app to /Applications..."
	rm -rf /Applications/$(BUNDLE_NAME)
	cp -R $(BUILD_DIR)/$(BUNDLE_NAME) /Applications/
	@echo "Installed to /Applications/$(BUNDLE_NAME)"

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

