package main

import "fmt"

func main() {
	// Example of for loop with break statement
	fmt.Println("For loop with break statement:")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 3 {
			break // Exit the loop when i reaches 3
		}
	}
}
