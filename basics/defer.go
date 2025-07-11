package main

import "fmt"

func main() {
	process()
}

func process() {
	defer fmt.Println("Ini yang dityunda 1 ")
	defer fmt.Println("Ini yang dityunda 2 ")
	defer fmt.Println("Ini yang dityunda 3")
	fmt.Println("Ini yang kaga")
}
