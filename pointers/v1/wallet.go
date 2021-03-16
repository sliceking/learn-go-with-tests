package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")


// Bitcoin represents a number of Bitcoins.
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of Bitcoin someone owns.
type Wallet struct {
	balance Bitcoin
}

// Deposit will add some Bitcoin to a wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw will remove some Bitcoin to a wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if (w.balance - amount) < 0 {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance returns the number of Bitcoin a wallet has.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
