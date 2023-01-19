//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	"k8s.io/client-go/rest"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	edgev1alpha1 "github.com/kcp-dev/edge-mc/pkg/client/clientset/versioned/typed/edge/v1alpha1"
	"github.com/kcp-dev/logicalcluster/v2"
)

type EdgeV1alpha1ClusterInterface interface {
	EdgeV1alpha1ClusterScoper
	EdgePlacementsClusterGetter
	SinglePlacementSlicesClusterGetter
}

type EdgeV1alpha1ClusterScoper interface {
	Cluster(logicalcluster.Name) edgev1alpha1.EdgeV1alpha1Interface
}

type EdgeV1alpha1ClusterClient struct {
	clientCache kcpclient.Cache[*edgev1alpha1.EdgeV1alpha1Client]
}

func (c *EdgeV1alpha1ClusterClient) Cluster(name logicalcluster.Name) edgev1alpha1.EdgeV1alpha1Interface {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(name)
}

func (c *EdgeV1alpha1ClusterClient) EdgePlacements() EdgePlacementClusterInterface {
	return &edgePlacementsClusterInterface{clientCache: c.clientCache}
}

func (c *EdgeV1alpha1ClusterClient) SinglePlacementSlices() SinglePlacementSliceClusterInterface {
	return &singlePlacementSlicesClusterInterface{clientCache: c.clientCache}
}

// NewForConfig creates a new EdgeV1alpha1ClusterClient for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*EdgeV1alpha1ClusterClient, error) {
	client, err := rest.HTTPClientFor(c)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(c, client)
}

// NewForConfigAndClient creates a new EdgeV1alpha1ClusterClient for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*EdgeV1alpha1ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*edgev1alpha1.EdgeV1alpha1Client]{
		NewForConfigAndClient: edgev1alpha1.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.New("root")); err != nil {
		return nil, err
	}
	return &EdgeV1alpha1ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new EdgeV1alpha1ClusterClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *EdgeV1alpha1ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
