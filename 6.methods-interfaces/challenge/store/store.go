package store

import (
	"errors"
	"fmt"
)

type Account struct {
	FirstName string
	LastName  string
}

type Employee struct {
	Account
	Credits float64
}

func (a *Account) ChangeName(name string) {
	a.FirstName = name
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s\nCredits: %.2f", e.FirstName, e.LastName, e.Credits)
}

func NewEmployee(firstName, lastName string, credits float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credits}, nil
}

func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0 {
		e.Credits += amount
		return e.Credits, nil
	}
	return 0, errors.New("invalid credit amount")
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0 {
		if amount < e.Credits {
			e.Credits -= amount
			return e.Credits, nil
		}
		return 0, errors.New("you can't remove more credits than the account has")
	}
	return 0, errors.New("you can't remove negative numbers")
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}
