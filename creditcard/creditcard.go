package creditcard

import (
	"errors"
)

type card struct {
	number string
}

func NewCreditCard(n string) (card, error) {
	if len(n) > 0 {
		return card{n}, nil
	}
	return card{}, errors.New("invalid cc number!")
}

func (c *card) Number() string {
	return c.number
}
