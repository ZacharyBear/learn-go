package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func someNumber() int {
	return 11
}

func location(city string) (string, string) {
	var region, continent string
	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func main() {
	// # If
	x := 26
	if x%2 == 0 {
		fmt.Println(x, "is even")
	}

	// ## Composite if expression
	if num := someNumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// fmt.Println(num)

	// # Switch
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)

	switch i {
	case 0:
		fmt.Println("Zero...")
	case 1:
		fmt.Println("One...")
	case 2:
		fmt.Println("Two...")
	default:
		fmt.Println("No match...")
	}
	fmt.Println("ok")

	region, continent := location("Irvine")
	fmt.Printf("John works in %s, %s\n", region, continent)

	// ## Calling function
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's the weekend, time to rest!")
	}

	fmt.Println(time.Now().Weekday().String())
	// ## Calling function in case
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?\d{3}[). \-]*\d{4}[.\-]?\d{4}`) // It's Chinese phone number format

	// contact := "zq@zenkie.cn"
	// contact := "12312344321"
	contact := "http://blog.zenkie.cn/"

	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case phone.MatchString(contact):
		fmt.Println(contact, "is a phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}

	// ## Ignore Condition
	rand.Seed(time.Now().Unix())
	r := rand.Float64()
	switch {
	case r > .1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}

	// ## fallthrough
	switch num := 15; {
	// switch num := 90; {
	// switch num := 120; {
	case num < 30:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200\n", num)
	}
}
