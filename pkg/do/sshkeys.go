package do

import (
	ctx "context"
	_ "embed"
	"errors"

	"github.com/digitalocean/godo"
)

//go:embed keys/tleague.sec
var defaultTleagueKeyPriv string

//go:embed keys/tleague.pub
var defaultTleagueKeyPub string

func (c *Client) KeyList() ([]godo.Key, error) {
	list := []godo.Key{}
	opt := &godo.ListOptions{}

	for {
		keys, resp, err := c.do.Keys.List(ctx.Background(), opt)
		if err != nil {
			return nil, err
		}

		list = append(list, keys...)

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

func (c *Client) KeyGetByName(name string) (*godo.Key, error) {
	keys, err := c.KeyList()

	if err != nil {
		return nil, err
	}

	for _, k := range keys {
		if k.Name == name {
			return &k, nil
		}
	}

	return nil, errors.New("key not found")
}

func (c *Client) KeyCreateDefault(name string) (*godo.Key, error) {
	key, _, err := c.do.Keys.Create(ctx.Background(), &godo.KeyCreateRequest{
		Name:      name,
		PublicKey: defaultTleagueKeyPub,
	})

	return key, err
}
