package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"zenkie.cn/bank"
)

var accounts = map[float64]*CustomAccount{}

func main() {
	accounts[1001] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "John",
				Address: "Los Angeles, Carlifornia",
				Phone:   "(213) 555 0147",
			},
			Number: 1001,
		},
	}
	accounts[1989] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "Taylor Swift",
				Address: "Cornelia St., New York",
				Phone:   "013 555 1989",
			},
			Number: 1989,
		},
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, exists := accounts[number]
		if !exists {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			fmt.Fprint(w, bank.Statement(account))
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprint(w, bank.Statement(account))
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprint(w, bank.Statement(account))
			}
		}
	}
}

func getAccount(number float64) (*CustomAccount, error) {
	account, exists := accounts[number]
	if !exists {
		return nil, fmt.Errorf("account with number %v can’t be found", number)
	}
	return account, nil
}

func transfer(w http.ResponseWriter, r *http.Request) {
	transferorqs := r.URL.Query().Get("transferor")
	receiverqs := r.URL.Query().Get("receiver")
	amountqs := r.URL.Query().Get("amount")

	if transferorqs == "" || receiverqs == "" {
		fmt.Fprint(w, "Transferor or receiver is missing!")
		return
	}

	if transferorNumber, err := strconv.ParseFloat(transferorqs, 64); err != nil {
		fmt.Fprint(w, "Invalid transferor account number")
	} else if receiverNumber, err := strconv.ParseFloat(receiverqs, 64); err != nil {
		fmt.Fprint(w, "Invalid receiver account number")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprint(w, "Invalid amount number")
	} else {
		transferor, err := getAccount(transferorNumber)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		receiver, err := getAccount(receiverNumber)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		err = transferor.Account.Transfer(amount, receiver.Account)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		fmt.Fprint(w, bank.Statement(transferor))
	}
}

type CustomAccount struct {
	*bank.Account
}

func (c *CustomAccount) Statement() string {
	json, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(json)
}
