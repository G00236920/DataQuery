package main

import "fmt"

func main() {

	menu()

}

func menu() {

	var choice int
	var entry string

	for i := 0; i < 1; {

		fmt.Println()
		fmt.Println("Please Choose an Option")
		fmt.Println("\t1. Test for a Palindrome")
		fmt.Println("\t2. Exit")
		fmt.Println()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println()
			fmt.Println("Enter the word you wish to use")
			fmt.Scan(&entry)
			fmt.Println()

			testPal(entry)
			fmt.Println()

		case 2:
			return
		default:
			fmt.Println()
			fmt.Println("INVALID OPTION")
			fmt.Println()
		}

	}
}

func testPal(entry string) {

	z := 0

	for i := len(entry) - 1; i >= 0; i-- {

		if (string(entry[i])) != string(entry[z]) {

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
