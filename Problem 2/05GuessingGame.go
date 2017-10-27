package main

import (	
	"fmt"												//Imports
	"net/http"
	"text/template"
	"math/rand"
	//"strconv"
)

type TemplateData struct {
	Message string
}

func templateHandler(w http.ResponseWriter, r *http.Request) {	//Handle Http requests
		
		Target := 0 
	
		// Try to read the cookie.
		var cookie, err = r.Cookie("Target")
		
		if err == nil {
			// If we could read it, try to convert its value to an int.
			if(Target == 0){

				Target = rand.Intn(19) + 1

			}
		}
	
		// Create a cookie instance and set the cookie.
		cookie = &http.Cookie{
	
			Name: "Target",
			Value: "Target",
	
		}

		http.SetCookie(w, cookie)

		fmt.Printf("\nRandom Number is: %d",Target) 

	t, _ := template.ParseFiles("template/guess.html")											//Parse the template File
	t.Execute(w, TemplateData{Message: "Guess a Number Between 1 and 20"})						// Execute the Tmpl file

}


func main() {

	http.Handle("/", http.FileServer(http.Dir("./")))		//Handle http request
	http.HandleFunc("/guess",templateHandler)				//handle requests for templates
	
	http.ListenAndServe(":8080", nil)						//Listen and report from port 8080

}

//strconv.Itoa()