package main

import "fmt"

func main() {
	var a, b int = 10, 22.0

	var result int
	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Substraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Modulus:", result)

	const pi float64 = 22 / 7.0
	fmt.Println(pi)
}
