package main

import (
	"fmt"       //import Fmt
	"io/ioutil" //Import ioutil
)

type Page struct {
	Title string
	Body  []byte
}

func main() {

	p1 := &Page{Title: "Guessing Game", Body: []byte("Gueesing Game.")}

	p1.save()

	p2, _ := loadPage("TestPage")

	fmt.Println(string(p2.Body))

}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := ioutil.ReadFile(filename)

	if err != nil {

		return nil, err

	}

	return &Page{Title: title, Body: body}, nil

}

func (p *Page) save() error {

	filename := p.Title + ".txt"

	return ioutil.WriteFile(filename, p.Body, 0600)

}
