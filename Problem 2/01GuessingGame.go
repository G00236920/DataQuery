package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { //Handler function

	fmt.Fprintf(w, "Guessing Game 3 %s", r.URL.Path[1:]) //Print formatted Text to the body

}

func main() { //Main Function

	http.HandleFunc("/", handler) //Handle function after '/'

	http.ListenAndServe(":8080", nil) //Listen for Port Number 8080

}
