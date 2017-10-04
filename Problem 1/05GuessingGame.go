package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	game()
}

func game() {

	var prevGuess int
	var guess int
	var count int

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(100-1) + 1

	for i := 0; i <= 1; {

		fmt.Println("Please Guess a Number between 1 and 100 (-1 to quit)")
		fmt.Scan(&guess)

		if guess == prevGuess {

			fmt.Println("You have guess the wrong number again")

		} else {

			count++

			if guess == -1 {
				return
			} else if guess == randNum {

				fmt.Println("WELL DONE!, You have Guessed Correctly")
				fmt.Print("\n\n*****   NEW GAME   *****\n\n")
				count = 0
				game()

			} else if guess > randNum {

				fmt.Println("Your Guess was Too high")

			} else {

				fmt.Println("Your Guess was Too low")

			}
		}

		fmt.Printf("You have made %d Guesses\n", count)

	}

}
