package addition

import "fmt"

func Add(a, b int) (result int) {
	result = a + b
	return result
}

func main() {
	result := Add(1, 3)

	fmt.Println("The result is \n", result)
}
