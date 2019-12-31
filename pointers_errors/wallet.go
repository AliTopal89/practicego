package main

import (
	"fmt"
)

type Wallet struct {
	balance int
	//balance variable in our struct to store the state
 }

// remember methods
// func (receiverName RecieverType) MethodName(args)

func(w Wallet) Deposit(amount int) {
	fmt.Printf("origin of balance in Deposit method is %v \n", &w.balance)
	w.balance += amount
	fmt.Printf("origin of balance_two in Deposit method is %v \n", &w.balance)
}

func(w Wallet) Balance() int {
	return w.balance
}