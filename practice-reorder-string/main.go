package main

import (
	"errors"
	"fmt"
	"strings"
)

func reverseWord(str string) (string, error) {
	var result string

	if str == "" {
		err := errors.New("The string provided is invalid")
		return str, err
	}

	for _, letter := range str {
		result = string(letter) + result
	}

	return result, nil
}

func reverseString(str string) (string, error) {
	strArr := strings.Fields(str)
	var revArr []string
	strArr = reverseArr(strArr)

	if str == "" {
		err := errors.New("The string provided is invalid")
		return str, err
	}

	for _, word := range strArr {
		revWord, err := reverseWord(word)

		if err != nil {
			return str, err
		}

		revArr = append(revArr, revWord)
	}

	result := strings.Join(revArr, " ")

	return result, nil
}

func reverseArr(arr []string) []string {
	first := 0
	last := len(arr) - 1

	for first < last {
		arr[first], arr[last] = arr[last], arr[first]
		first++
		last--
	}

	return arr
}

func main() {
	str := ""

	result, err := reverseString(str)

	if err != nil {
		fmt.Printf("New Error: %s\n", err)
	} else {
		fmt.Println("The reversed String is: \n", result)
	}
}
