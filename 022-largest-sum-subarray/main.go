package main

// Largest Sum Subarray https://www.educative.io/blog/apple-coding-interview-questions

import (
	"errors"
	"fmt"
)

func findLargest(arr []int) ([]int, error) {
	err := errors.New("No intergers identified.")
	var result []int

	if len(arr) <= 1 {
		return result, err
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			result = append(result, arr[i])
		}
	}

	if result == nil {
		return result, err
	}

	return result, nil

}

func main() {
	arr := []int{-2, -5, 0}
	result, err := findLargest(arr)

	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("The largest subarray is %v", result)
	}
}
