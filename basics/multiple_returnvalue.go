package main

import "fmt"

func main() {

	first, second := divide(3, 5)

	// fmt.Println(first, second)

	fmt.Printf("A is : %d and b is %d", first, second)
}

func divide(a, b int) (int, int) {
	first := a / b
	second := a % b

	return first, second
}
