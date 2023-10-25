package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// # for loop
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of 1...100 is", sum)

	// ## Without pre-handle and suf-handle statement
	var num int64
	rand.Seed(time.Now().Unix())
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Print(num)
	}
	fmt.Println()

	// ## Infinity loop and break statement
	rand.Seed(time.Now().Unix())

	for {
		fmt.Print("Wrinting inside the loop...")
		if num = int64(rand.Int31n(10)); num == 5 {
			fmt.Println("Finished")
			break
		}
		fmt.Println(num)
	}

	sum = 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum)
}
