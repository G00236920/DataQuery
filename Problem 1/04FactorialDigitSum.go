package main

import (
	"fmt"      //fmt import
	"math/big" //Math Import
	"strconv"  //string conversion import
)

func main() {

	menu() //Menu Function call

}

func menu() {

	//Option menu
	fmt.Println("Please Choose an Option")
	fmt.Print("1. Find Factorial: \n")
	fmt.Print("2. Exit: \n")

	var input int
	fmt.Scan(&input) //User input

	switch input { //Switch which decides menu ouput
	case 1:
		option1() //Option 1 function
	case 2:
		return //return from menu function to end program
	default:
		fmt.Println("Incorrect Entry, Try Again")
	}

	menu() //re call menu function
}

func option1() { //Option 1

	total := new(big.Int) //Converts value to a big integer

	//Request for user input
	fmt.Println("This Program will calculate the Factorial of any Integer Number Entered")
	fmt.Print("Enter A Number Please: \n")

	var x int
	fmt.Scan(&x) //User input stored as x

	str := strconv.Itoa(x) //convert x to a string

	total.SetString(str, 10) //convert x to an integer

	fmt.Printf("The Factorial of %d is:\n%d\n", total, factorial(total, x)) //output value

}

func factorial(total *big.Int, x int) *big.Int { //factorial function, outputs a big integer, takes in a big integer

	x = x - 1

	if x == 0 {
		//ends the recursive function when x is 0, ending the multiplication

		return total
		//return value

	}

	y := new(big.Int)                //declares big integer
	y.SetString(strconv.Itoa(x), 10) //converts x to integer then to y as a string

	total = big.NewInt(0).Mul(total, y) //multiply total by y,(which is the decreasing value)

	return factorial(total, x) //call the factorial function using recursion, until x reaches 0, then ends returning the total
}
