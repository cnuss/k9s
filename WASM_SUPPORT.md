# K9s WebAssembly Support

## Overview

K9s now includes experimental WebAssembly (WASM) support, enabling it to run in a web browser without requiring a backend server or real Kubernetes cluster connection. This feature provides a demo mode with mock Kubernetes data.

## Quick Start

### Building

```bash
make build-wasm
```

This command:
1. Compiles `cmd/web/main.go` to `web/k9s.wasm`
2. Copies the Go WASM runtime (`wasm_exec.js`) to the web directory

### Running

```bash
cd web
python3 -m http.server 8080
```

Open http://localhost:8080/index.html in your browser.

## Features

### Current Implementation (v1)

✅ **Working:**
- WebAssembly compilation using Go's js/wasm target
- Browser-based demo with mock cluster data
- JavaScript API for programmatic access
- Mock configuration (demo-cluster, demo-context)
- Console output showing cluster state
- Proper resource scoping (cluster vs namespaced)
- Enhanced error handling with diagnostics

### JavaScript API

The WASM module exposes three API functions accessible from the browser console:

```javascript
// Get demo status and cluster information
k9sDemo()
// Returns: {
//   version: "v0.50.16-wasm-demo",
//   status: "running",
//   mode: "demo",
//   cluster: "demo-cluster",
//   context: "demo-context"
// }

// List available namespaces
k9sGetNamespaces()
// Returns: ["default", "kube-system", "demo-app"]

// List mock pods
k9sGetPods()
// Returns: Array of pod objects with name, namespace, status, ready fields
```

### Mock Data

The demo includes realistic mock Kubernetes data:

- **Context:** demo-context
- **Cluster:** demo-cluster
- **Namespaces:** default, kube-system, demo-app
- **Pods:** 4 total (nginx, coredns, api-server, worker)
  - 3 Running
  - 1 Pending (ImagePullBackOff)
- **Deployments:** 2 (nginx-deployment, demo-api-deployment)
- **Services:** 2 (kubernetes, demo-api-service)
- **Nodes:** 1 (demo-node-1)

## Architecture

### Build System

The implementation uses Go build tags to separate browser-specific code:

```go
//go:build js && wasm
```

Files with this tag are only compiled for WASM builds.

```go
//go:build !js || !wasm
```

Files with this tag are excluded from WASM builds.

### File Structure

```
cmd/web/main.go              # WASM entry point
internal/client/wasm.go      # Mock client for WASM
internal/config/wasm.go      # Mock config for WASM
web/
  ├── index.html             # HTML host page
  ├── main.js                # JavaScript initialization
  ├── k9s.wasm               # Compiled WASM binary (generated)
  ├── wasm_exec.js           # Go WASM runtime (copied)
  └── README.md              # Detailed documentation
```

### Key Components

1. **WASM Entry Point (`cmd/web/main.go`)**
   - Initializes demo mode
   - Registers JavaScript callbacks
   - Displays mock cluster information
   - Maintains Go runtime lifecycle

2. **Mock Client (`internal/client/wasm.go`)**
   - Implements APIClient interface for demo mode
   - Returns mock cluster and resource data
   - Provides proper resource scoping
   - Always reports healthy connectivity

3. **Mock Config (`internal/config/wasm.go`)**
   - Provides fake kubeconfig
   - Implements InitLocs/InitLogLoc for browser
   - Returns demo cluster configuration

4. **Web Interface**
   - Loads and initializes WASM module
   - Handles errors with specific diagnostics
   - Displays status and output
   - Provides clean UI with terminal styling

## Technical Details

### Compilation

```bash
GOOS=js GOARCH=wasm go build -tags=wasm -o web/k9s.wasm cmd/web/main.go
```

The resulting WASM binary:
- Size: ~2.4 MB
- Format: WebAssembly binary
- Runtime: Go 1.21+ WASM runtime

### Browser Requirements

- Modern browser with WebAssembly support:
  - Chrome 57+
  - Firefox 52+
  - Safari 11+
  - Edge 16+
- JavaScript enabled
- Local web server (for MIME types and CORS)

### Resource Scoping

The mock client properly identifies resource scope:

