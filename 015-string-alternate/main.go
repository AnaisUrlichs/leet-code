package main

import (
	"fmt"
)

// String Alternate question https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75

func stringAlternate(strOne, strTwo string) []string {
	var result []string
	i := 0

	var sliceOne []string
	var sliceTwo []string

	for _, let := range strTwo {
		sliceTwo = append(sliceTwo, string(let))
	}

	for _, let := range strOne {
		sliceOne = append(sliceOne, string(let))
	}

	lenOne := len(sliceOne)
	lenTwo := len(sliceTwo)

	if i < lenOne {
		for _, letterOne := range sliceOne {
			result = append(result, string(letterOne))
			if i < lenTwo {
				result = append(result, sliceTwo[i])
			}
			i++
		}
		fmt.Println(i)
	}

	if i >= lenOne && i < lenTwo {
		result = append(result, sliceTwo[i])
		i++
	}

	return result
}

func main() {
	strOne := "abc"
	strTwo := "efgh"

	result := stringAlternate(strOne, strTwo)
	fmt.Println(result)
}
