package main

// https://leetcode.com/problems/determine-if-two-strings-are-close/description/?envType=study-plan-v2&envId=leetcode-75

import (
	"fmt"
	"strings"
)

func areClose(strOne, strTwo string) bool {
	result := false

	if len(strOne) != len(strTwo) {
		return result
	}

	sliceOne := convertSlice(strOne)
	sliceTwo := convertSlice(strTwo)

	resultOne := operationOne(sliceOne, sliceTwo)
	resultTwo := operationTwo(sliceOne, sliceTwo)

	if resultOne == true || resultTwo == true {
		return true
	}

	return result
}

func convertSlice(input string) []string {
	input = strings.ToLower(input)
	var result []string
	for _, letter := range input {
		result = append(result, string(letter))
	}
	return result
}

func operationOne(strOne, strTwo []string) bool {
	var outcome bool
	var count int

	for i := 0; i < len(strOne); i++ {
		for y := 0; y < len(strOne); y++ {
			if strOne[i] == strTwo[y] {
				count++
			}
		}
	}

	if count == len(strOne) {
		outcome = true
	}

	return outcome
}

func operationTwo(strOne, strTwo []string) bool {
	mpOne := make(map[string]int)
	mpTwo := make(map[string]int)
	var count int

	for i := 0; i < len(strOne); i++ {
		mpOne[strOne[i]] = mpOne[strOne[i]] + 1
	}

	for y := 0; y < len(strTwo); y++ {
		mpTwo[strTwo[y]] = mpTwo[strTwo[y]] + 1
	}

	for _, valueA := range mpOne {
		for _, valueB := range mpTwo {
			fmt.Println(valueA)
			fmt.Println(valueB)
			if valueA == valueB {
				count++
			}
		}
	}

	if count == len(mpOne) {
		return true
	}

	return false
}

func main() {
	strOne := "aac"
	strTwo := "bbc"

	result := areClose(strOne, strTwo)
	fmt.Println(result)
}
