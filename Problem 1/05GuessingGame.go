package main

import (
	"fmt"
)

func main(){

	var randNum int 
	var prevGuess int
	var guess int
	var count int

    for i := 0; i <= 1; {

		fmt.Printf("Please Guess a Number between 1 and %d\n", randNum)

		if guess == prevGuess{

			fmt.Println("You have guess the wrong number again")
			
		} else{
			
			count++

			if guess == randNum {

				fmt.Println("WELL DONE!, You have Guessed Correctly")
				i++

			} else if guess > randNum{

				fmt.Println("Your Guess was Too high")

			} else{

				fmt.Println("Your Guess was Too low")

			}
		}

		fmt.Printf("You have made %d Guesses\n", count)

    }

}