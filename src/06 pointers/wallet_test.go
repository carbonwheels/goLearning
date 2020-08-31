package wallet

import (
	"fmt"
	"testing"
)

// In Go, when you call a function or a method the arguments are copied.
func TestWalletAsCopy(t *testing.T) {

	wallet := Wallet{}

	// Originally was passing as wallet.Deposit(10), however need to convert to proper datatype
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	// It will return 0 because it's a copy that is passed
	want := Bitcoin(0)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// When calling function with receiver type is *Wallet rather than Wallet, it will pass "a pointer to a wallet".
func TestWalletAsPointer(t *testing.T) {

	walletPointer := Wallet{}

	// Originally was passing as wallet.DepositAsPointer(10), however need to convert to proper datatype
	walletPointer.DepositAsPointer(Bitcoin(10))

	got := walletPointer.BalanceAsPointer()

	fmt.Printf("address of balance in test is %v \n", &walletPointer.balance)

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
