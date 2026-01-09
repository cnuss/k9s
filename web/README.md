# K9s WebAssembly Demo

This directory contains the WebAssembly build of K9s for running in a web browser.

## Overview

This is a proof-of-concept demo that compiles K9s to WebAssembly, allowing it to run entirely in a web browser without requiring a backend server or real Kubernetes cluster.

## Current Status

✓ **Working:**
- WebAssembly compilation successful
- Basic Go runtime working in browser
- Web hosting infrastructure (HTML, JS, WASM loading)
- Mock Kubernetes configuration
- JavaScript/Go interop

⧗ **In Progress:**
- Terminal UI integration (tcell WASM driver)
- Complete mock Kubernetes API client
- Browser storage for configuration persistence

## Building

Build the WASM binary using:

```bash
make build-wasm
```

This will:
1. Compile `cmd/web/main.go` to `web/k9s.wasm`
2. Copy the Go WASM runtime (`wasm_exec.js`) to the web directory

## Running

1. Build the WASM binary:
   ```bash
   make build-wasm
   ```

2. Start a local web server from the `web` directory:
   ```bash
   cd web
   python3 -m http.server 8080
   ```

3. Open http://localhost:8080/index.html in your browser

## Architecture

### Files

- **index.html**: HTML host page with terminal container
- **main.js**: JavaScript initialization and WASM loading
- **k9s.wasm**: Compiled K9s WebAssembly binary (generated)
- **wasm_exec.js**: Go WebAssembly runtime (copied from Go SDK)

### Build Tags

The WASM build uses Go build tags to separate browser-specific code:

- `//go:build js && wasm`: Files only included in WASM builds
- `//go:build !js || !wasm`: Files excluded from WASM builds

### Mock Data

The demo includes mock Kubernetes data:
- Fake context: `demo-context`
- Fake cluster: `demo-cluster`
- Mock namespaces: `default`, `kube-system`, `demo-app`
- Mock resources: pods, deployments, services

## JavaScript API

You can interact with the running WASM instance from the browser console:

```javascript
// Get demo status
k9sDemo()
// Returns: {version: "v0.50.16-wasm-demo", status: "running", mode: "demo", cluster: "demo-cluster", context: "demo-context"}

// List namespaces
k9sGetNamespaces()
// Returns: ["default", "kube-system", "demo-app"]

// List pods
k9sGetPods()
// Returns: Array of pod objects with name, namespace, status, ready fields
```

## Limitations

This is a proof-of-concept demo with the following limitations:

- **No real cluster connection**: All data is mocked
- **TUI not yet integrated**: Terminal rendering pending tcell WASM driver
- **Limited features**: Features requiring cluster interaction are not available
- **Read-only**: Cannot modify cluster state

## Next Steps

1. Integrate tcell WASM terminal driver for full TUI rendering
2. Implement complete mock Kubernetes API client with watch events
3. Add browser storage (IndexedDB/localStorage) for configuration persistence
4. Enable interactive navigation through mock resources
5. Add more comprehensive mock data

## Development

When developing WASM-specific code:

1. Use `//go:build js && wasm` tag for WASM-only files
2. Use `//go:build !js || !wasm` tag to exclude files from WASM
3. Test locally with a web server (WASM requires proper MIME types)
4. Check browser console for Go runtime logs and errors

## Troubleshooting

**WASM module fails to load:**
- Ensure `k9s.wasm` and `wasm_exec.js` are in the `web` directory
- Check that the web server is serving with correct MIME types
- Check browser console for detailed error messages

**Blank page:**
- Check browser console for JavaScript errors
- Verify the HTTP server is running and accessible
- Ensure all files are present (index.html, main.js, k9s.wasm, wasm_exec.js)

**Build fails:**
- Ensure Go 1.21+ is installed
- Check that GOOS=js GOARCH=wasm are supported by your Go version
- Verify all dependencies are available
