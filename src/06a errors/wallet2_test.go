package wallet2

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		//// 1) Code is refactored because of duplication - thus, creating assertBalance
		// got := wallet.Balance()
		// want := Bitcoin(10)
		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		// wallet.Withdraw(10)
		err := wallet.Withdraw(Bitcoin(10))

		//// 1) Code is refactored because of duplication
		// got := wallet.Balance()
		// want := Bitcoin(10)
		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		//// 2) Code is refactored because of duplication
		// if err == nil {
		// 	t.Error("wanted an error but didn't get one")
		// }
		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}

// 1) Created to reduce duplicate code
func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

// 2) Created to reduce duplicate code
func assertError(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
