// Digital Ocean
package do

import (
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
