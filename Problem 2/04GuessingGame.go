package main

import (													//Imports
	"fmt"
	"net/http"
	"text/template"
)

type TemplateData struct {
	Message string
}

func templateHandler(w http.ResponseWriter, r *http.Request) {	//Handle Http requests

	t, _ := template.ParseFiles("template/guess.html")											//Parse the template File
	t.Execute(w, TemplateData{Message: "Guess a Number Between 1 and 20"})						// Execute the Tmpl file

}


func main() {

	http.Handle("/", http.FileServer(http.Dir("./")))		//Handle http request
	http.HandleFunc("/guess",templateHandler)				//handle requests for templates
	
	http.ListenAndServe(":8080", nil)						//Listen and report from port 8080

}