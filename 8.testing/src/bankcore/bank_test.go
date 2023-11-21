package bank

import (
	"testing"
)

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, Carlifornia",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)

	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, Carlifornia",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	if err := account.Deposit(-10); err == nil {
		t.Error("only positive numbers should be allowed to deposit")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, Carlifornia",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("balance is not being updated after withdraw")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, Carlifornia",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := Statement(account)

	if statement != "1001 - John - 100" {
		t.Error("statement doesn't have the proper format")
	}
}

func TestTransfer(t *testing.T) {
	a := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, Carlifornia",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 100,
	}
	b := Account{
		Customer: Customer{
			Name:    "Taylor Swift",
			Address: "Cornelia St., New York",
			Phone:   "013 555 1989",
		},
		Number: 1989,
	}

	err := a.Transfer(10, &b)
	if err != nil {
		t.Error(err)
	}
	if a.Balance != 90 || b.Balance != 10 {
		t.Error("Unexpected result")
	}
}
