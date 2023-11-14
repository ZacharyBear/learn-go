package main

import (
	"fmt"
	"strings"

	"github.com/zenkiebear/learn-go/geometry"
)

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

type square struct {
	size int
}

func (t square) perimeter() int {
	return t.size * 4
}

func (t *triangle) scale(times int) {
	t.size *= times
}

// custom type
type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

// embedding structs
type coloredTriangle struct {
	triangle
	color string
}

// override methods
func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}

func main() {
	t := triangle{3}
	s := square{3}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())

	t.scale(2)
	fmt.Println("Size:", t.size)

	msg := upperstring("Learning Go!")
	fmt.Println(msg)
	fmt.Println(msg.Upper())

	t1 := coloredTriangle{
		triangle{3},
		"blue",
	}
	fmt.Println("Size:", t1.size)
	fmt.Println("Perimeter (colored):", t1.perimeter())
	fmt.Println("Perimeter (normal):", t1.triangle.perimeter())

	// Use another package
	t2 := geometry.Triangle{}
	t2.SetSize(3)
	// t2.doubleSize()
	fmt.Println("Perimeter:", t2.Perimeter())
}
