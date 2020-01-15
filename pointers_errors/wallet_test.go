package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T)  {

	assertBalance := func (t *testing.T, wallet Wallet, want Bitcoin)  {
	   t.Helper()
	   got := wallet.Balance()
	   
	   if got != want {
		t.Errorf("got %s want %s", got, want)
	   }
	}

	assertNoError := func (t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	assertError := func (t *testing.T, got error, want error)  {
		t.Helper()
		if got == nil {
			// Errors can be nil because the return type of Withdraw 
			// will be "error", which is an interface.
			t.Fatal("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		
	}

	t.Run("Deposit", func(t *testing.T){
	    wallet := Wallet{}
	    wallet.Deposit(Bitcoin(10))

		fmt.Printf("origin of balance in test is %v \n", &wallet.balance)
														//&myVal - origin of that bit of memory
		assertBalance(t, wallet, Bitcoin(10))

	})

	t. Run("Withdraw", func(t *testing.T){
		 wallet := Wallet{balance: Bitcoin(20)}

		 err := wallet.Withdraw(Bitcoin(10))

		 fmt.Printf("origin of withdrawal balance in test is %v \n", &wallet.balance)

		 assertBalance(t, wallet, Bitcoin(10))
		 assertNoError(t, err)
	})

	t.Run("Withdraw insuficient funds", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	} )
}
// in our very secure wallet we don't want to expose our inner state to the rest of the world
// by doing "Bitcoin(999)" its making a new type to declare methods on.
// update our test format strings so they will use String() instead.
