package main

// leet code exercise https://leetcode.com/problems/greatest-common-divisor-of-strings/description/?envType=study-plan-v2&envId=leetcode-75

import "fmt"

func greatestCommon(strOne, strTwo string) []string {
	var result []string

	sliceOne := strToSlice(strOne)
	sliceTwo := strToSlice(strTwo)

	mp := make(map[string]bool)

	if sliceOne[0] == sliceTwo[0] {
		result = append(result, sliceOne[0])
	}

	for i := 1; i < len(sliceOne); i++ {
		for y := 1; y < len(sliceOne); y++ {
			if sliceOne[i] == sliceTwo[y] {
				if sliceOne[i-1] == sliceTwo[y-1] && len(result) >= 1 {
					if _, m := mp[sliceOne[i]]; !m {
						mp[sliceOne[i]] = true
						result = append(result, sliceOne[i])
					}
				}
			}
		}
	}

	return result
}

func strToSlice(instr string) []string {
	var outputSlice []string
	for _, letter := range instr {
		outputSlice = append(outputSlice, string(letter))
	}

	return outputSlice
}

func main() {
	strOne := "ABCAB"
	strTwo := "ABABCABAB"

	result := greatestCommon(strOne, strTwo)
	fmt.Println(result)
}
