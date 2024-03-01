package main

// Compare longest common prefix in an array of strings https://leetcode.com/problems/longest-common-prefix/description/
// https://medium.com/codex/leetcode-with-golang-longest-common-prefix-89d856b0749b

import (
	"fmt"
	"sort"
	"strings"
)

var arr []string

func longestPrefix(arr []string) string {

	var result string
	var endPrefix = false

	if len(arr) > 0 {

		sort.Strings(arr)

		first := strings.ToLower(arr[0])
		last := strings.ToLower(arr[len(arr)-1])

		for i := 0; i < len(first); i++ {
			if !endPrefix && string(last[i]) == string(first[i]) {
				result += string(last[i])
			} else {
				endPrefix = true
			}
		}
	}

	return result
}

func main() {
	arr = []string{"fly", "flight", "fling"}
	longPref := longestPrefix(arr)
	fmt.Println(longPref)
}
