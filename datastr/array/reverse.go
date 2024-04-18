package main

import (
	"fmt"
)

func reverseArray(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		// Swap elements arr[i] and arr[n-i-1]
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

func main() {
	// Sample array
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println("Original array:", arr)

	// Reverse the array
	reverseArray(arr)

	fmt.Println("Reversed array:", arr)
}
