package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love this %s !", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	// listen on port 8080 on any interface, this function will block until program is terminated
	log.Fatal(http.ListenAndServe(":8080", nil))
}
