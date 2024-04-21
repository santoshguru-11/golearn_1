package main

import "fmt"

func main() {
	matrixA := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	matrixB := [][]int{
		{7, 8, 9},
		{10, 11, 12},
	}

	result := addMatrices(matrixA, matrixB)
	fmt.Println("Sum of Matrices:")
	for _, row := range result {
		fmt.Println(row)
	}
}

func addMatrices(a, b [][]int) [][]int {
	result := make([][]int, len(a))
	for i := range a {
		result[i] = make([]int, len(a[i]))
		for j := range a[i] {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}
