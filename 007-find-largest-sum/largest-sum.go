package main

// https://www.educative.io/largest-sum-subarray

import "fmt"

var arr []int

func findLargest(arr []int) (result []int) {
	for _, val := range arr {
		if val > 0 {
			result = append(result, val)
		}
	}
	return
}

func main() {
	arr = []int{1, 2, 5, 3, -1}
	result := findLargest(arr)
	fmt.Println(result)
}
