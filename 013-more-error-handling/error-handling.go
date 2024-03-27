package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (result int, err error) {
	err = errors.New("An error occured")

	if a == 0 {
		return 0, err
	}

	return a / b, nil
}

func main() {
	result, err := divide(0, 5)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
