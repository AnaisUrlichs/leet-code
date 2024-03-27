package main

//Problem statement: Given an array of integers and a value,
//determine if there are any three integers in the array whose sum equals the given value.

import (
	"fmt"
)

var arr []int
var comb []int
var equal int

func findNumbers(arr []int, m int) string {
	for i, value := range arr {
		fmt.Println("M is ", m)
		fmt.Println("First Number is ", value)
		result := m - value
		for y := 0; y < len(arr); y++ {
			if i != y {
				valueOne := arr[y]
				fmt.Println("Second Number is ", valueOne)
				for f := 0; f < len(arr); f++ {
					if i != f {
						valueTwo := arr[f]
						fmt.Println("The Third Number is ", valueTwo)
						currentSum := valueOne + valueTwo
						if currentSum == result {
							comb = append(comb, value)
							comb = append(comb, valueOne)
							comb = append(comb, valueTwo)

							return "Values Found"
						}
					}
				}
			}
		}
	}
	return "No Values Found"
}

func main() {
	arr = []int{2, 5, 2, 4}
	equal = 11
	answer := findNumbers(arr, equal)
	fmt.Println(answer, comb)
}
