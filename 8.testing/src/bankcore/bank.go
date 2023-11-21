package bank

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

type Statable interface {
	Statement() string
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}

	a.Balance -= amount
	return nil
}

func Statement(state Statable) string {
	return state.Statement()
}

func (a Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

func (transferor *Account) Transfer(amount float64, receiver *Account) error {
	err := transferor.Withdraw(amount)
	if err != nil {
		return err
	}
	err = receiver.Deposit(amount)
	if err != nil {
		err1 := transferor.Deposit(amount)
		if err1 != nil {
			panic(err1)
		}
		return err
	}
	return nil
}
