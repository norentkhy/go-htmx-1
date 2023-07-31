package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("hi")

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func h1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "hi", Director: "mofo"},
			{Title: "ho", Director: "mufu"},
		},
	}
	tmpl.Execute(w, films)
}

func h2(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Header.Get("HX-Request"))
	time.Sleep(500 * time.Millisecond)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	htmlStr := fmt.Sprintf("<li>%s - %s</li>", title, director)
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, nil)
}
