package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	newarr := make([]int, 0, len(left)+len(right))

	i, y := 0, 0

	for i < len(left) && y < len(right) {
		if left[i] < right[y] {
			newarr = append(newarr, left[i])
			i++
		} else {
			newarr = append(newarr, right[y])
			y++
		}
	}

	newarr = append(newarr, left[i:]...)
	newarr = append(newarr, right[y:]...)

	return newarr

}

func removeDup(arr []int) []int {
	mp := make(map[int]bool)
	var result []int

	for _, num := range arr {
		if _, y := mp[num]; !y {
			result = append(result, num)
			mp[num] = false
		}
	}

	return result

}

func revWords(str string) string {
	// strSlice := strings.Fields(str)
	var newStr string

	for _, letter := range str {
		newStr = string(letter) + newStr
	}

	return newStr
}

func main() {
	arr := []int{4, 5, 2, 3, 7, 2, 3}
	result := mergeSort(arr)
	remDup := removeDup(result)

	str := "This is a house"

	newStr := revWords(str)

	fmt.Println("The sorted array is: ", result)
	fmt.Println("The array without dup is: ", remDup)
	fmt.Println("The new String reversed is: ", newStr)

}
