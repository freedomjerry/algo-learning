package Wallet

import "testing"

func TestWallet(t *testing.T)  {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	assertNoError := func(t *testing.T, err error) {
		if err != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
			wallet := Wallet{}

			wallet.Deposite(Bitcoin(10))
			assertBalance(t, wallet, Bitcoin(10))
		//	got := wallet.Balance()
		//	want := Bitcoin(10)
		//
		//	if got != want {
		//		t.Errorf("got %s, want %s", got, want)
		//}
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
		//got := wallet.Balance()
		//want := Bitcoin(10)
		//
		//if got != want {
		//	t.Errorf("got %s want %s", got, want)
		//}
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		walllet := Wallet{startingBalance}
		err := walllet.Withdraw(Bitcoin(100))

		assertBalance(t, walllet, Bitcoin(20))
		assertError(t, err, InsufficientFundsError)
	})
}
