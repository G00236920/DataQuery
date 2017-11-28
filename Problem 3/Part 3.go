package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

//list of inputs
var str = []string{
	"People say I look like both my mother and father.",
	"Father was a teacher.",
	"I was my Father’s favourite.",
	"I’m looking forward to the weekend.",
	"My grandfather was French!",
	"I am happy.",
	"I am not happy with your responses.",
	"I am not sure that you understand the effect that your questions are having on me.",
	"I am supposed to just take what you’re saying at face value?",
}

func main() {

	for i := 0; i < 9; i++ {

		//User Input strings
		fmt.Println("intput: " + str[i] + "\n")

		//Response
		fmt.Println("output: " + ElizaResponse(str[i]) + "\n")

	}

}

//Response function
func ElizaResponse(input string) string {

	re := regexp.MustCompile(`I am ([^.?!]*)[.?!]?`)

	if matched := re.MatchString(input); matched { //Compare question with the users input

		return re.ReplaceAllString(input, "How do you know you are $1??")

	}

	re2 := regexp.MustCompile("(?i)" + `\bfather\b`) //Read the input for the question

	//If the input mentions Father
	if matched := re2.MatchString(input); matched { //Compare question with the users input

		return "Why don’t you tell me more about your father?"

	}

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
