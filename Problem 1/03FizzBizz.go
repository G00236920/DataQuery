package main

import ( //Imports
	"fmt"
)

func main() { //Main Function

	menu() //Call menu function

}

func menu() { //menu function

	fmt.Println("Please Choose an Option") //Print menu options
	fmt.Print("1. Play FizzBizz: \n")
	fmt.Print("2. Exit: \n")

	var input int //user input
	var fizz int  //value for what counts as a  fizz value
	var bizz int  //value for what counts as a  bizz value
	//fizzbizz will be calculated based on which values are dividable by both fizz and bizz

	fizz = 5
	bizz = 3

	fmt.Scan(&input) //Input Value

	switch input { //switch
	case 1:
		output(fizz, bizz) //output function
	case 2:
		return //leave the menu function
	default:
		fmt.Println("Incorrect Entry, Try Again")
	}

	menu()
}

func output(fizz int, bizz int) { //output function

	for i := 1; i <= 100; i++ { //loop through 100 values

		if i%fizz == 0 && i%bizz == 0 { //if it divides evenly into both values, prints fizzbizz
			fmt.Println("FizzBizz")
		} else if i%fizz == 0 {
			fmt.Println("Fizz") //fizz print
		} else if i%bizz == 0 {
			fmt.Println("Bizz") //bizz print
		} else {
			fmt.Println("", i)
		}
	}
}
