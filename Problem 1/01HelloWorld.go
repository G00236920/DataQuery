// https://gobyexample.com/hello-world

package main

import (
	"fmt"
)

func main() {

	menu()

}

func menu() {

	fmt.Println("This Program will say 'Hello World'  in either Japanese or English")
	fmt.Print("Enter: \n1. Japanese \n2. English: \n3. Exit \n")

	var x int
	fmt.Scan(&x)

	if x == 1 {
		fmt.Println("こんにちは, 世界")
	} else if x == 2 {
		fmt.Println("Hello World")
	} else {
		return
	}

	menu()

}
