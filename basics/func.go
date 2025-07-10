package main

import "fmt"

func main() {
	fmt.Println(add(1, 3))

	result :=
		applyOperation(9, 7, add)

	fmt.Println("add : ", result)

	multiplyBy2 := createMultiaplier(2)

	println("multiply func : ", multiplyBy2(4))
}

func add(a, b int) int {
	return a + b
}

func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiaplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
