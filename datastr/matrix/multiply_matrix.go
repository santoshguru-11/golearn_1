package main

import "fmt"

func main() {
	matrixA := [][]int{
		{1, 2},
		{3, 4},
	}

	matrixB := [][]int{
		{5, 6},
		{7, 8},
	}

	result := multiplyMatrices(matrixA, matrixB)
	fmt.Println("Product of Matrices:")
	for _, row := range result {
		fmt.Println(row)
	}
}

func multiplyMatrices(a, b [][]int) [][]int {
	rowsA, colsA := len(a), len(a[0])
	rowsB, colsB := len(b), len(b[0])
	if colsA != rowsB {
		fmt.Println("Cannot multiply matrices: incompatible dimensions")
		return nil
	}

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
		for j := range result[i] {
			sum := 0
			for k := 0; k < colsA; k++ {
				sum += a[i][k] * b[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}
