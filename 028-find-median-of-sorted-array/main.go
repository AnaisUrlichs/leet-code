package main

import "fmt"

var median float64

func findMedian(arr []int) (float64, error) {

	sortedArr := mergeSort(arr)
	fmt.Println(sortedArr)
	lengthArr := len(sortedArr)
	center := lengthArr / 2

	if center == 1 {
		median = float64((sortedArr[center-1] + sortedArr[center]) / 2)
	} else {
		median = float64(sortedArr[center])
	}

	return median, nil
}

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
	newarr := make([]int, 0, len(left)+len(left))
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

func main() {
	arr := []int{1, 2, 3, 4, 6, 5, 7}

	sortedArr := mergeSort(arr)
	result, err := findMedian(sortedArr)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The median of the array is", result)
}
