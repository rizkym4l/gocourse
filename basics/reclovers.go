package main

import "fmt"

func main() {
	// fmt.Println("Hello, world!")
	proccess()
}

func proccess() {
	defer func() {

		if r := recover(); r != nil {
			fmt.Println("recovered", r)
		}
	}()

	fmt.Println("star Proses")

	// panic("somethink when wrong")

	fmt.Print("sdad")

}
