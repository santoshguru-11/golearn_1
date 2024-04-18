package main

import "fmt"

func findMin(arr []int) int {
	min := arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {
	numbers := []int{5, 9, 3, 7, 2, 10}
	min := findMin(numbers)
	fmt.Println("Minimum element:", min)
}
