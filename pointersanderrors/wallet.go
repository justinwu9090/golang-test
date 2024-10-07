package Wallet

import (
	"errors"
	"fmt"
)

// var defines values global to the package
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int // creating type from existing data types

type Wallet struct {
	balance Bitcoin
}

// w Wallet --> this will fail since w is copied. we actually need a pointer
// w *Wallet will work. notice w.balance means it's dereferenced + member access
// happened without needing -> or (*w).balance. also (*w).balance is still valid code
// but struct pointers are automatically dereferenced
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance // explicit dereference is done here
}

// define withdraw method on type wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// define String() method on bitcoin type.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
