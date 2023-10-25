package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// # Use defer
	testDefer()

	// Example
	writeFile()

	// # panic()
	// highlow(2, 0)

	// # recover()
	testRecover()

	// # Callenges
	fizzBuzz()
	testFindPrimes()
	testNumber()
	fmt.Println("Program finished successfully!")
}

func testDefer() {
	for i := 1; i < 5; i++ {
		defer fmt.Println("defered", -i)
		fmt.Println("regular", i)
	}
}
func writeFile() {
	newFile, error := os.Create("learnGo.txt")
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}
	// File will be closed after function execute finished
	defer newFile.Close()

	if _, error = io.WriteString(newFile, "Learning Go!"); error != nil {
		fmt.Println("Error: Could not write to file.")
		return
	}
	newFile.Sync()
}

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}

	defer fmt.Println("Deferred: highlow(", high, ", ", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}
func handlePanics() {
	handler := recover()
	if handler != nil {
		fmt.Println("main(): recover:", handler)
	}
}
func testRecover() {
	defer handlePanics()
	highlow(2, 0)
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func findPrimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number != 1 {
		return true
	} else {
		return false
	}
}
func testFindPrimes() {
	fmt.Println("Prime numbers less than 20:")

	for number := 1; number <= 20; number++ {
		if findPrimes(number) {
			fmt.Printf("%v ", number)
		}
	}
}

func testNumber() {
	val := 0
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		switch {
		case val < 0:
			panic("Warning: negative is not allowed")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		case val > 0:
			fmt.Println("You entered:", val)
		}
	}
}
