package main

import (
	"fmt" //Imports
	"net/http"
	"text/template"
)

type TemplateData struct {
	Message string
}

func templateHandler(w http.ResponseWriter, r *http.Request) { //Handle Http requests

	// Try to read the cookie.
	var cookie, err = r.Cookie("Target") //Variables for cookie and its errors

	var usersGuess = r.FormValue("guess") //Get input that user added to text field

	fmt.Printf("\nUser Guessed: %s", usersGuess)

	if err == nil {

	}

	// Create a cookie instance and set the cookie.
	cookie = &http.Cookie{ //Cookie Response

		Name:  "Target", //Name of Cookie
		Value: "Target", //Value Cookie Holds

	}

	http.SetCookie(w, cookie) //Set the Cookie

	t, _ := template.ParseFiles("template/guess.html")                     //Parse the template File
	t.Execute(w, TemplateData{Message: "Guess a Number Between 1 and 20"}) // Execute the Tmpl file

}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./"))) //Handle http request
	http.HandleFunc("/guess", templateHandler)        //handle requests for templates
	http.ListenAndServe(":8080", nil)                 //Listen and report from port 8080

}
