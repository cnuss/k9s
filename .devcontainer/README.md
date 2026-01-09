# K9s Devcontainer

This devcontainer is configured for developing k9s, including full support for WebAssembly (WASM) compilation and browser-based testing.

## What's Included

### Base Image
- **Go 1.24** (Debian Bookworm)
- Full Go toolchain with WASM support

### Tools
- **Python 3.12** - For running the WASM demo web server
- **Node.js LTS** - For potential JavaScript tooling
- **Docker-in-Docker** - For container builds
- **golangci-lint** - Code linting
- **gopls** - Go language server
- **goimports** - Import formatting

### VS Code Extensions
- Go extension with language server
- Docker extension
- Makefile tools
- YAML support
- Markdown linting

## Quick Start

### 1. Open in Codespaces

Click the "Code" button on GitHub and select "Open with Codespaces" or "Create codespace on branch".

The devcontainer will automatically:
- Install all dependencies
- Download Go modules
- Build both regular and WASM binaries
- Set up the development environment

### 2. Build K9s

**Regular Build:**
```bash
k9s-build
# or
go build -o k9s main.go
```

**WASM Build:**
```bash
k9s-wasm
# or
make build-wasm
```

### 3. Test WASM Demo

```bash
# Start web server
k9s-serve
# or
cd web && python3 -m http.server 8080
```

Then:
1. Click on the "Ports" tab in VS Code
2. Find port 8080 (labeled "K9s WASM Demo")
3. Click the globe icon to open in browser
4. Navigate to `index.html`

### 4. Run Tests

```bash
k9s-test
# or
go test ./...
```

### 5. Lint Code

```bash
k9s-lint
# or
golangci-lint run
```

## Custom Aliases

The devcontainer includes helpful aliases:

- `k9s-build` - Build k9s binary
- `k9s-wasm` - Build WASM binary
- `k9s-test` - Run tests
- `k9s-serve` - Start WASM demo server
- `k9s-lint` - Run linter
- `cdweb` - Go to web directory
- `cdcmd` - Go to cmd directory
- `cdinternal` - Go to internal directory

## Port Forwarding

The following ports are automatically forwarded:

- **8080** - Primary WASM demo server
- **8081** - Alternative port 1
- **8082** - Alternative port 2

Access forwarded ports via the VS Code "Ports" tab.

## Directory Structure

```
k9s/
├── .devcontainer/
│   ├── devcontainer.json    # Devcontainer configuration
│   ├── post-create.sh        # Post-creation setup script
│   ├── bashrc                # Custom bash configuration
│   └── README.md             # This file
├── cmd/
│   └── web/main.go           # WASM entry point
├── internal/
│   ├── client/wasm.go        # Mock Kubernetes client
│   └── config/wasm.go        # Mock configuration
├── web/
│   ├── index.html            # WASM demo page
│   ├── main.js               # JavaScript loader
│   ├── k9s.wasm              # Compiled WASM binary (generated)
│   ├── wasm_exec.js          # Go WASM runtime (generated)
│   └── README.md             # Web interface docs
├── WASM_SUPPORT.md           # WASM development guide
└── Makefile                  # Build targets
```

## WASM Development Workflow

1. **Make code changes** to WASM-specific files:
   - `cmd/web/main.go`
   - `internal/client/wasm.go`
   - `internal/config/wasm.go`

2. **Rebuild WASM binary:**
   ```bash
   make build-wasm
   ```

3. **Test in browser:**
   ```bash
   cd web && python3 -m http.server 8080
   ```

4. **Open browser** to forwarded port and test JavaScript API:
   ```javascript
   k9sDemo()
   k9sGetNamespaces()
   k9sGetPods()
   ```

## Troubleshooting

### Build Fails

If the initial build fails:
```bash
go mod download
go mod tidy
make build-wasm
```

### Port Already in Use

Use alternative ports:
```bash
cd web && python3 -m http.server 8081
# or
cd web && python3 -m http.server 8082
```

### WASM Binary Not Found

Rebuild the WASM binary:
```bash
make build-wasm
ls -lh web/k9s.wasm
```

### Go Tools Not Found

Reinstall Go tools:
```bash
go install golang.org/x/tools/gopls@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## Documentation

- **WASM_SUPPORT.md** - Comprehensive WASM feature guide
- **web/README.md** - Web interface documentation
- **Main README.md** - General k9s documentation

## Environment Variables

The devcontainer sets:
- `GOPATH=/go`
- `CGO_ENABLED=0` (for WASM compatibility)
- Go binaries in PATH

## Performance Tips

1. **Incremental builds**: Go's build cache speeds up rebuilds
2. **Parallel tests**: Use `go test -p 4 ./...`
3. **Skip WASM build**: Only rebuild WASM when changing web-related code

## Additional Resources

- [Go WebAssembly Guide](https://github.com/golang/go/wiki/WebAssembly)
- [VS Code Remote Containers](https://code.visualstudio.com/docs/remote/containers)
- [GitHub Codespaces](https://github.com/features/codespaces)

## Support

For issues specific to this devcontainer, check:
1. The post-create script output
2. VS Code Output panel (View > Output)
3. Terminal for any error messages

For k9s WASM support issues, see WASM_SUPPORT.md.
