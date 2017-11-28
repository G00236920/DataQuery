package main

import (
	"fmt"
	"math/rand"
	"time"
)

//list of inputs
var str = []string{
	"People say I look like both my mother and father.",
	"Father was a teacher.",
	"I was my father’s favourite.",
	"I’m looking forward to the weekend.",
	"My grandfather was French!",
}

func main() {

	for i := 0; i < 4; i++ {

		//User Input strings
		fmt.Println("intput: " + str[i] + "\n")

		//Response
		fmt.Println("output: " + ElizaResponse() + "\n")

	}

}

//Response function
func ElizaResponse() string {

	//answer posibilities
	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	//Seed time
	rand.Seed(time.Now().UTC().UnixNano())

	//Return random answer
	return answers[rand.Intn(len(answers))]

}
