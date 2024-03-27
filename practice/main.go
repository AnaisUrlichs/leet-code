package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 5, 6},
		{6, 3, 8},
	}

	for i, arr := range matrix {
		fmt.Printf("The array %v, has the following numbers: \n", i)
		for _, num := range arr {
			fmt.Println(num)
		}
	}
}
