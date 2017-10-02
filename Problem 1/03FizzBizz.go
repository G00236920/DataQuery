package main

import (
	"fmt"
)

func main(){

	for i := 1;  i<=100; i++ {
		
		if i % 3 == 0 && i % 5 == 0{
			fmt.Println("FizzBizz ")
		}else if i % 3 == 0 {
			fmt.Println("Fizz ")
		} else if i % 5 == 0 {
			fmt.Println("Bizz")
		}else{
			fmt.Println("", i)
		}
	}
}