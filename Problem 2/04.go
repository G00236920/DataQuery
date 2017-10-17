package main

import (
	"fmt" //import Fmt
)

func main() {

	menu() //Call menu function

}

func menu() {

	fmt.Println("")
	fmt.Print("Enter: \n1. Option 1\n2. Option 2 \n3. Exit \n")

	var x int    //Integer value
	fmt.Scan(&x) //Take input from user

	if x == 1 {
		option1()
	} else if x == 2 {
		option2()
	} else {
		return //End program by returning from the menu function
	}

	menu()

}

func option1() {

}

func option2() {

}
