package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	transposed := transposeMatrix(matrix)
	fmt.Println("Transposed Matrix:")
	for _, row := range transposed {
		fmt.Println(row)
	}
}

func transposeMatrix(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
		for j := range transposed[i] {
			transposed[i][j] = matrix[j][i]
		}
	}
	return transposed
}
