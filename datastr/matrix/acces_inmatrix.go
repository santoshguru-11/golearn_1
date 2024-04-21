package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	element := matrix[1][2] // Accessing element at row 1, column 2
	fmt.Println("Element at (1, 2):", element)
}
