package main

import (
	"fmt"
	"io/ioutil"
)

// Page data structure represents a wiki page which consists of a title and body(the page content)
type Page struct {
	Title string
	Body  []byte // body is byte instead of string , coz that is the type what io libraries expect
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("the title you are looking for is not present")
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
