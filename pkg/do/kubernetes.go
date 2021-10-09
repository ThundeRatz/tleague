package do

import (
	ctx "context"

	"github.com/digitalocean/godo"
)

func (c *Client) ClusterCreate(name string) (*godo.KubernetesCluster, error) {
	cluster, _, err := c.do.Kubernetes.Create(ctx.Background(), &godo.KubernetesClusterCreateRequest{
		Name:        name + "-cluster",
		RegionSlug:  "nyc1",
		VersionSlug: "1.20.7-do.0",
		NodePools: []*godo.KubernetesNodePoolCreateRequest{{
			Name:  name + "-pool",
			Size:  "s-4vcpu-8gb",
			Count: 10,
		}},
	})

	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (c *Client) ClusterGetCredentials(clusterID string) (*godo.KubernetesClusterCredentials, error) {
	creds, _, err := c.do.Kubernetes.GetCredentials(ctx.Background(), clusterID, nil)

	if err != nil {
		return nil, err
	}

	return creds, nil
}
