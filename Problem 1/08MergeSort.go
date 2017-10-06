package main

import "fmt"

func main() {

	menu()

}

func menu() {

	var choice int

	var a = [3]int{1, 4, 6}
	var b = [3]int{2, 3, 5}

	for i := 0; i < 1; {

		fmt.Println()
		fmt.Println("Please Choose an Option")
		fmt.Println("\t1. Show the 2 sorted Arrays")
		fmt.Println("\t2. Merge and Sort both Arrays")
		fmt.Println("\t3. Exit")
		fmt.Println()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			display(a, b)
		case 2:
			sort(a, b)
		case 3:
			return
		default:
			fmt.Println()
			fmt.Println("INVALID OPTION")
			fmt.Println()
		}

	}
}

func display(a [3]int, b [3]int) {
	fmt.Println()
	fmt.Print("List A is: \n{")

	for i := 0; i < 3; i++ {
		fmt.Print(a[i], ", ")
	}

	fmt.Println("}")
	fmt.Println()
	fmt.Print("List B is: \n{")

	for i := 0; i < 3; i++ {
		fmt.Print(b[i], ", ")
	}

	fmt.Println("}")
}

func sort(a [3]int, b [3]int) {
	var c = [6]int{}

	i, j := 0, 0

	for k := 0; k < 6; k++ {
		if i > len(a)-1 && j <= len(b)-1 {
			c[k] = b[j]
			j++
		} else if j > len(b)-1 && i <= len(a)-1 {
			c[k] = a[i]
			i++
		} else if a[i] < b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
	}

	fmt.Print("List C is: \n{")

	for i := 0; i < 6; i++ {

		fmt.Print(c[i], ", ")

	}

	fmt.Println("}")

}
