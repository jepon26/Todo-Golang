package main

import (
	"html/template"
	"log"
	"net/http"
)

// Declare a global variable to store the parsed HTML template
var tmpl *template.Template

// Define a struct type to represent individual todo items
type List struct {
	object string
	finish bool
}

// Define a struct type to represent the data needed for rendering the web page
type PageInfo struct {
	Title string
	Todos []List
}

// Handler function for the "/list" endpoint
func list(w http.ResponseWriter, r *http.Request) {
	// Create an instance of PageInfo with sample todo data
	data := PageInfo{
		Title: "Todo List",
		Todos: []List{
			{object: "Wash the dishes", finish: true},
			{object: "Do the homework", finish: false},
			{object: "Listen to music", finish: false},
		},
	}
	// Execute the template with the provided data and write the output to the ResponseWriter
	tmpl.Execute(w, data)
}

func main() {
	// Create a new ServeMux, which acts as a router for incoming requests
	mux := http.NewServeMux()

	// Parse the HTML template file and store the parsed template in the tmpl variable
	tmpl = template.Must(template.ParseFiles("template/index.html"))

	// Create a file server to handle static files (e.g., CSS, JavaScript) and register it with the ServeMux
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register the list handler function with the ServeMux to handle requests to the "/list" endpoint
	mux.HandleFunc("/list", list)

	// Start the server and log any errors that occur
	log.Fatal(http.ListenAndServe(":8080", mux))
}
