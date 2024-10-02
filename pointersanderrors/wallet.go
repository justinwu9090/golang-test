package Wallet

type Bitcoin int // creating type

type Wallet struct {
	balance Bitcoin
}

// this will fail since w is copied. we actually need a pointer
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance // explicit dereference is done here
}
