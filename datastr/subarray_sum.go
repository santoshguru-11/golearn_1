package main

import (
	"fmt"
)

func findEqualSumSubarrays(arr []int) [][]int {
	subarrays := [][]int{}
	sumMap := make(map[int][]int)

	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if sum == 0 {
			subarrays = append(subarrays, []int{0, i})
		}

		if _, found := sumMap[sum]; found {
			for _, index := range sumMap[sum] {
				subarrays = append(subarrays, []int{index + 1, i})
			}
		}

		sumMap[sum] = append(sumMap[sum], i)
	}

	return subarrays
}

func main() {
	arr := []int{4, 6, 3, -4, -5, 1}
	equalSumSubarrays := findEqualSumSubarrays(arr)
	fmt.Println("Equal sum subarrays:", equalSumSubarrays)
}
