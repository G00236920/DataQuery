package main

import (
	"net/http"
	"text/template"
)

type Conversation struct {
	User     string
	Computer string
}

//With this Project we had the Oppertunity to use Boostrap for the implementation of the GUI in html format
//I decided to Design my own instead

func main() {

	//Add to list of messages between user and Eliza

	http.Handle("/", http.FileServer(http.Dir("./"))) //Handle http request
	http.HandleFunc("/chat", templateHandler)         //handle requests for templates
	http.ListenAndServe(":8080", nil)                 //Listen and report from port 8080

}

func templateHandler(w http.ResponseWriter, r *http.Request) { //Handle Http requests

	var cookie, _ = r.Cookie("myList") //Variables for cookie and its errors

	var usersAnswer = r.FormValue("userinput")

	type convoList []Conversation

	mylist := convoList{}

	mylist = append(mylist, Conversation{Computer: "New", User: usersAnswer})

	cookie = &http.Cookie{ //Cookie Response

		Name:  "myList", //Name of Cookie
		Value: "myList", //Value Cookie Holds

	}

	http.SetCookie(w, cookie) //Set the Cookie

	t, _ := template.ParseFiles("chat/eliza.html") //Parse the template File

	run(t, w, mylist)
}

func run(t *template.Template, w http.ResponseWriter, myList []Conversation) {

	t.Execute(w, myList) // Execute the Tmpl file
}
