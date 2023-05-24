package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Lis struct {
	object string 
	finish   bool
}




func createItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	log.WithFields(log.Fields{"description": description}).Info("Add new Todo Item. Saving to database.")
	todo := &TodoItemModel{Description: description, Completed: false}
	db.Create(&todo)
	result := db.Last(&todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Value)
}

func StartServer(w http.ResponseWriter, r *http.Request) {
	log.Info("API is ok")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}
func main() {
	

	log.Info("Starting Todo API server")

	r := mux.NewRouter()
	r.HandleFunc("/StartServer", StartServer).Methods("GET")
	r.HandleFunc("/todo", createItem).Methods("POST")

	http.ListenAndServe(":8000", r)
}
