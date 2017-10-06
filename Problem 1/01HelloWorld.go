// https://gobyexample.com/hello-world

package main

import (
	"fmt" //import Fmt
)

func main() {

	menu() //Call menu function

}

func menu() {

	fmt.Println("This Program will say 'Hello World'  in either Japanese or English")
	fmt.Print("Enter: \n1. Japanese \n2. English: \n3. Exit \n")

	var x int    //Integer value
	fmt.Scan(&x) //Take input from user

	if x == 1 {
		fmt.Println("こんにちは, 世界") //Print Hello world in Japanese
	} else if x == 2 {
		fmt.Println("Hello World") //Print hello world in english
	} else {
		return //End program by returning from the menu function
	}

	menu()

}
