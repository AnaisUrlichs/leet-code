package main

import "fmt"

func funcPrint(strfn func(string) string) {
	fmt.Println(strfn("Hello World"))
}

func main() {

	myString := func(str string) string {
		return str
	}

	funcPrint(myString)
}
