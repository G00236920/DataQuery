package main

import (
	"fmt" //Imports
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type TemplateData struct {
	Message string
	Guess   string
	Status  string
	Button  string
}

func templateHandler(w http.ResponseWriter, r *http.Request) { //Handle Http requests

	// Try to read the cookie.
	var cookie, err = r.Cookie("Target") //Variables for cookie and its errors

	var usersGuess = r.FormValue("guess") //Get input that user added to text field

	fmt.Printf("\nUser Guessed: %s", usersGuess)

	if err == nil {
		// If we could read it, try to convert its value to an int.
		if Target == 0 { //if Target has no random value yet
			rand.Seed(time.Now().UnixNano())                           //Seed using clock
			Target = rand.Intn(19) + 1                                 //Random Number
			fmt.Printf("\nRandom Number is: %s", strconv.Itoa(Target)) //Convert to string just because
		} else {
			fmt.Printf("\nRandom Number is: %s", strconv.Itoa(Target)) //Convert to string just because

		}
	}

	// Create a cookie instance and set the cookie.
	cookie = &http.Cookie{ //Cookie Response

		Name:  "Target", //Name of Cookie
		Value: "Target", //Value Cookie Holds

	}

	http.SetCookie(w, cookie) //Set the Cookie

	textGuess := TemplateData{Guess: usersGuess} //Change text when User Guesses

	status := compare(textGuess) //Compare Guess with random number

	t, _ := template.ParseFiles("template/guess.html")                                                                        //Parse the template File
	t.Execute(w, TemplateData{Message: "Guess a Number Between 1 and 20", Guess: usersGuess, Status: status, Button: button}) // Execute the Tmpl file

}

var Target int = 0
var count int = 0
var button string = "Guess"

func main() {

	fmt.Printf("\nRandom Number is: %s", strconv.Itoa(Target))
	http.Handle("/", http.FileServer(http.Dir("."))) //Handle http request
	http.HandleFunc("/guess", templateHandler)       //handle requests for templates

	http.ListenAndServe(":8080", nil) //Listen and report from port 8080

}

func compare(textGuess TemplateData) string {

	guessInt, _ := strconv.Atoi(textGuess.Guess)

	if count == 0 {
		button = "Guess"                   //Reset button Text
		count++                            //Increment Game counter
		return "You Have Not Guessed Yet!" //return text to user
	} else if guessInt == Target { //User Guessed Correctly
		Target = rand.Intn(20) + 1          //Generate New random Number
		count = 0                           //Reset Game Counter
		button = "New_Game"                 //New game Button
		return "You Have Guessed Correctly" //return text to user
	} else if guessInt > Target { //User Guessed High
		return "You Have Guessed Too High" //return text to user
	} else if guessInt < Target { //User Guessed Low
		return "You Have Guessed Too Low" //return text to user
	} else {
		return "You Have Not entered an Integer Value" //Catch All - Dosent Work
	}

}

//strconv.Itoa()
