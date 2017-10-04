package main

import (
	"fmt"
)

func main() {

	menu()

}

func menu() {

	fmt.Println("Please Choose an Option")
	fmt.Print("1. Play FizzBizz: \n")
	fmt.Print("2. Exit: \n")

	var input int
	var fizz int
	var bizz int

	fizz = 5
	bizz = 3

	fmt.Scan(&input)

	switch input {
	case 1:
		output(fizz, bizz)
	case 2:
		return
	default:
		fmt.Println("Incorrect Entry, Try Again")
	}

	menu()
}

func output(fizz int, bizz int) {

	for i := 1; i <= 100; i++ {

		if i%fizz == 0 && i%bizz == 0 {
			fmt.Println("FizzBizz")
		} else if i%fizz == 0 {
			fmt.Println("Fizz")
		} else if i%bizz == 0 {
			fmt.Println("Bizz")
		} else {
			fmt.Println("", i)
		}
	}
}
