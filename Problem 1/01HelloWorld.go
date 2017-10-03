// https://gobyexample.com/hello-world

package main

import (
	"fmt"
)

func main() {

	fmt.Println("This Program will say 'Hello World'  in either Japanese or English")
	fmt.Print("Enter: \n1 for Japanese \n2 for English: \n")

	var x int
	fmt.Scan(&x)

	if x == 1 {
		fmt.Println("こんにちは, 世界")
	} else if x == 2 {
		fmt.Println("Hello World")
	}

}
