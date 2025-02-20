#!/bin/bash

# Exit on any error
set -e

CONTRACTS_DIR="../packages/contracts/contracts"
BUILD_DIR="build/contracts"
GO_OUT_DIR="internal/integration/go"

# Check for required dependencies
check_dependency() {
    if ! command -v "$1" &> /dev/null; then
        echo "Error: $1 is not installed"
        case "$1" in
            "solc")
                echo "Please install solc (Solidity compiler) using one of these methods:"
                echo "  - npm install -g solc"
                echo "  - brew install solidity"
                echo "  - snap install solc"
                ;;
            "abigen")
                echo "Please install abigen by running:"
                echo "  go install github.com/ethereum/go-ethereum/cmd/abigen@latest"
                ;;
        esac
        exit 1
    fi
}

check_dependency "solc"
check_dependency "abigen"

# Create the output directories if they don't exist
mkdir -p "$BUILD_DIR"
mkdir -p "$GO_OUT_DIR"

# Check if there are any .sol files
if ! ls "$CONTRACTS_DIR"/*.sol 1> /dev/null 2>&1; then
    echo "Error: No .sol files found in $CONTRACTS_DIR"
    exit 1
fi

# First pass: Compile all contracts to get combined JSON output
echo "Compiling all contracts..."
solc --combined-json abi,bin,bin-runtime \
    --optimize --optimize-runs 200 \
    --allow-paths "$CONTRACTS_DIR" \
    "$CONTRACTS_DIR"/*.sol > "$BUILD_DIR/combined.json"

# Second pass: Generate Go bindings using the combined JSON
echo "Generating Go bindings..."
for SOL_FILE in "$CONTRACTS_DIR"/*.sol; do
    FILENAME=$(basename "$SOL_FILE" .sol)
    # Convert to lowercase for package name
    PACKAGE_NAME=$(echo "$FILENAME" | tr '[:upper:]' '[:lower:]')
    
    # Create package directory
    PACKAGE_DIR="$GO_OUT_DIR/$PACKAGE_NAME"
    mkdir -p "$PACKAGE_DIR"
    
    echo "Processing $FILENAME..."
    
    # Use combined JSON for abigen
    abigen --combined-json "$BUILD_DIR/combined.json" \
           --pkg "$PACKAGE_NAME" \
           --type "$FILENAME" \
           --out "$PACKAGE_DIR/${PACKAGE_NAME}.go"
done

echo "Compilation and binding generation complete."
