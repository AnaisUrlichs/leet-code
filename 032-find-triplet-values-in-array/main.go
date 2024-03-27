package main

import (
	"fmt"
	"sort"
)

func findArrayOfThree(arr []int, target int) (result []int) {
	sort.Ints(arr)

	low, high, triple := 0, 0, 0

	for i := 0; i < len(arr)-2; i++ {
		low = i + 1
		high = len(arr) - 1

		for low < high {
			triple = arr[i] + arr[high] + arr[low]

			if triple == target {
				result = append(result, arr[i])
				result = append(result, arr[low])
				result = append(result, arr[high])
				return result
			} else if triple < target {
				low++
			} else {
				high--
			}
		}
	}

	return []int{0, 0, 0}
}

func main() {
	arr := []int{2, 5, 1, 2, 7, 8}
	target := 10

	result := findArrayOfThree(arr, target)

	fmt.Println(result)
}
