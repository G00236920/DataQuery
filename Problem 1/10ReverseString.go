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
		fmt.Println("\t1. Enter a new String")
		fmt.Println("\t2. Exit")
		fmt.Println()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println()
			fmt.Println("The String you wish to Reverse")
			fmt.Scan(&entry)
			fmt.Println()
			fmt.Print("This Word Reversed is ")

			for i := len(entry) - 1; i >= 0; i-- {
				fmt.Print(string(entry[i]))
			}
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
