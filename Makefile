# Suppress command echoing in recipes
.SILENT:

# Use one shell for entire recipe
.ONESHELL:

# List of platforms to build for
PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64

# Output directory
OUTPUT_DIR := dist
MAIN_PACKAGE := ./cmd/event_dispatcher/main.go

## all: (default target) executes clean and build
.PHONY: all
all: clean build

## help: print this help message
.PHONY: help
help:
	echo 'Usage:'
	sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## clean: remove previously built binaries
.PHONY: clean
clean:
	echo "\033[0;33m⚙ Removing previous builds ...\033[0m"
	rm -rf $(OUTPUT_DIR)

## build: build for all supported platforms
build: $(PLATFORMS)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p $(OUTPUT_DIR)
	OS=$$(echo $@ | cut -d'/' -f1)
	ARCH=$$(echo $@ | cut -d'/' -f2)
	OUTPUT_NAME="pillarbox-event-dispatcher-$${OS}-$${ARCH}"
	echo "\033[0;33m⚙ Building for OS=$${OS} ARCH=$${ARCH} ...\033[0m"
	GOOS=$${OS} GOARCH=$${ARCH} go build -o "$(OUTPUT_DIR)/$${OUTPUT_NAME}" $(MAIN_PACKAGE)
	if [ $$? -eq 0 ]; then
		echo "\033[0;32m✔ Successfully built $${OUTPUT_NAME}\033[0m"
	else
		echo "\033[0;31m✘ Failed to build for $${OS}/$${ARCH}\033[0m"
		exit 1
	fi
