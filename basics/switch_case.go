package main

import (
	"fmt"
)

func main() {
	// switch expression {
	// case value1:

	// }

	// fruit := "dsadas"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("its an Apple")
	// case "banana":
	// 	fmt.Println("its a banana")
	// default:
	// 	fmt.Println("unknown Fruit")
	// }

	// day := "Sunday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday":
	// 	fmt.Println("weekday")
	// case "Sunday":
	// 	fmt.Println("WeekEnd")
	// }

	num := 2

	switch {
	case num > 1:
		fmt.Println("Greaterthan 1")
		fallthrough
	case num == 2:
		fmt.Println("it is 2")
	default:
		fmt.Println("Idk")

	}
	checkType(false)
}

func checkType(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println(" it is an Integer DUde")
	case float64:
		fmt.Println("this is an FLoat 64 geys")
	case string:
		fmt.Println("this is a String babe")
	default:
		fmt.Println("Ora Jelas")
	}
}
