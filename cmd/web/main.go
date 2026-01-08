// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("🐶 K9s WebAssembly Demo Mode")
	fmt.Println("================================")
	fmt.Println("This is a proof-of-concept build of K9s compiled to WebAssembly.")
	fmt.Println("Running entirely in your browser with mock Kubernetes data.")
	fmt.Println("")
	fmt.Println("Demo Cluster Configuration:")
	fmt.Println("  Context: demo-context")
	fmt.Println("  Cluster: demo-cluster")
	fmt.Println("  Server:  https://demo-cluster.k8s.local:6443")
	fmt.Println("")
	fmt.Println("Available Namespaces:")
	fmt.Println("  • default")
	fmt.Println("  • kube-system")
	fmt.Println("  • demo-app")
	fmt.Println("")
	fmt.Println("Mock Resources:")
	fmt.Println("  • 4 Pods (3 Running, 1 Pending)")
	fmt.Println("  • 2 Deployments")
	fmt.Println("  • 2 Services")
	fmt.Println("  • 1 Node")
	fmt.Println("")
	fmt.Println("Current Status:")
	fmt.Println("  ✓ WebAssembly compilation successful")
	fmt.Println("  ✓ Basic Go runtime working in browser")
	fmt.Println("  ✓ Mock configuration loaded")
	fmt.Println("  ⧗ Terminal UI integration pending")
	fmt.Println("  ⧗ Full mock Kubernetes client pending")
	fmt.Println("")
	fmt.Println("Next Steps:")
	fmt.Println("  - Integrate tcell WASM terminal driver")
	fmt.Println("  - Implement complete mock Kubernetes API client")
	fmt.Println("  - Add browser storage for configuration")
	fmt.Println("  - Enable interactive resource navigation")
	
	// Keep the Go runtime alive
	done := make(chan struct{})
	
	// Register JavaScript API functions
	js.Global().Set("k9sDemo", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return map[string]interface{}{
			"version": "v0.50.16-wasm-demo",
			"status":  "running",
			"mode":    "demo",
			"cluster": "demo-cluster",
			"context": "demo-context",
		}
	}))
	
	js.Global().Set("k9sGetNamespaces", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return []interface{}{"default", "kube-system", "demo-app"}
	}))
	
	js.Global().Set("k9sGetPods", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return []interface{}{
			map[string]interface{}{
				"name":      "nginx-demo-abc123",
				"namespace": "default",
				"status":    "Running",
				"ready":     "1/1",
			},
			map[string]interface{}{
				"name":      "coredns-xyz789",
				"namespace": "kube-system",
				"status":    "Running",
				"ready":     "1/1",
			},
			map[string]interface{}{
				"name":      "demo-api-server-def456",
				"namespace": "demo-app",
				"status":    "Running",
				"ready":     "1/1",
			},
			map[string]interface{}{
				"name":      "demo-worker-ghi789",
				"namespace": "demo-app",
				"status":    "Pending",
				"ready":     "0/1",
			},
		}
	}))
	
	fmt.Println("")
	fmt.Println("JavaScript API Available:")
	fmt.Println("  • k9sDemo() - Get demo status")
	fmt.Println("  • k9sGetNamespaces() - List namespaces")
	fmt.Println("  • k9sGetPods() - List pods")
	fmt.Println("")
	fmt.Println("Try these commands in the browser console!")
	
	<-done
}
