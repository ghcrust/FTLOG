package bank

import (
	"errors"
)

func Withdrawal(balance, amount int) (int, error) {
	if balance < amount {
		return 0, errors.New("insufficient balance!")
	}
	return balance-amount, nil
}
