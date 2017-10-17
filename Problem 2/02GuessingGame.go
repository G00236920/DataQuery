package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { //Handler function

	fmt.Fprintf(w, "\n\n<h1>Guessing Game %s </h1>\n\n", r.URL.Path[1:]) //Print formatted Text to the body

}

func main() { //Main Function

	http.HandleFunc("/", handler) //Handle function after '/'

	http.ListenAndServe(":8080", nil) //Listen for Port Number 8080

}
