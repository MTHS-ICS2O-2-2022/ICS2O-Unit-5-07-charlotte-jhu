// Created By: Charlotte Jhu
// Created On: May 2023
//
// This function does multiplication

package main

import (
	"fmt"
)

func main() {
	// This function does multiplication
	// variables
	var number1 int
	var number2 int
	var counter = 0
	var answer = 0

	// input
	fmt.Println("Let's do multiplication!")
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&number1)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&number2)

	// process
	for counter < number2 {
		answer += number1
		counter++
	}

	// output
	fmt.Println("")
	fmt.Println(number1, "x", number2, "is", answer)
	fmt.Println("\nDone.")
}
