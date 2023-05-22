package main 

import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

+var db, _= gorm.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")


// StartServer starts the server
func StartServer(w http.ResponseWriter, r *http.Request) {
log.Info("API is ok")
w.Header().Set("Content-Type","application/json")
io.WriteString(w, `{"alive": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}


func main () {
	log.Info("Starting TodoList API server")
	r := mux.NewRouter()
	r.HandleFunc("/StartServer", StartServer).Methods("GET")
	http.ListenAndServe(":8000", r)
    
}
