package main

import ( //Multiple imports
	"fmt"
	"math/rand" //math random import
	"time"      //time import
)

func main() {

	game() //Game function, start the game
}

func game() { // Game function

	//variables
	var prevGuess int
	var guess int
	var count int

	rand.Seed(time.Now().UnixNano()) //seed the random value, so that it resets on each run
	randNum := rand.Intn(100-1) + 1  //generate random value, between 1-100

	for i := 0; i <= 1; { //loop until value increased

		//ask the user to input a value
		fmt.Println("Please Guess a Number between 1 and 100 (-1 to quit)")
		fmt.Scan(&guess)
		//User input

		if guess == prevGuess { //if the user enters the same value as previous, the count does not increase

			fmt.Println("You have guess the wrong number again") //user entered the wrong input

		} else {

			count++ //increase count value

			if guess == -1 {
				return //return and end the function
			} else if guess == randNum { //if the user guesses the correct value

				fmt.Println("WELL DONE!, You have Guessed Correctly")
				fmt.Print("\n\n*****   NEW GAME   *****\n\n")
				count = 0
				game() //game function

			} else if guess > randNum {

				fmt.Println("Your Guess was Too high") //user guessed to high

			} else {

				fmt.Println("Your Guess was Too low") //User guessed too low

			}
		}

		fmt.Printf("You have made %d Guesses\n", count) //print the amount of tries made by user

	}

}
