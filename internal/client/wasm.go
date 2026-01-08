// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

//go:build js && wasm

package client

import (
	"errors"
	"log/slog"

	"k8s.io/apimachinery/pkg/util/cache"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/tools/clientcmd/api"
)

// InitMockConnection creates a mock connection for WASM/browser demo mode.
func InitMockConnection(config *Config, log *slog.Logger) (*APIClient, error) {
	return &APIClient{
		config: config,
		cache:  cache.NewLRUExpireCache(cacheSize),
		connOK: true,
		log:    log,
	}, nil
}

// CheckConnectivity checks if cluster is reachable (always true in demo mode).
func (a *APIClient) CheckConnectivity() bool {
	return true
}

// ServerVersion returns a mock server version for demo mode.
func (a *APIClient) ServerVersion() (*version.Info, error) {
	return &version.Info{
		Major:      "1",
		Minor:      "28",
		GitVersion: "v1.28.0-demo",
		Platform:   "wasm/js",
	}, nil
}

// ValidNamespaceNames returns mock namespace names.
func (a *APIClient) ValidNamespaceNames() ([]string, error) {
	return []string{"default", "kube-system", "demo-app"}, nil
}

// ActiveCluster returns mock cluster name.
func (a *APIClient) ActiveCluster() string {
	return "demo-cluster"
}

// ActiveNamespace returns the active namespace.
func (a *APIClient) ActiveNamespace() string {
	ns, err := a.CurrentNamespaceName()
	if err != nil {
		return "default"
	}
	return ns
}

// CurrentNamespaceName returns current namespace name.
func (a *APIClient) CurrentNamespaceName() (string, error) {
	return "default", nil
}

// Config returns the client config.
func (a *APIClient) Config() *Config {
	return a.config
}

// ConnectionOK returns connection status.
func (a *APIClient) ConnectionOK() bool {
	return a.connOK
}

// IsNamespaced checks if a resource is namespaced (mock implementation).
func (a *APIClient) IsNamespaced(gvr string) bool {
	// Most resources are namespaced
	return true
}

// CanI checks if user can perform an action (always true in demo mode).
func (a *APIClient) CanI(ns, gvr string, verbs []string) (bool, error) {
	return true, nil
}

// GetContext returns a mock context.
func (c *Config) GetContext(name string) (*api.Context, error) {
	if name == "demo-context" {
		return &api.Context{
			Cluster:   "demo-cluster",
			AuthInfo:  "demo-user",
			Namespace: "default",
		}, nil
	}
	return nil, errors.New("context not found")
}

// CurrentContextName returns the demo context name.
func (c *Config) CurrentContextName() (string, error) {
	return "demo-context", nil
}

// CurrentClusterName returns the demo cluster name.
func (c *Config) CurrentClusterName() (string, error) {
	return "demo-cluster", nil
}

// CurrentNamespaceName returns the demo namespace name.
func (c *Config) CurrentNamespaceName() (string, error) {
	return "default", nil
}

// ClusterNameFromContext returns cluster name from context.
func (c *Config) ClusterNameFromContext(ctx string) (string, error) {
	return "demo-cluster", nil
}

// Contexts returns list of available contexts.
func (c *Config) Contexts() (map[string]*api.Context, error) {
	return map[string]*api.Context{
		"demo-context": {
			Cluster:   "demo-cluster",
			AuthInfo:  "demo-user",
			Namespace: "default",
		},
	}, nil
}
