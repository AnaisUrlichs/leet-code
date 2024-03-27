package main

import (
	"fmt"
	"strings"
)

func removeDup(str []string) (output []string, err error) {
	for _, town := range str {
		// lowercase the string
		toLower := strings.ToLower(town)
		// separate by whitespace
		lowerStr := strings.Fields(toLower)
		// rejoin without whitespace
		resultString := strings.Join(lowerStr, "")
		// append new string
		output = append(output, resultString)
	}

	return
}

func findNumTowns(str []string) (result int, err error) {
	mp := make(map[string]bool)

	for _, town := range str {
		if _, y := mp[town]; !y {
			mp[town] = true
			result++
		}
	}

	return
}

func main() {
	str := []string{"town TOYOTOTO", "Town tretetetet", "City BOBOBOB", "Place TWTTWTWTW", "town TOYOTOTO"}
	removedDup, err := removeDup(str)
	result, err := findNumTowns(removedDup)

	if err != nil {
		fmt.Printf("Something went wrong %s\n", err)
	}

	fmt.Println("The number of towns in the list is: ", result)
}
