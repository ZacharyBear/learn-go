package main

import (
	"encoding/json"
	"fmt"
)

// # Define a struct
type Person struct {
	ID        int
	FirstName string `json:name`
	LastName  string
	Address   string `json:address,omitempty`
}

func main() {
	// Intialize
	// var john Person
	// person := Person{1001, "John", "Cena", "Cena's Icecream St."}
	person := Person{LastName: "Bear", FirstName: "Zenkie"}
	fmt.Println(person)

	// Access
	fmt.Println(person.FirstName)
	personPtr := &person
	personPtr.FirstName = "David"
	fmt.Println(person)

	// Struct Embedding
	type Employee struct {
		Infomation Person
		ManagerID  int
	}
	employee := Employee{}
	employee.Infomation.FirstName = "John"
	// ↑ This will break our code

	type Contractor struct {
		Person
		ManagerID int
	}
	contractor := Contractor{
		Person: Person{
			FirstName: "John",
		},
	}
	contractor.LastName = "Cena"
	fmt.Println(contractor.FirstName)

	// JSON encode and decode
	employees := []Employee{
		Employee{
			Infomation: Person{
				FirstName: "Zenkie", LastName: "Bear",
			},
		},
		Employee{
			Infomation: Person{
				FirstName: "Coco", LastName: "Chanel",
			},
		},
	}

	// encode
	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	// decode
	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v\n", decoded)

	// Challenges
	testFibonacci() // Fibonacci
	testRomanNumberals()
}

func fibonacci(n int) []int {
	if n < 2 {
		return nil
	}

	nums := make([]int, n)
	nums[0], nums[1] = 1, 1

	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums
}

func testFibonacci() {
	var n int
	fmt.Print("Please an number:")
	fmt.Scanf("%d", &n)
	fmt.Println(fibonacci(n))
}

func romanNumberals(numberal string) int {
	romanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'I': 1,
	}

	result, prev := 0, 0
	for _, char := range numberal {
		curr, exist := romanMap[char]
		if !exist {
			panic(fmt.Sprintf("Wrong character: %c\n", rune(curr)))
		}
		if prev < curr {
			result += curr - prev*2
		} else {
			result += curr
		}
		prev = curr
	}
	return result
}
func testRomanNumberals() {
	var roman string
	fmt.Print("Please enter an Roman Numberal: ")
	fmt.Scanf("%s", &roman)
	num := romanNumberals(roman)
	fmt.Printf("%s = %d \n", roman, num)
}
