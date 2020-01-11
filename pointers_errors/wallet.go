package pointers

import (
	"fmt"
)

type Bitcoin int 

// type Myname original type so that our struct is more descriptives

type Wallet struct {
	balance Bitcoin
}
//balance variable in our struct to store the state

 type Stringer interface {
	String() string
}
// implement Stringer interface on Bitcoin

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// remember methods
// func (receiverName RecieverType) MethodName(args)

func(w *Wallet) Deposit(amount Bitcoin) {
	//*Wallet ~ pointer to a wallet
	fmt.Printf("origin of balance in Deposit method is %v \n", &w.balance)
	w.balance += amount
}

func(w *Wallet) Balance() Bitcoin {
	return w.balance
}

func(w *Wallet) Withdraw(amount Bitcoin) Bitcoin{
	fmt.Printf("origin of balance in Withdraw method is %v \n", &w.balance)
	w.balance -= amount

	return w.balance
}
