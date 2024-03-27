package main

import "fmt"

var currentPosition int

func findPeak(arr []int) (int, int) {
	if len(arr) <= 1 {
		return 0, 0
	}

	highest := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > highest {
			highest = arr[i]
			currentPosition = i
		}
	}

	return highest, currentPosition
}

func main() {
	arr := []int{1, 2, 3, 1}

	result, position := findPeak(arr)

	fmt.Printf("Number %v has been the highest at position %v", result, position)
}
