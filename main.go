package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type List struct {
	object string
	finish bool
}

type PageInfo struct {
	Title string
	Todos []List
}

func list(w http.ResponseWriter, r *http.Request) {

	data := PageInfo{

		Title: "Todo List",

		Todos: []List{

			{object: "Wash the dishes", finish: true},

			{object: "Do the homework", finish: false},

			{object: "Listen to music", finish: false},
		},
	}
	tmpl.Execute(w, data)
}

func main() {

	mux := http.NewServeMux()

	tmpl = template.Must(template.ParseFiles("templates/mark.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/list", list)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
