package main

import (
	"fmt"
)

func findCommonValue(arrOne, arrTwo []int) (result int, err error) {

	var tmp []int

	for _, numOne := range arrOne {
		for _, numTwo := range arrTwo {
			if numOne == numTwo {
				tmp = append(tmp, numOne)
			}
		}
	}

	result, err = FindSmallest(tmp)

	if err != nil {
		return 0, err
	}

	return result, nil
}

func FindSmallest(arr []int) (result int, err error) {
	tmp := arr[0]
	for i := 1; i < len(arr); i++ {
		secVal := arr[i]
		if secVal < tmp {
			tmp = secVal
		}
	}

	return tmp, nil
}

func main() {
	arrOne := []int{1, 5, 3}
	arrTwo := []int{5, 3}

	result, err := findCommonValue(arrOne, arrTwo)

	if err != nil {
		fmt.Printf("There has been an error %s\n", err)
	}

	fmt.Println("The result is \n", result)
}
