#!/bin/bash
# Creates a macOS .app bundle for Lidless
# Usage: ./scripts/create-app-bundle.sh [binary_path] [output_dir]

set -e

BINARY_PATH="${1:-./build/lidless}"
OUTPUT_DIR="${2:-./build}"
APP_NAME="Lidless"
BUNDLE_NAME="${APP_NAME}.app"
BUNDLE_PATH="${OUTPUT_DIR}/${BUNDLE_NAME}"

# Validate binary exists
if [ ! -f "$BINARY_PATH" ]; then
    echo "Error: Binary not found at $BINARY_PATH"
    echo "Run 'make build' first."
    exit 1
fi

echo "Creating ${BUNDLE_NAME}..."

# Create bundle directory structure
rm -rf "$BUNDLE_PATH"
mkdir -p "$BUNDLE_PATH/Contents/MacOS"
mkdir -p "$BUNDLE_PATH/Contents/Resources"

# Copy binary
cp "$BINARY_PATH" "$BUNDLE_PATH/Contents/MacOS/lidless"
chmod +x "$BUNDLE_PATH/Contents/MacOS/lidless"

# Copy Info.plist
cp "./resources/Info.plist" "$BUNDLE_PATH/Contents/"

# Copy icon if it exists
if [ -f "./resources/AppIcon.icns" ]; then
    cp "./resources/AppIcon.icns" "$BUNDLE_PATH/Contents/Resources/"
fi

echo "Created ${BUNDLE_PATH}"
