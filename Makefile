.PHONY: build run clean install tidy test app app-amd64 install-app release

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

release:
	@CURRENT_TAG=$$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"); \
	echo "Current version: $$CURRENT_TAG"; \
	read -p "Bump type (major/minor/patch): " BUMP_TYPE; \
	VERSION=$${CURRENT_TAG#v}; \
	VERSION=$${VERSION%%-*}; \
	MAJOR=$$(echo $$VERSION | cut -d. -f1); \
	MINOR=$$(echo $$VERSION | cut -d. -f2); \
	PATCH=$$(echo $$VERSION | cut -d. -f3); \
	case $$BUMP_TYPE in \
		major) MAJOR=$$((MAJOR + 1)); MINOR=0; PATCH=0 ;; \
		minor) MINOR=$$((MINOR + 1)); PATCH=0 ;; \
		patch) PATCH=$$((PATCH + 1)) ;; \
		*) echo "Invalid bump type. Use major, minor, or patch."; exit 1 ;; \
	esac; \
	NEW_VERSION="v$$MAJOR.$$MINOR.$$PATCH"; \
	echo "New version: $$NEW_VERSION"; \
	read -p "Create and push tag $$NEW_VERSION? (y/n): " CONFIRM; \
	if [ "$$CONFIRM" = "y" ]; then \
		git tag -a $$NEW_VERSION -m "Release $$NEW_VERSION"; \
		git push origin $$NEW_VERSION; \
		echo "Tag $$NEW_VERSION pushed. GitHub Actions will create the release."; \
	else \
		echo "Aborted."; \
	fi
