package main

import (
	"fmt"
	"maps"
)

func main() {
	myMapp := make(map[string]int)

	// fmt.Println(myMapp)

	myMapp["key1"] = 9
	myMapp["code"] = 18

	// fmt.Println(myMapp["code"])

	// delete(myMapp, "key1")

	// fmt.Println(myMapp)

	// myMapp["key1"] = 9
	// myMapp["key2"] = 3
	// myMapp["key3"] = 32
	// fmt.Println(myMapp)

	// // clear(myMapp)

	// fmt.Println(myMapp)

	// value, unknownValue := myMapp["key1"]

	// fmt.Println(value)
	// fmt.Println(unknownValue)

	myMapp2 := map[string]int{"a": 1, "b": 2}
	myMapp3 := map[string]int{"a": 1, "b": 2}

	fmt.Println(myMapp2)

	if maps.Equal(myMapp, myMapp2) {
		fmt.Println("my map 1 and map 2 is equal")
	} else {
		fmt.Println("nah")
	}
	if maps.Equal(myMapp3, myMapp2) {
		fmt.Println("my map 3 and map 2 is equal")
	}

	for _, v := range myMapp3 {
		fmt.Println(v)
	}

	myMap5 := make(map[string]map[string]int)

	myMap5["map1"] = myMapp2
	myMap5["map2"] = myMapp3

	fmt.Print(myMap5)

}
