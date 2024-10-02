package Wallet

type Wallet struct {
	balance int
}

// this will fail since w is copied. we actually need a pointer
func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
