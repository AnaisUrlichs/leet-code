package main

// Blog post https://www.educative.io/answers/how-to-reverse-an-array-in-go
// https://www.educative.io/answers/how-to-remove-duplicate-values-from-a-slice-in-go

import (
	"fmt"
	"strings"
)

var str []string
var nodup []string
var count int

func reverseString(str []string) {
	left := 0
	right := len(str) - 1

	for left < right {
		strings.ToLower(str[left])
		strings.ToLower(str[right])
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
}

func countNonDuplicate(str []string) {
	mp := make(map[string]bool)

	for _, element := range str {
		if _, y := mp[element]; !y {
			mp[element] = true
			nodup = append(nodup, element)
			count++
		}
	}
}

func main() {
	str = []string{"Apple", "Banana", "Pancake", "Apple"}
	reverseString(str)
	countNonDuplicate(str)
	fmt.Println(str)
	fmt.Println(nodup)
	fmt.Println(count)
}
