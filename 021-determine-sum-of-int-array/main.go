package main

// Problem 1: https://www.educative.io/blog/apple-coding-interview-questions

import "fmt"

func findNums(arr []int, num int) (result []int) {
	if len(arr) < 2 {
		// = []int{0}
		return result
	}

	for i := 0; i < len(arr)-2; i++ {
		for y := 1; i < len(arr)-1; y++ {
			for z := 2; z < len(arr); z++ {
				currentSum := arr[i] + arr[y] + arr[z]
				if currentSum == num {
					result = []int{arr[i], arr[y], arr[z]}
					return result
				}
			}
		}
	}

	return result
}

func main() {
	arr := []int{1, 4, 3, 2, 6}
	num := 8

	result := findNums(arr, num)
	fmt.Printf("The following numbers equal the sum %v", result)
}
