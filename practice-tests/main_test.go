package addition

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	have := Add(1, 3)
	want := 4

	if have != want {
		t.Errorf("We want %q but we have %q", want, have)
	} else {
		fmt.Println("The test passed!")
	}

}
