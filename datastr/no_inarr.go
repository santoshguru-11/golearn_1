package main

import "fmt"

func countOccurrences(arr []int, target int) int {
	count := 0
	for _, num := range arr {
		if num == target {
			count++
		}
	}
	return count
}

func main() {
	numbers := []int{5, 9, 3, 7, 2, 10, 3, 3}
	target := 3
	fmt.Println("Occurrences of", target, ":", countOccurrences(numbers, target))
}
