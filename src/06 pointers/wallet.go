package wallet

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// Wallet is passed as a copy
func (w Wallet) Deposit(amount Bitcoin) {
	// The \n escape character, prints new line after outputting the memory address. We get the pointer to a thing with the address of symbol; &.
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

// The difference is the receiver type is *Wallet rather than Wallet which you can read as "a pointer to a wallet".
func (w *Wallet) DepositAsPointer(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// func (w Wallet) Balance() int {
func (w *Wallet) BalanceAsPointer() Bitcoin {
	return w.balance
}
