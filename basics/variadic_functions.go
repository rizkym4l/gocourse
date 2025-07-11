package main

import "fmt"

func main() {

	statement, total := sum("Ini loh totalnya", 3, 4, 5, 6, 3, 1, 3, 4, 8)
	fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5}

	ora, jelas := sum("iki oran jelss", numbers...)

	fmt.Println(ora, jelas)
}
func sum(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}

	return returnString, total
}
