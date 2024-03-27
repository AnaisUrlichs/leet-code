package main

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/

import "fmt"

func findPositions(arr []int, toFind int) []int {
	var result []int

	for i, num := range arr {
		if num == toFind {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	arr := []int{2, 5, 4, 3, 4, 2}
	toFind := 4

	result := findPositions(arr, toFind)

	fmt.Printf("Number %v has been found at the following positions \n%v", toFind, result)
}
