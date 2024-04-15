package main

import (
	"fmt"
)

func main() {
	// Sample array
	numbers := []int{3, 5, 7, 9, 11}

	// Indices of the numbers to add
	index1 := 1
	index2 := 3

	// Check if the indices are within the bounds of the array
	if index1 >= 0 && index1 < len(numbers) && index2 >= 0 && index2 < len(numbers) {
		// Add the numbers at the specified indices
		result := numbers[index1] + numbers[index2]
		fmt.Printf("The sum of %d and %d is %d\n", numbers[index1], numbers[index2], result)
	} else {
		fmt.Println("Invalid indices")
	}
}
