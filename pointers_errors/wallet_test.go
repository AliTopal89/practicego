package main

import (
	"testing"
)

func TestWallet(t *testing.T)  {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10
	
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
// in our very secure wallet we don't want to expose our inner state to the rest of the world