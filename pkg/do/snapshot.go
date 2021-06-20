package do

import (
	ctx "context"
	"errors"

	"github.com/digitalocean/godo"
)

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

func (c *Client) GetSnapshotByName(name string) (*godo.Snapshot, error) {
	snaps, err := c.SnapshotList()

	if err != nil {
		return nil, err
	}

	for _, s := range snaps {
		if s.Name == name {
			return &s, nil
		}
	}

	return nil, errors.New("snapshot not found")
}
