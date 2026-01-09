#!/bin/bash
set -e

echo "🐶 Setting up K9s development environment..."

# Install Go tools
echo "📦 Installing Go tools..."
go install golang.org/x/tools/gopls@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install golang.org/x/tools/cmd/goimports@latest

# Download Go dependencies
echo "📥 Downloading Go dependencies..."
go mod download

# Build regular k9s binary to verify setup
echo "🔨 Building k9s (regular)..."
go build -o /tmp/k9s-test main.go && rm /tmp/k9s-test

# Build WASM binary
echo "🌐 Building k9s WASM binary..."
make build-wasm

# Verify web directory
echo "✓ Web directory contents:"
ls -lh web/

echo ""
echo "✅ Development environment ready!"
echo ""
echo "📝 Quick Start Commands:"
echo "  - Build regular k9s:   go build -o k9s main.go"
echo "  - Build WASM:          make build-wasm"
echo "  - Run tests:           go test ./..."
echo "  - Run WASM demo:       cd web && python3 -m http.server 8080"
echo ""
echo "🌐 To test WASM in browser:"
echo "  1. Run: cd web && python3 -m http.server 8080"
echo "  2. Open the forwarded port 8080 in your browser"
echo "  3. Navigate to index.html"
echo ""
echo "📚 Documentation:"
echo "  - WASM_SUPPORT.md - Comprehensive WASM guide"
echo "  - web/README.md   - Web interface documentation"
echo ""
