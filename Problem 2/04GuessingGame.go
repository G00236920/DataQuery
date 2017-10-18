package main

import (													//Imports
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {											//Page Struct to store attributes of the page
	Title string
	Body  []byte
}

func (p *Page) save() error {								// write to a page 
	filename := p.Title + ".html"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {				//Load a page
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)

	if err != nil {											//Catch Nil Pointer
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil				//Return Page and Values
}

func viewHandler(w http.ResponseWriter, r *http.Request) {	//Handle Http requests
	title := r.URL.Path[len("/"):]
	
	if len(title) == 0 {									//Make index the Home page for localhost:8080
		title = "index"	 									//Im sure there is an official way to do this but
	}														//This will do until I know what it is

	p, _ := loadPage(title)									//Page
	fmt.Fprintf(w, "<div>%s</div>", p.Body)					//Print out
}

func main() {
	http.HandleFunc("/", viewHandler)						//Call Http Handler
	http.ListenAndServe(":8080", nil)						//Listen and report from port 8080
}
