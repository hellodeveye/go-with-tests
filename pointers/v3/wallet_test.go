package wallet

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, err, want error) {
		if err == nil {
			t.Error("wanted an error but didnt get one")
		}
		if err != want {
			t.Errorf("got %s want %s", err, want)
		}
	}

	assertNoError := func(t *testing.T, err error) {
		if err != nil {
			t.Fatal("got an error but didnt want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)

		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})
}
