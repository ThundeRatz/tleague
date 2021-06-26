package do

import (
	ctx "context"

	"github.com/digitalocean/godo"
)

func (c *Client) ClusterCreate() (*godo.KubernetesCluster, error) {
	cluster, _, err := c.do.Kubernetes.Create(ctx.Background(), &godo.KubernetesClusterCreateRequest{})

	if err != nil {
		return nil, err
	}

	return cluster, nil
}
