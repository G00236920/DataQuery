package main

import (
	"fmt"
	"math"
)

func main() {

	menu()

}

func menu() {

	var choice int
	var entry float64 //float Variable

	for i := 0; i < 1; { //infinite loop for menu - thought id change it up

		fmt.Printf("\nPlease Choose an Option\n")
		fmt.Println("\t1. Find the Square root of a value")
		fmt.Println("\t2. Exit")
		fmt.Println()
		fmt.Scan(&choice) // User input

		switch choice {
		case 1:
			fmt.Println("Enter the Number you wish to find the square root of")
			fmt.Scan(&entry) //user input

			Square := sqrt(entry)
			fmt.Printf("\nThe square root of %.2f is %.4f\n", entry, Square) //Output the square root and original value
		case 2:
			return
		default:
			fmt.Printf("\n\nINVALID OPTION\n\n")
		}

	}
}

func sqrt(x float64) float64 {

	value := 1.0
	minValue := 0.00000000001 //limit the amount the square root will trace down as far as
	guess := 1.0

	for i := 0; math.Abs(guess) > minValue; i++ {

		guess = ((value * value) - x) / (2 * value) //Newtons Equation for square root
		value = value - guess                       //Subtract the Guess from our Value

	}

	return value // return the square root

}
