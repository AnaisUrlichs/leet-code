package main

// https://leetcode.com/problems/reverse-integer/description/

import (
	"fmt"
	"strings"
	"unicode"
)

func reverseStringInt(input string) string {
	var result string

	toSlice, specChar := sliceString(input)
	fmt.Println(toSlice)
	revSlice := revTheSlice(toSlice)
	backToString := strings.Join(revSlice, "")
	result = specChar + backToString

	return result
}

func sliceString(input string) (result []string, specChar string) {
	for _, num := range input {
		if unicode.IsDigit(num) {
			result = append(result, string(num))
		} else {
			specChar = string(num)
		}
	}
	return result, specChar
}

func revTheSlice(input []string) []string {
	inputLast := len(input) - 1
	inputFirst := 0

	if input[inputLast] == "0" {
		inputLast--
	}

	for inputFirst < inputLast {
		input[inputFirst], input[inputLast] = input[inputLast], input[inputFirst]
		inputFirst++
		inputLast--
	}

	return input
}

func main() {
	input := "-123"
	result := reverseStringInt(input)
	fmt.Printf("The result of reeversing the string is \n%v", result)
}
