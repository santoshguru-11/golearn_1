package main

import "fmt"

func findDuplicates(arr []int) []int {
	// Create a map to store the count of each element
	counts := make(map[int]int)
	var duplicates []int

	// Iterate over the array and update the count for each element
	for _, num := range arr {
		counts[num]++
	}

	// Iterate over the map and find elements with count > 1
	for num, count := range counts {
		if count > 1 {
			duplicates = append(duplicates, num)
		}
	}

	return duplicates
}

func main() {
	// Sample array
	numbers := []int{1, 2, 3, 4, 2, 5, 6, 3, 7, 8, 5}

	// Find duplicates
	duplicates := findDuplicates(numbers)

	// Print duplicates
	fmt.Println("Duplicates in the array:", duplicates)
}
