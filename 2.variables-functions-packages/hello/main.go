package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"math/cmplx"
	"zenkie.cn/calculator"
)


// Execute: go run variables.go 2 3
func main() {
	// # Variables
	// var firstName, lastName string
	// var age int

	// var (
	// 	firstName, lastName string
	// 	age                 int
	// )

	// var (
	// 	firstName string = "Taylor"
	// 	lastName  string = "Swift"
	// 	age       int    = 33
	// )

	// var (
	// 	firstName = "Taylor"
	// 	lastName  = "Swift"
	// 	age       = 33
	// )

	// var (
	// 	firstName, lastName, age = "Taylor", "Swift", 33
	// )

	firstName, lastName := "Taylor", "Swift"
	// age := 33
	age := 33
	fmt.Println(firstName, lastName, age)

	// # Constants
	const HTTPStatusOK = 200
	// HTTPStatusOK := 200;

	const (
		StatusOK              = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)

	// # Basic Data Types
	// ## Integer
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	fmt.Println(integer8, integer16, integer32, integer64)
	// fmt.Println(integer16 + integer32)

	rune := 'G'
	fmt.Println(rune)

	// Challenge 1
	// var challengeInt int32 = 2147483648;
	var challengeInt int32 = 2147483647
	fmt.Println(challengeInt)

	// Challenge 2
	// var integer uint = -10;
	var integer uint = 10
	fmt.Println(integer)

	// ## Float
	var flt32 float32 = 2147483647
	var flt64 float64 = 9223372036854775807
	fmt.Println(flt32, flt64)

	fmt.Println(math.MaxFloat32, math.MaxFloat64)

	const e = 2.71828
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957 - 34
	fmt.Println(e, Avogadro, Planck)

	// ## Boolean
	var featureFlag bool = true
	fmt.Println(featureFlag)

	// ## String
	var firstName1 string = "Taylor"
	lastName1 := "Swift"
	fmt.Println(firstName1, lastName1)

	fullName := "Taylor Swift \t(alias \"Taylor\\Swiftie\")\n"
	fmt.Println(fullName)

	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)

	// # Type Convert
	fmt.Println(int32(integer16) + integer32)
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)

	// # Reading Bootstrap Arguments
	number1, _ := strconv.Atoi(os.Args[1])
	number2, _ := strconv.Atoi(os.Args[2])
	fmt.Println("Sum: ", number1+number2)

	// # Custom Function
	sum := sum(os.Args[1], os.Args[2])
	fmt.Println("Sum: ", sum)

	// ## Multi-result
	_, multiple := calc(os.Args[1], os.Args[2])
	fmt.Println("Multiple: ", multiple)

	// ## Function Pointer Parameter
	firstName = "Evan"
	updateName(&firstName)
	fmt.Println(firstName)

	// # Packages
	fmt.Println(cmplx.Inf())

	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)
	// fmt.Println("logMessage: ", calculator.logMessage)
	// fmt.Println("Internal sum: ", calculator.internalSum(5))
}

func sum(num1 string, num2 string) int {
	int1, _ := strconv.Atoi(num1)
	int2, _ := strconv.Atoi(num2)
	return int1 + int2
}

func calc(num1 string, num2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(num1)
	int2, _ := strconv.Atoi(num2)
	sum = int1 + int2
	mul = int1 * int2
	return
}

func updateName(name *string) {
	*name = "David"
}
