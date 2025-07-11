package main

import "fmt"

func main() {
	process(209)

	process(-3)
}

func process(input int) {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer2")
	if input < 0 {
		fmt.Println("before Panic")
		panic("input must be non negative")
		// fmt.Println("After Panic")
		// defer fmt.Println("Defer 3")

	}
	fmt.Println("PRossesing Inputs", input)
}
