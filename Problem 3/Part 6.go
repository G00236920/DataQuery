package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

//list of inputs
var str = []string{
	"People say I look like both my mother and father.",
	"Father was a teacher.",
	"I was my Father’s favourite.",
	"I'm looking forward to the weekend.",
	"My grandfather was French!",
	"I am happy.",
	"I am not happy with your responses.",
	"I am not sure that you understand the effect that your questions are having on me.",
	"I am supposed to just take what you're saying at face value?",
	"Why can't I drive my car without a license",
	"Why don't you go for a walk?",
	"Are you human?",
}

var psychobable = [][]string{
	{`Why don\'?t you ([^\?]*)\??`,
		"Perhaps eventually I will $1?."},

	{`Why can\'?t I ([^\?]*)\??`,
		"Do you think you should be able to $1??"},

	{`(?i)i\'?(?:\s?am|m)([^.?!]*)[.?!]?`,
		"How do you know you are $1??"},

	{`Are you ([^\?]*)\??`,
		"I may be $1? -- what do you think?"},

	{`\bfather\b`,
		"Why don’t you tell me more about your father?"},
}

func main() {

	for i := 0; i < 12; i++ {

		//User Input strings
		fmt.Println("intput: " + str[i] + "\n")

		//Response
		fmt.Println("output: " + ElizaResponse(str[i]) + "\n")

	}

}

//Response function
func ElizaResponse(input string) string {

	var counter int

	for range psychobable {

		re := regexp.MustCompile("(?i)" + psychobable[counter][0]) //Read the find index in the row - for the question

		if matched := re.MatchString(input); matched { //Compare question with the users input

			str := re.ReplaceAllString(input, psychobable[counter][1]) //Replace the question with the answer for output

			return str

		}

		counter++ //Increment the index

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

func Reflect(input string) string {

	// Split the input on word boundaries.
	boundaries := regexp.MustCompile("(?i)" + `\b`)
	tokens := boundaries.Split(input, -1)

	// List the reflections.
	reflections := [][]string{
		{`you're `, `i'm`},
		{`your`, `my`},
		{`my`, `your`},
		{`you`, `I`},
		{`me`, `you`},
	}

	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}

	// Put the tokens back together.
	return strings.Join(tokens, ``)
}
