// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

//go:build js && wasm

package config

import (
	"fmt"
	"time"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// getMockKubeConfig returns a mock kubeconfig for WASM demo mode.
func getMockKubeConfig() clientcmdapi.Config {
	return clientcmdapi.Config{
		Kind:       "Config",
		APIVersion: "v1",
		Clusters: map[string]*clientcmdapi.Cluster{
			"demo-cluster": {
				Server:                "https://demo-cluster.k8s.local:6443",
				InsecureSkipTLSVerify: true,
			},
		},
		Contexts: map[string]*clientcmdapi.Context{
			"demo-context": {
				Cluster:   "demo-cluster",
				AuthInfo:  "demo-user",
				Namespace: "default",
			},
		},
		CurrentContext: "demo-context",
		AuthInfos: map[string]*clientcmdapi.AuthInfo{
			"demo-user": {
				Username: "demo",
			},
		},
	}
}

// InitLocs initializes k9s directories for WASM mode (browser storage).
func InitLocs() error {
	// In WASM mode, we use browser storage instead of filesystem
	fmt.Println("InitLocs: Using browser storage for WASM mode")
	return nil
}

// InitLogLoc initializes k9s log location for WASM mode.
func InitLogLoc() error {
	// In WASM mode, logs go to browser console
	fmt.Println("InitLogLoc: Logs will be sent to browser console")
	return nil
}

// GetMockClientConfig returns a mock client config for WASM demo mode.
func GetMockClientConfig() clientcmd.ClientConfig {
	config := getMockKubeConfig()
	return clientcmd.NewDefaultClientConfig(config, &clientcmd.ConfigOverrides{
		Timeout: (30 * time.Second).String(),
	})
}
