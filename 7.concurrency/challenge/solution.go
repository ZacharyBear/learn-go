package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fibonacci(ch chan<- string, number int) {
	x, y := 1, 1
	for i := 0; i < number; i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- fmt.Sprintf("Fibonacci(%v): %v\n", number, x)
}

func main() {
	start := time.Now()

	size := 15
	ch := make(chan string, size)

	for i := 0; i < size; i++ {
		go fibonacci(ch, i)
	}

	for i := 0; i < size; i++ {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
