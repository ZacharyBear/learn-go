package main

import (
	"errors"
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employee, err := getInfomation(1000)
	// if err != nil {
	if err.Is(ErrNotFound) { // You can check the error's type
		// Something is wrong. Do something
		fmt.Println(err)
	} else {
		fmt.Println(employee)
	}
}

func getInfomation(id int) (*Employee, error) {
	// # simple strategy
	// employee, err := apiCallEmployee(id)
	// if err != nil {
	// 	// return nil, err // Simply return the error to the caller
	// 	return nil, fmt.Errorf("Got an error when getting the employee infomation: %v", err)
	// }
	// return employee, err

	// # retry strategy
	for i := 0; i < 3; i++ {
		employee, err := apiCallEmployee(id)
		if err == nil {
			return employee, err
		}

		fmt.Println("Something is wrong, we're retrying...")
		time.Sleep(time.Second * 2) // sleep 2 second
	}

	return nil, fmt.Errorf("Server has failed to respond to get the employee infomation, please retry later :|")
}

// Reusable Errors
var ErrNotFound = errors.New("Employee not found")

func apiCallEmployee(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	employee := Employee{FirstName: "Zenkie", LastName: "Bear"}
	return &employee, nil
}
