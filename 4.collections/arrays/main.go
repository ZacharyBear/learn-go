package main

import (
	"fmt"
	"strconv"
)

func main() {
	// # Array
	// Define an array
	var a [3]int
	// Access an array
	a[1] = 10
	fmt.Println(a[0], a[1], a[len(a)-1])

	// ## Initial an array
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)

	// Without length
	cities1 := [...]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities1)

	// Ignored length and define the last value
	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))

	// ## Multidimensional array
	var binaryArr [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			binaryArr[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, binaryArr[i])
	}
	fmt.Println("\nAll at once:", binaryArr)

	// # Slice
	months := []string{"January", "February", "March", "April", "May",
		"June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))

	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))
	// Extend
	quarter2Extended := quarter2[:4]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))

	// ## Append Item
	testAppending()

	// ## Remove Item
	testRemove()

	// ## Copy Item
	testCopy()
}

func testAppending() {
	var numbers []int
	var lastCap int
	for i := 0; i < 1024; i++ {
		numbers = append(numbers, i)
		capInfo := strconv.Itoa(lastCap)
		if cap(numbers) != lastCap {
			lastCap = cap(numbers)
			capInfo += "!"
		}
		fmt.Printf("%d\t%s\t%v\n", i, capInfo, len(numbers))
	}
}

func testRemove() {
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2

	if remove < len(letters) {
		fmt.Println("Before", letters, "Remove", letters[remove])
		letters = append(letters[:remove], letters[remove+1:]...)
		fmt.Println("After", letters)
	}
}

func testCopy() {
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[:2]
	// slice2 := letters[1:4]
	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])

	slice1[1] = "Z"
	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)
}
