// Created By: Charlotte Jhu
// Created On: May 2023
//
// This program adds all the numbers between 1 and the user's inputted number

package main

import (
	"fmt"
)

func main() {
	// This function adds all the numbers between 1 and the user's inputted number
	// variables
	var number int
	var counter int
	var answer = 0

	// input
	fmt.Println("This program adds all the numbers between 1 and your inputted number")
	fmt.Print("Enter your number: ")
	fmt.Scanln(&number)

	// process
	for counter < number {
		answer += counter
		counter++
	}
	answer += number

	// output
	fmt.Println("")
	fmt.Println("The sum of all the numbers between 1 and", number, "is", answer)
	fmt.Println("\nDone.")
}
