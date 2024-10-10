#!/bin/sh

# ANSI color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No color

# List of platforms to build for
PLATFORMS="linux/amd64 linux/arm64 darwin/amd64 darwin/arm64"

# Output directory
OUTPUT_DIR="dist"
mkdir -p "$OUTPUT_DIR"

# Main package
MAIN_PACKAGE="./cmd/event_dispatcher/main.go"

# Build for each platform
for PLATFORM in $PLATFORMS; do

  GOOS=$(echo "$PLATFORM" | cut -d'/' -f1)
  GOARCH=$(echo "$PLATFORM" | cut -d'/' -f2)

  OUTPUT_NAME="pillarbox-event-dispatcher_${GOOS}_${GOARCH}"

  # Set environment variables and build
  env GOOS="$GOOS" GOARCH="$GOARCH" \
    go build \
      -o "${OUTPUT_DIR}/${OUTPUT_NAME}" "$MAIN_PACKAGE"

  if [ $? -ne 0 ]; then
    echo "${RED}✘ Failed to build for ${GOOS}/${GOARCH}${NC}"
    exit 1
  else
    echo "${GREEN}✔ Successfully built ${OUTPUT_NAME}${NC}"
  fi
done

echo "${GREEN}✔ All builds completed successfully!${NC}"
