package main

import "fmt"

func main() {
	// # Map
	studentsAge := map[string]int{
		"Zach": 18,
		"Bob":  29,
	}
	fmt.Println(studentsAge)

	// # Initialize & Appending
	studentsAge = make(map[string]int)
	studentsAge["Zach"] = 18
	studentsAge["Bob"] = 29
	fmt.Println(studentsAge)

	// # Accessing
	age, exist := studentsAge["Milo"]
	if exist {
		fmt.Println("Milo's age is", age)
	} else {
		fmt.Println("Milo's age couldn't be found")
	}

	// # Delete Item
	delete(studentsAge, "Bob")
	fmt.Println(studentsAge)

	// # Traversal Maps
	studentsAge["Milo"] = 32
	for name, age := range studentsAge {
		fmt.Printf("%s\t %d\n", name, age)
	}
	for _, age := range studentsAge {
		fmt.Println("Ages", age)
	}
	for name, _ := range studentsAge {
		fmt.Println("Names", name)
	}
}
