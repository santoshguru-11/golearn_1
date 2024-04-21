package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
