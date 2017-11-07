package main

import (
	"net/http"
	"strconv"
	"text/template"
)

type Message struct {
	Heading      string
	UserMessage  string
	ElizaMessage string
	Button       string
}

//With this Project we had the Oppertunity to use Boostrap for the implementation of the GUI in html format
//I decided to Design my own instead

func main() {

	var Messages []Message // an empty list

	for i := 0; i < 3; i++ {

		Messages = append(Messages,
			Message{UserMessage: "I Am the User" + strconv.Itoa(i), ElizaMessage: "I am Eliza" + strconv.Itoa(i)})
		//Add to list of messages between user and Eliza

		//fmt.Println(Messages[i].UserMessage, Messages[i].ElizaMessage)

	}

	http.Handle("/", http.FileServer(http.Dir("./"))) //Handle http request
	http.HandleFunc("/chat", templateHandler)         //handle requests for templates
	http.ListenAndServe(":8080", nil)                 //Listen and report from port 8080

}

func templateHandler(w http.ResponseWriter, r *http.Request) { //Handle Http requests

	t, _ := template.ParseFiles("chat/eliza.html")    //Parse the template File
	t.Execute(w, Message{Heading: "Chat with ELIZA"}) // Execute the Tmpl file

}
