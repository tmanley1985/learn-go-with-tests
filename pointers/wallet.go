package main

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in test is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	// We could have written this like so: return (*w).balance
	// and that would be an explicit dereference. But in Go they
	// thought this was cumbersome. So they made it to where it's an implicit
	// dereference. In fact, these are called Struct Pointers.
	// Also, I don't actually have to use the pointer receiver here as the copy
	// of wallet would have been fine. However, for consistency I'll keep it.
	return w.balance
}
