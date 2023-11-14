package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	size float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printShape(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

// Extend Spring interface
type Springer interface {
	String() string
}
type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

func main() {
	var s Shape = Square{3}
	printShape(s)

	c := Circle{6}
	printShape(c)

	zk := Person{"Zenkie Bear", "China"}
	kt := Person{"Kevin Ting", "China"}
	fmt.Printf("%s\n%s\n", zk, kt)

	// Extending Writer interface
	res, err := http.Get("https://api.github.com/users/zenkiebear/repos?page=1&per_page=3")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// io.Copy(os.Stdout, resp.Body)
	writer := customWriter{}
	io.Copy(writer, res.Body)

	testServerAPI()
}

type customWriter struct{}
type GithubResponse []struct {
	FullName string `json:"full_name"`
}

func (w customWriter) Write(p []byte) (n int, err error) {
	var res GithubResponse
	err = json.Unmarshal(p, &res)
	if err != nil {
		fmt.Println(err)
	}
	for _, r := range res {
		fmt.Println(r)
	}
	return len(p), nil
}

// Server API
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func testServerAPI() {
	db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	log.Fatal(http.ListenAndServe("localhost:8000", db))

}
