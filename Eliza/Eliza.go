package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"text/template"
)

var user string

var mtx sync.Mutex

//With this Project we had the Oppertunity to use Bootstrap for the implementation of the GUI in html format
//I decided to Design my own instead

func main() {

	http.Handle("/", http.FileServer(http.Dir("./"))) //Handle http request
	http.HandleFunc("/chat", templateHandler)         //handle requests for templates
	http.HandleFunc("/ajax", receiveAjax)
	http.ListenAndServe(":8080", nil) //Listen and report from port 8080
	//Need response handler. Which receives the message through ajax.
}

func templateHandler(w http.ResponseWriter, r *http.Request) { //Handle Http requests

	t, _ := template.ParseFiles("chat/eliza.html") //Parse the template File
	t.Execute(w, t)                                // Execute the Tmpl file

}

func responseFromEliza(usersAnswer string) string {

	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, usersAnswer); matched {
		return "Why don’t you tell me more about your father?"
	}

	re := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)
	if matched := re.MatchString(usersAnswer); matched {
		return re.ReplaceAllString(usersAnswer, "How do you know you are $1?")
	}

	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	return answers[rand.Intn(len(answers))]

}

func Reflect(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)

	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
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

func receiveAjax(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		Q := r.FormValue("Question")

		fmt.Println(responseFromEliza(Q))

	}
}
