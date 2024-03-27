package main

import (
	"errors"
	"fmt"
)

func removeDup(arr []int) (result []int, err error) {
	if len(arr) <= 1 {
		err = errors.New("The array is too short to sort")
		return arr, err
	}

	mp := make(map[int]bool)

	for _, num := range arr {
		if _, y := mp[num]; !y {
			mp[num] = true
			result = append(result, num)
		}
	}

	return result, nil
}

func main() {
	arr := []int{1, 5, 4, 4, 6, 6, 8, 2}

	result, err := removeDup(arr)

	if err != nil {
		fmt.Printf("The error is %s\n", err)
	} else {
		fmt.Println("The arr without duplicates is \n", result)
	}
}
