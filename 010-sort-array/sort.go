package main

import (
	"fmt"
)

func mergeSort(arr []int) (result []int, err error) {
	if len(arr) <= 1 {
		return arr, err
	}

	mid := len(arr) / 2
	left, err := mergeSort(arr[:mid])
	right, err := mergeSort(arr[mid:])

	return merge(left, right), nil
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, y := 0, 0

	for i < len(left) && y < len(right) {
		if left[i] < right[y] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[y])
			y++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[y:]...)

	return result

}

func main() {
	arr := []int{1, 5, 2, 3, 2}

	result, err := mergeSort(arr)

	if err != nil {
		fmt.Printf("We have a problem %s\n", err)
	}

	fmt.Println("The sorted array is this \n", result)
}