**Cluster-scoped resources:**
- nodes (no)
- namespaces (ns)
- persistentvolumes (pv)
- clusterroles
- clusterrolebindings
- storageclasses (sc)

**Namespaced resources:**
- pods
- deployments
- services
- configmaps
- secrets
- etc.

## Limitations

This is a proof-of-concept demo with several limitations:

1. **No Real Cluster Connection**
   - All data is mocked
   - Cannot connect to actual Kubernetes clusters
   - No authentication or authorization

2. **TUI Not Integrated**
   - Full k9s terminal UI not yet rendered in browser
   - Pending tcell WASM terminal driver integration
   - Currently shows status output only

3. **Read-Only**
   - Cannot modify cluster state
   - No kubectl operations
   - No exec, logs, port-forward capabilities

4. **Static Mock Data**
   - Data doesn't update dynamically
   - No watch events
   - No metrics

## Future Enhancements

### Phase 2: TUI Integration
- Integrate tcell WASM terminal driver
- Render full k9s interface in browser
- Enable keyboard navigation
- Support all view modes

### Phase 3: Complete Mock Client
- Implement full Kubernetes API mock
- Add watch events for resource updates
- Simulate pod lifecycle changes
- Add more resource types

### Phase 4: Browser Storage
- Implement IndexedDB/localStorage adapter
- Persist configuration across sessions
- Save view preferences
- Store command history

### Phase 5: Enhanced Mock Data
- Add more realistic resource relationships
- Simulate events and conditions
- Add resource metrics
- Include custom resources

## Development

### Adding New Mock Resources

To add a new resource type to the mock client:

1. Add data structure in `internal/client/wasm.go`:
```go
func getMockConfigMaps(namespace string) []v1.ConfigMap {
    // Return mock configmap data
}
```

2. Add JavaScript API function in `cmd/web/main.go`:
```go
js.Global().Set("k9sGetConfigMaps", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    return getMockConfigMapData()
}))
```

### Testing WASM Build

```bash
# Build
make build-wasm

# Test in browser
cd web
python3 -m http.server 8080
# Open http://localhost:8080/index.html

# Test JavaScript API
# Open browser console and run:
k9sDemo()
k9sGetNamespaces()
k9sGetPods()
```

### Debugging

Enable verbose WASM debugging:
1. Open browser DevTools (F12)
2. Check Console tab for Go runtime messages
3. Check Network tab for WASM loading issues
4. Use browser's WASM debugging tools

## FAQ

**Q: Why can't I see the full k9s interface?**
A: The terminal UI (TUI) integration is pending. The current implementation is a proof-of-concept showing that k9s can compile to WASM. Full TUI rendering requires tcell's WASM terminal driver.

**Q: Can I connect to a real Kubernetes cluster?**
A: No, this is a demo mode with mock data only. Real cluster connections are not supported in WASM mode due to browser security restrictions.

**Q: Why is the WASM file so large?**
A: The WASM binary includes the entire Go runtime and dependencies. This is normal for Go WASM applications. Future optimizations could reduce size.

**Q: Does this work offline?**
A: Yes, once loaded. The WASM binary and all resources are cached by the browser. You can work offline after the initial load.

**Q: Can I use this in production?**
A: No, this is an experimental demo for evaluation and demonstration purposes only. It's not suitable for production use.

## Contributing

To contribute to WASM support:

1. Use proper build tags (`//go:build js && wasm` or `//go:build !js || !wasm`)
2. Test both regular and WASM builds
3. Update mock data to match real Kubernetes behavior
4. Add JavaScript API functions for new features
5. Document all changes in web/README.md

## References

- [Go WebAssembly Wiki](https://github.com/golang/go/wiki/WebAssembly)
- [tcell Terminal Library](https://github.com/gdamore/tcell)
- [WebAssembly Specification](https://webassembly.org/specs/)

## License

Same as k9s: Apache License 2.0

## Security Summary

The WASM implementation:
- ✅ No sensitive data stored or transmitted
- ✅ No network requests to external services
- ✅ No filesystem access (browser sandbox)
- ✅ No credential handling
- ✅ Enhanced error handling prevents information leakage
- ✅ All mock data is safe for public viewing

This demo mode is safe for public demonstrations and does not expose any sensitive information.
