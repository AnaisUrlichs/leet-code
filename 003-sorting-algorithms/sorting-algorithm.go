package main

// based on https://www.educative.io/answers/merge-sort-in-go

import (
	"fmt"
)

var nums []int

func mergeSort(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	nums = []int{1, 5, 3, 7, 2, 1, 4}
	fmt.Println("Original array: ", nums)

	sorted := mergeSort(nums)
	fmt.Println("Sorted Array: ", sorted)
}
