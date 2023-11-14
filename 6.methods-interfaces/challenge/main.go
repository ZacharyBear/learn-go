package main

import (
	"challenge/store"
	"fmt"
)

func main() {
	zenkie, _ := store.NewEmployee("Zenkie", "Bear", 1000)
	fmt.Println(zenkie.CheckCredits())
	credits, err := zenkie.AddCredits(250)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance is =", credits)
	}

	_, err = zenkie.RemoveCredits(2500)
	if err != nil {
		fmt.Println("Cant't withdraw or overdraw!", err)
	}

	zenkie.ChangeName("Swift")
	fmt.Println(zenkie)
}
