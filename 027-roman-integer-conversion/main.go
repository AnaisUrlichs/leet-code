package main

// Leetcode premium https://leetcode.com/problems/roman-to-integer/description/?envType=study-plan-v2&envId=google-spring-23-high-frequency
// https://medium.com/@anuragsahani0123/roman-to-integer-solution-in-golang-22e156ebe7f9

import "fmt"

func convertToInt(s string) (resultInt int) {

	know := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	lengthOfString := len(s)
	lastElement := s[len(s)-1 : lengthOfString]
	var result int
	result = know[lastElement]
	for i := len(s) - 1; i > 0; i-- {
		if know[s[i:i+1]] <= know[s[i-1:i]] {
			result += know[s[i-1:i]]
		} else {
			result -= know[s[i-1:i]]
		}
	}
	return result
}

func main() {
	romanStr := "XVII"

	resultInt := convertToInt(romanStr)

	fmt.Println(resultInt)

}
