package main

import "fmt"

func MakeMultiplier(base int) func(int) int {
	return func(factor int) int {
		fmt.Printf("Thise is base %q and factor %q", base, factor)
		return base * factor
	}
}

func main() {
	result := MakeMultiplier(2)
	fmt.Println("This is the result", result)
}
