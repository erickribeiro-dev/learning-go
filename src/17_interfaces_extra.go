package main

import (
	"errors"
	"fmt"
)

// To use interfaces:
// 1 - Create/define an interface with method signatures.
// Types need to implement all of these methods

type BankAccount interface {
	GetBalance() int           // return an int, representing cents
	Deposit(amount int)        // takes an int as the amount to deposit
	Withdraw(amount int) error // takes an int as the amount to withdraw. Can return an error when the user tries to withdraw more than the balance
}

// 2 - Create different types that share similarities with each other.
// That is, we'll create different types of Bank Accounts
type WellsFargo struct {
	balance int
}
type BitcoinAccount struct {
	balance int
	fee     int
}

// 3 - Implement all the methods defined in the BankAccount interface for the types you want to
func (w *WellsFargo) GetBalance() int {
	return w.balance
}
func (w *WellsFargo) Deposit(amount int) {
	w.balance += amount
}
func (w *WellsFargo) Withdraw(amount int) error {
	if amount > w.balance {
		return errors.New("insufficient funds")
	}
	w.balance -= amount
	return nil
}

// 4 - Also implement it for the Bitcoin type
func (w *BitcoinAccount) GetBalance() int {
	return w.balance
}
func (w *BitcoinAccount) Deposit(amount int) {
	w.balance += amount
}
func (w *BitcoinAccount) Withdraw(amount int) error {
	if amount > w.balance-w.fee {
		return errors.New("insufficient funds")
	}
	w.balance -= amount + w.fee
	return nil
}

// 5 - Constructor methods
func NewWellsFargo() *WellsFargo {
	return &WellsFargo{
		balance: 0,
	}
}
func NewBitcoinAccount() *BitcoinAccount {
	return &BitcoinAccount{
		balance: 0,
		fee:     200,
	}
}

func interfaces_extra() {
	// Creating new instances and testing the methods
	wf := NewWellsFargo()
	wf.Deposit(500)
	wf.Withdraw(150)
	fmt.Println(wf.GetBalance())

	btc := NewBitcoinAccount()
	btc.Deposit(500)
	btc.Withdraw(150)
	fmt.Println(btc.GetBalance())

	// Elements inside a BankAccount slice are treated as a BankAccount type
	myAccounts := []BankAccount{
		NewWellsFargo(),
		NewBitcoinAccount(),
	}

	// Since all the accounts are in the same slice, you can iterate over them
	for _, account := range myAccounts {
		account.Deposit(1000)
		if err := account.Withdraw(200); err != nil {
			fmt.Println(err)
		}
		fmt.Println(account.GetBalance())
	}
}
