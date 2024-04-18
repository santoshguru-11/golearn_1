package main

import "fmt"

func findMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	numbers := []int{5, 9, 3, 7, 2, 10}
	max := findMax(numbers)
	fmt.Println("Maximum element:", max)
}
