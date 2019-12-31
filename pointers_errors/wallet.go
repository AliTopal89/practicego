package pointers

import (
	"fmt"
)

type Wallet struct {
	balance int
	//balance variable in our struct to store the state
 }

// remember methods
// func (receiverName RecieverType) MethodName(args)

func(w *Wallet) Deposit(amount int) {
	//*Wallet ~ pointer to a wallet
	fmt.Printf("origin of balance in Deposit method is %v \n", &w.balance)
	w.balance += amount
}

func(w *Wallet) Balance() int {
	return w.balance
}