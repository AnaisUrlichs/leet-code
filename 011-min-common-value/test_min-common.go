package main

import (
	"fmt"
	"testing"
)

func TestFindSmallest(t *testing.T) {
	arr := []int{1, 2}
	have, err := FindSmallest(arr)
	want := 1

	if err != nil {
		fmt.Printf("There has been an error %s\n", err)
	} else {
		if have != want {
			t.Errorf("We wanted %q but we have %q", want, have)
		}
	}
}
