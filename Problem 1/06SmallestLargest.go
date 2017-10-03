package main

import (
	"fmt"
)

func main() {

	//var numOfElements int
	var elements []string

	for i := 0; i < 2; {

		fmt.Printf("Please Enter Element No. %d\n", (i + 1))
		fmt.Print("Enter text: \n")
		fmt.Scan(&elements[i])
		i++

	}
}
