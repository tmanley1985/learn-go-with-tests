package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

// We're implementing this interface:
// https://pkg.go.dev/fmt#Stringer
// Remember, that we don't have to explicitly state that
// Bitcoin implements the interface, we just have to actually
// implement the interface structurally.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in test is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// We could have written this like so: return (*w).balance
	// and that would be an explicit dereference. But in Go they
	// thought this was cumbersome. So they made it to where it's an implicit
	// dereference. In fact, these are called Struct Pointers.
	// Also, I don't actually have to use the pointer receiver here as the copy
	// of wallet would have been fine. However, for consistency I'll keep it.
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
