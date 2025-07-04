package main

import "fmt"

func main() {
	// var numbers [5]int

	// // fmt.Println(numbers)

	// numbers[4] = 20

	// // fmt.Println(numbers)

	// // fruits := [3]string{"Apple", "Banana", "grape"}

	// // fmt.Println(fruits[2])

	// coppieArray[0] = 1000

	// fmt.Println(coppieArray)
	// fmt.Println(originalArray)

	// for index, value := range numbers {
	// 	fmt.Printf("index %d  value %d\n", index, value)
	// }

	// a, b := someFunction()

	// fmt.Println(a)
	// fmt.Println(b)

	// var matrix [3][3]int = [3][3]int{
	// 	{1, 2, 3},
	// 	{1, 2, 3},
	// 	{1, 2, 3},
	// }

	// fmt.Print(matrix)

	originalArray := [3]int{1, 2, 3}
	var coppieArray *[3]int

	coppieArray = &originalArray
	coppieArray[0] = 100

	fmt.Println(originalArray)
}

func someFunction() (int, int) {
	return 1, 2
}
