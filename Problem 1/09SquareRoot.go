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
	var entry float64

	for i := 0; i < 1; {

		fmt.Printf("\nPlease Choose an Option\n")
		fmt.Println("\t1. Find the Square root of a value")
		fmt.Println("\t2. Exit")
		fmt.Println()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter the Number you wish to find the square root of")
			fmt.Scan(&entry)

			Square := sqrt(entry)
			fmt.Printf("\nThe square root of %.2f is %.4f\n", entry, Square)
		case 2:
			return
		default:
			fmt.Printf("\n\nINVALID OPTION\n\n")
		}

	}
}

func sqrt(x float64) float64 {

	value := 1.0
	minValue := 0.00000000001
	guess := 1.0

	for i := 0; math.Abs(guess) > minValue; i++ {

		guess = ((value * value) - x) / (2 * value) //Newtons Equation for square root
		value = value - guess                       //Subtract the Guess from our Value

	}

	return value

}
