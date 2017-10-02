package main

import (
	"fmt"
)

func main(){

	factorial(9, 9)

}

func factorial(num int, total int) int {

	num = num - 1
	total = total * num

	if num != 1 {
		factorial(num, total)
	} else if num == 1 {
		fmt.Println(total)
		return total
	}

	return 0
}