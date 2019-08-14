package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/faozimipa/belajargo/controllers"
	"github.com/faozimipa/belajargo/driver"
	"github.com/faozimipa/belajargo/models"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = driver.ConnectDb()

	r := mux.NewRouter()
	c := controllers.Controller{}

	r.HandleFunc("/books", c.GetBooks(db)).Methods("GET")
	r.HandleFunc("/books/{id}", c.GetBook(db)).Methods("GET")
	r.HandleFunc("/books", c.AddBook(db)).Methods("POST")
	r.HandleFunc("/books", c.UpdateBook(db)).Methods("PUT")
	r.HandleFunc("/books/{id}", c.RemoveBook(db)).Methods("DELETE")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))

}
