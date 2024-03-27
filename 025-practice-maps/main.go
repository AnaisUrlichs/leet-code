package main

import "fmt"

func main() {
	mp := map[string]int{
		"Apple":     2,
		"Banana":    3,
		"Blueberry": 6,
	}

	mp["Orange"] = 3

	_, found := mp["Banana"]

	if found {
		fmt.Println("Banana exists")
	}

	delete(mp, "Blueberry")

	fmt.Println(mp)
}
