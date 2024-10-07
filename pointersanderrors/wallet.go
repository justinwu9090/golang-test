package Wallet

type Bitcoin int // creating type

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
