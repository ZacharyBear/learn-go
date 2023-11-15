package solutions

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func fibonacci(ch chan<- string, number float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- fmt.Sprintf("Fibonacci(%v): %v\n", number, x)
}

func Solution() {
	ch := make(chan string, 100)

	for i := 1; i < 15; i++ {
		go fibonacci(ch, float64(i))
	}

	for i := 1; i < 15; i++ {
		fmt.Println(<-ch)
	}
}

func parse(chData chan<- int, chQuit chan<- bool, input string) {
	if input == "quit" {
		chQuit <- true
		return
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Erorr: ", err)
		return
	}

	x, y := 1.0, 1.0
	for i := 0; i < num; i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	chData <- int(x)
}
func Solution1() {
	chData := make(chan int)
	chQuit := make(chan bool)
	for {
		fmt.Print("Please enter: ")
		var input string
		fmt.Scanf("%s", &input)

		go parse(chData, chQuit, input)

		select {
		case fibonacci := <-chData:
			fmt.Printf("%d\n\n", fibonacci)
		case <-chQuit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}
