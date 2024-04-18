package main

import "fmt"

func contains(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}

func main() {
	numbers := []int{5, 9, 3, 7, 2, 10}
	target := 7
	fmt.Println("Does", target, "exist?", contains(numbers, target))
}
