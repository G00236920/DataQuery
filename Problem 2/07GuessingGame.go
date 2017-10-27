package main

import (	
	"time"
	"fmt"												//Imports
	"net/http"
	"text/template"
	"math/rand"
	"strconv"
)

type TemplateData struct {
	Message string
	Guess string
	Status string
	Button string
}

func templateHandler(w http.ResponseWriter, r *http.Request) {	//Handle Http requests

		// Try to read the cookie.
		var cookie, err = r.Cookie("Target")

		var usersGuess = r.FormValue("guess")

		fmt.Printf("\nUser Guessed: %s",usersGuess)
		
		if err == nil {
			// If we could read it, try to convert its value to an int.
			if(Target == 0){
				rand.Seed(time.Now().UnixNano())
				Target = rand.Intn(20) + 1
				fmt.Printf("\nRandom Number is: %s",strconv.Itoa(Target))
			}else{
				fmt.Printf("\nRandom Number is: %s",strconv.Itoa(Target)) 
			}
		}
	
		// Create a cookie instance and set the cookie.
		cookie = &http.Cookie{
	
			Name: "Target",
			Value: "Target",
	
		}

		http.SetCookie(w, cookie)

		textGuess := TemplateData{Guess: usersGuess}

		status := compare(textGuess)

	t, _ := template.ParseFiles("template/guess.html")											//Parse the template File
	t.Execute(w, TemplateData{Message: "Guess a Number Between 1 and 20", Guess: usersGuess, Status: status, Button: button})						// Execute the Tmpl file

}

var Target int = 0
var count int = 0
var button string = "Guess"

func main() {

	fmt.Printf("\nRandom Number is: %s",strconv.Itoa(Target))
	http.Handle("/", http.FileServer(http.Dir("./")))		//Handle http request
	http.HandleFunc("/guess",templateHandler)				//handle requests for templates
	
	http.ListenAndServe(":8080", nil)						//Listen and report from port 8080

}

func compare(textGuess TemplateData) string{

	guessInt  , _ := strconv.Atoi(textGuess.Guess)

	if(count == 0){
		button ="Guess"
		count++
		return "You Have Not Guessed Yet!"
	}else if(guessInt == Target){
		Target = rand.Intn(20) + 1
		count = 0
		button ="New_Game"
		return "You Have Guessed Correctly"
	}else if (guessInt > Target){
		return "You Have Guessed Too High"
	}else if (guessInt < Target){
		return "You Have Guessed Too Low"
	}else{
		return "You Have Not entered an Integer Value"
	}

}

//strconv.Itoa()