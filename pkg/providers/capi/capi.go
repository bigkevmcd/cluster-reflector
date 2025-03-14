package capi

import (
	"context"
	"fmt"

	clustersv1alpha1 "github.com/weaveworks/cluster-reflector-controller/api/v1alpha1"
	"github.com/weaveworks/cluster-reflector-controller/pkg/providers"
	capiclusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CAPIProvider queries a Kubernetes cluster for CAPI Cluster resources.
type CAPIProvider struct {
	Namespace            string
	ManagementClusterRef clustersv1alpha1.Cluster
	Client               client.Reader
}

var _ providers.Provider = (*CAPIProvider)(nil)

// NewCAPIProvider creates and returns a CAPIProvider ready for use.
func NewCAPIProvider(client client.Reader, namespace string, managementClusterRef clustersv1alpha1.Cluster) *CAPIProvider {
	provider := &CAPIProvider{
		Client:               client,
		Namespace:            namespace,
		ManagementClusterRef: managementClusterRef,
	}
	return provider
}

func (p *CAPIProvider) ListClusters(ctx context.Context) ([]*providers.ProviderCluster, error) {
	capiClusters := &capiclusterv1.ClusterList{}
	err := p.Client.List(ctx, capiClusters, &client.ListOptions{Namespace: p.Namespace})
	if err != nil {
		return nil, fmt.Errorf("querying CAPI clusters: %w", err)
	}

	var clusters []*providers.ProviderCluster
	for _, capiCluster := range capiClusters.Items {
		clusters = append(clusters, &providers.ProviderCluster{
			Name:       capiCluster.Name,
			ID:         capiCluster.Name,
			KubeConfig: nil,
			Labels:     capiCluster.Labels,
		})
	}

	return clusters, nil
}

// ProviderCluster has an ID to identify the cluster, but CAPI cluster doesn't have a Cluster ID
// therefore wont't match in the case of CAPI
func (p *CAPIProvider) ClusterID(ctx context.Context, kubeClient client.Reader) (string, error) {
	return p.ManagementClusterRef.String(), nil
}
