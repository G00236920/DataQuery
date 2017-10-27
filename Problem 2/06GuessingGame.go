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
}

func templateHandler(w http.ResponseWriter, r *http.Request) {	//Handle Http requests

		// Try to read the cookie.
		var cookie, err = r.Cookie("Target")	//Variables for cookie and its errors
		
		var usersGuess = r.FormValue("guess")

		fmt.Printf("\nUser Guessed: %s",usersGuess)
		
		if err == nil {
			// If we could read it, try to convert its value to an int.
			if(Target == 0){												//if Target has no random value yet
				rand.Seed(time.Now().UnixNano())							//Seed using clock
				Target = rand.Intn(19) + 1									//Random Number
				fmt.Printf("\nRandom Number is: %s",strconv.Itoa(Target))	//Convert to string just because
			}else{
				fmt.Printf("\nRandom Number is: %s",strconv.Itoa(Target)) 	//Convert to string just because
				
			}
		}
	
		// Create a cookie instance and set the cookie.
		cookie = &http.Cookie{												//Cookie Response	
	
			Name: "Target",													//Name of Cookie
			Value: "Target",												//Value Cookie Holds
	
		}

		http.SetCookie(w, cookie)											//Set the Cookie


	t, _ := template.ParseFiles("template/guess.html")											//Parse the template File
	t.Execute(w, TemplateData{Message: "Guess a Number Between 1 and 20",Guess: usersGuess})						// Execute the Tmpl file

}

var Target int = 0

func main() {

	fmt.Printf("\nRandom Number is: %s",strconv.Itoa(Target))
	http.Handle("/", http.FileServer(http.Dir("./")))		//Handle http request
	http.HandleFunc("/guess",templateHandler)				//handle requests for templates
	
	http.ListenAndServe(":8080", nil)						//Listen and report from port 8080

}

//strconv.Itoa()