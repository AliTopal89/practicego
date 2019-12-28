package main

type Wallet struct {
	balance int
	//balance variable in our struct to store the state
 }

// remember methods
// func (receiverName RecieverType) MethodName(args)

func(w Wallet) Deposit(amount int) {
	w.balance += amount
}

func(w Wallet) Balance() int {
	return w.balance
}