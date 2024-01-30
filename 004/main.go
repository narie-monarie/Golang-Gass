package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	hello := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("frontend/index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The God Father", Director: "Not Known"},
				{Title: "The God Mama", Director: "Not Known Again"},
			},
		}
		templ.Execute(w, films)
	}

	hello_two := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		htmlStr := fmt.Sprintf("<li class=\"mt-2 btn btn-wide\">%s - %s</li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", hello)
	http.HandleFunc("/add-film/", hello_two)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
