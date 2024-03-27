package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) []string {
	arrStr := strings.Fields(str)
	var reversedArr []string
	for _, words := range arrStr {
		var resultStr string
		for _, letter := range words {
			resultStr = string(letter) + resultStr
		}
		reversedArr = append(reversedArr, resultStr)
	}
	return reversedArr
}

func reverseWordOrder(str string) (result string) {
	arrStr := strings.Fields(str)
	first := 0
	last := len(arrStr) - 1
	for first < last {
		arrStr[first], arrStr[last] = arrStr[last], arrStr[first]
		first++
		last--
	}
	result = strings.Join(arrStr, " ")
	return
}

func main() {
	str := "This is a String"
	fmt.Println("Original input string ", str)
	reversedStr := reverseString(str)
	fmt.Println("Reversed string ", reversedStr)
	reverseOrder := reverseWordOrder(str)
	fmt.Println("Reversed word order ", reverseOrder)
}
