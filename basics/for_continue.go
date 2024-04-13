package main

import "fmt"

func main() {
	// Example of for loop with continue statement
	fmt.Println("For loop with continue statement:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skip the rest of the loop body when i equals 2
		}
		fmt.Println(i)
	}
}
