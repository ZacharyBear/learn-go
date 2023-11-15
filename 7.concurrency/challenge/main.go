package main

import (
	"challenge/solutions"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// solutions.Solution()
	solutions.Solution1()

	elapsed := time.Since(start)
	fmt.Printf("Done! It look %v seconds!\n", elapsed.Seconds())
}
