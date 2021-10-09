package do

import (
	ctx "context"
	"errors"
	"strconv"

	"github.com/digitalocean/godo"
)

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

func (c *Client) DropletCreateC32(name, snapshotID string, sshKeyID int) (*godo.Droplet, error) {
	sID, err := strconv.Atoi(snapshotID)

	if err != nil {
		return nil, errors.New("invalid SSH key ID")
	}

	droplet, _, err := c.do.Droplets.Create(ctx.Background(), &godo.DropletCreateRequest{
		Name:   name,
		Region: "nyc1",
		Size:   "c-32",
		Image: godo.DropletCreateImage{
			ID: sID,
		},
		SSHKeys: []godo.DropletCreateSSHKey{{ID: sshKeyID}},
	})

	if err != nil {
		return nil, err
	}

	return droplet, nil
}

func (c *Client) DropletDestroy(dropletID int) error {
	_, err := c.do.Droplets.Delete(ctx.Background(), dropletID)

	return err
}

func DropletGetPublicIP(d godo.Droplet) string {
	for _, net := range d.Networks.V4 {
		if net.Type == "public" {
			return net.IPAddress
		}
	}

	return ""
}
