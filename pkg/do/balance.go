package do

import (
	ctx "context"
	"strconv"
)

func (c *Client) GetCredits() (float64, error) {
	bal, _, err := c.do.Balance.Get(ctx.Background())

	if err != nil {
		return 0, err
	}

	balance, err := strconv.ParseFloat(bal.MonthToDateBalance, 64)

	if err != nil {
		return 0, err
	}

	if balance >= 0 {
		return 0, nil
	}

	return -balance, nil
}
