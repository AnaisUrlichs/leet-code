package main

import "fmt"

func findValue(input []int, n int) (result int) {
	revInput := reverseSlice(input)

	result = revInput[n-1]

	return result
}

func reverseSlice(arr []int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	return arr
}

func main() {
	inputs := []int{288, 224, 275, 390, 4, 383, 330, 60, 193}
	n := 6
	result := findValue(inputs, n)

	fmt.Println(result)
}
