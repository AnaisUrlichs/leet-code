package main

import (
	"fmt"
)

func isPalindrome(inputString string) bool {
	left := 0
	right := len(inputString) - 1
	for left < right {
		if inputString[left] != inputString[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	str := "abcca"
	isPalindrome := isPalindrome(str)

	fmt.Println("The string is Palindrome ", isPalindrome)
}
