package main

import (
	"fmt"
	"slices"
)

func main() {
	// var sliceName []ElementType

	// var numbers []int
	// var numbers1 = []int{1, 2, 3}

	// numbers2 = []int{9, 8, 7}

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:5]

	slice = append(slice, 4, 4)
	// fmt.Println(slice)

	sliceCoppy := make([]int, len(slice))

	copy(sliceCoppy, slice)
	fmt.Println("Iki coppy Loh", sliceCoppy)

	for i, v := range slice {
		fmt.Println(i, v)
	}

	if slices.Equal(slice, sliceCoppy) {
		fmt.Println("wong sama kok ")
	}

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1               // 1 2
		twoD[i] = make([]int, innerLen) //twoD[0] bikin slice [] twoD[1] = []
		for j := 0; j < innerLen; j++ { // u=innernya 2
			twoD[i][j] = i + j //twoD[0][0] == 0 TwoD[1][0] == 1 2
		}
	}

	fmt.Print(twoD)
}
