package main

// Reverse Words in a sentence https://www.educative.io/blog/apple-coding-interview-questions

import (
	"errors"
	"fmt"
	"strings"
)

func revString(str string) (string, error) {

	err := errors.New("The string is invalid")

	if str == "" {
		return str, err
	}

	strSlice := turnIntoSlice(str)

	if len(strSlice) <= 1 {
		return str, err
	}

	first := 0
	last := len(strSlice) - 1

	for first < last {
		strSlice[first], strSlice[last] = strSlice[last], strSlice[first]
		first++
		last--
	}

	result := turnIntoString(strSlice)

	return result, nil
}

func turnIntoSlice(str string) []string {
	result := strings.Fields(str)
	return result
}

func turnIntoString(input []string) string {
	result := strings.Join(input, " ")
	return result
}

func main() {
	str := "One Two Three Four Five"
	result, err := revString(str)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(result)
}
