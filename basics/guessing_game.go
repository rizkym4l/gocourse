package main

import (
	"Math/rand"
	"fmt"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(10) + 1

	fmt.Println("Welcome to Guessing Game")
	fmt.Println("I alredy Take a number between 1 - 10")
	fmt.Println("Can u guess it?")

	var guess int

	for {
		fmt.Println("Input your answer:")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("you are Correct, Get happy then")
			break
		} else if guess < target {
			fmt.Println("Too low")

		} else {
			fmt.Println("too high")

		}
	}
}
