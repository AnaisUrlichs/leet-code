package main

// Resource https://www.educative.io/answers/how-to-reverse-a-string-in-golang

import (
	"fmt"
	"strings"
)

// reverse single string
func reverseString(str string) (result string) {
	for _, letter := range str {
		result = string(letter) + result
	}

	return result
}

// reverse a slice of strings
func reverseStringSlice(str []string) (result []string) {
	for _, value := range str {
		revStr := reverseString(value)
		result = append(result, revStr)
	}
	return
}

func splitString(str string) (splitString []string) {
	splitString = strings.Fields(str)
	return
}

func main() {
	str := "Educative Espresso"
	fmt.Println(str)
	splitString := splitString(str)
	result := reverseStringSlice(splitString)
	fmt.Println(result)
}
