package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {

	menu()

}

func menu() {

	fmt.Println("Please Choose an Option")
	fmt.Print("1. Find Factorial: \n")
	fmt.Print("2. Exit: \n")

	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		option1()
	case 2:
		return
	default:
		fmt.Println("Incorrect Entry, Try Again")
	}

	menu()
}

func option1() {

	total := new(big.Int)

	fmt.Println("This Program will calculate the Factorial of any Integer Number Entered")
	fmt.Print("Enter A Number Please: \n")

	var x int
	fmt.Scan(&x)

	str := strconv.Itoa(x)

	total.SetString(str, 10)

	fmt.Printf("The Factorial of %d is:\n%d\n", total, factorial(total, x))

}

func factorial(total *big.Int, x int) *big.Int {

	x = x - 1

	if x == 0 {

		return total

	}

	y := new(big.Int)
	y.SetString(strconv.Itoa(x), 10)

	total = big.NewInt(0).Mul(total, y)

	return factorial(total, x)
}
