package main

import (
	"errors"
	"fmt"
)

func createSequence(arr []int, length int) ([]int, error) {

	if len(arr) <= 2 {
		err := errors.New("The array provided is to short")
		return arr, err
	}

	i := 0

	for i < length {
		lastOne := len(arr) - 1
		lastTwo := len(arr) - 2

		one := arr[lastOne]
		two := arr[lastTwo]

		new := one + two

		arr = append(arr, new)

		i++
	}

	return arr, nil
}

func main() {
	start := []int{0, 1, 2}
	length := 10
	result, err := createSequence(start, length)

	if err != nil {
		fmt.Printf("There has been an error %s\n", err)
	} else {
		fmt.Print("The sequence is: \n", result)
	}
}
