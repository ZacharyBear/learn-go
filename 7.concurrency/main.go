package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// testChannel()
	// testBufferedChannel()
	// testChannelDirection()
	testSelect()
}

func testChannel() {
	start := time.Now()

	apis := []string{
		"https://www.apple.com/",
		"https://bing.com",
		"https://jd.com",
		"https://api.somewhereintheinternet.com/",
		"https://picsum.photos/",
		"https://sspai.com/",
		"https://www.mi.com/",
		"https://www.bilibili.com/",
	}

	ch := make(chan string, 10)

	for _, api := range apis {
		go checkAPI(api, ch)
		// fmt.Print(<-ch)
	}
	for range apis {
		fmt.Print(<-ch)
	}

	// fmt.Print(<-ch)

	elapsed := time.Since(start)
	fmt.Printf("Done! It look %v seconds!\n", elapsed)
}

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func testBufferedChannel() {
	size := 2
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	go send(ch, "three")
	go send(ch, "four")
	fmt.Println("All data sent to the channel ...")

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done")
}

// it's a channel to only send data
func send(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

// it's a channel to only read data
func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
	// ch <- "Bye!" // this is not allowed
}

func testChannelDirection() {
	ch := make(chan string, 1)
	send(ch, "Hello World")
	read(ch)
}

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"
}

func testSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go process(ch1)
	go replicate(ch2)

	for i := 0; i < 2; i++ {
		select {
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)
		}
	}
}
