// Digital Ocean
package do

import (
	ctx "context"

	"github.com/digitalocean/godo"
)

type Client struct {
	do *godo.Client
}

func NewClient(token string) *Client {
	return &Client{
		do: godo.NewFromToken(token),
	}
}

func (c *Client) DropletList() ([]godo.Droplet, error) {
	list := []godo.Droplet{}
	opt := &godo.ListOptions{}

	for {
		droplets, resp, err := c.do.Droplets.List(ctx.Background(), opt)
		if err != nil {
			return nil, err
		}

		list = append(list, droplets...)

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		opt.Page = page + 1
	}

	return list, nil
}

func (c *Client) SnapshotList() ([]godo.Snapshot, error) {
	list := []godo.Snapshot{}
	opt := &godo.ListOptions{}

	for {
		snapshots, resp, err := c.do.Snapshots.List(ctx.Background(), opt)
		if err != nil {
			return nil, err
		}

		list = append(list, snapshots...)

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		opt.Page = page + 1
	}

	return list, nil
}
