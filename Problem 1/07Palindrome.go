package main

import "fmt"

func main() {

	menu() //call menu function

}

func menu() { //menu function

	var choice int
	var entry string

	for i := 0; i < 1; { //loop until i incremented

		//Print output info
		fmt.Println()
		fmt.Println("Please Choose an Option")
		fmt.Println("\t1. Test for a Palindrome")
		fmt.Println("\t2. Exit")
		fmt.Println()
		fmt.Scan(&choice)

		//Switch to decide what happens with user input
		switch choice {
		case 1:

			//Ask the user for their word
			fmt.Println()
			fmt.Println("Enter the word you wish to use")
			fmt.Scan(&entry)
			fmt.Println()

			//Call testPal, to test if it is a palindrome
			testPal(entry)
			fmt.Println()

		case 2:
			return //return from the menu function, ending the program
		default:
			fmt.Println()
			fmt.Println("INVALID OPTION")
			fmt.Println()
		}

	}
}

func testPal(entry string) { //Test for Palindrome function

	z := 0

	for i := len(entry) - 1; i >= 0; i-- {

		if (string(entry[i])) != string(entry[z]) { //Compare each character in the string

			fmt.Println()
			fmt.Println("This Word is NOT a Palindrome")
			fmt.Println()
			return

		} else if i == z {
			fmt.Println()
			fmt.Println("YAY! This Word is Definitely a Palindrome")
			fmt.Println()
		}
		z++
	}

}
