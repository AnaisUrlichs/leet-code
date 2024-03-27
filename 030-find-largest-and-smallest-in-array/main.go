package main

import (
	"fmt"
)

func findLargestSmallest(arr []int) (largest int, smallest int) {
	left := 1
	right := len(arr) - 1
	smallest = arr[left]

	for left < right {
		if arr[left] < arr[right] && arr[left] <= smallest {
			smallest = arr[left]

		} else if arr[right] <= smallest {
			smallest = arr[right]
		}

		if arr[left] > arr[right] && arr[left] >= largest {
			largest = arr[left]

		} else if arr[right] >= largest {
			largest = arr[right]
		}

		left++
		right--
	}

	return largest, smallest
}

func main() {
	arr := []int{2, 6, 4, 3, 2, 6, 5, 4, 8, 1}
	largest, smallest := findLargestSmallest(arr)

	fmt.Printf("The largest is %v and the smallest is %v", largest, smallest)
}
