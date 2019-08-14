package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/faozimipa/belajargo/models"
	bookrepository "github.com/faozimipa/belajargo/repository/book"
	"github.com/gorilla/mux"
)

// Controller struct
type Controller struct{}

var books []models.Book
var db *sql.DB

//GetBooks get all books
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var book models.Book
		var books []models.Book
		bookRepo := bookrepository.BookRepository{}
		books = bookRepo.GetBooks(db, book, books)
		json.NewEncoder(w).Encode(books)
	}
}

//GetBook get spesific books
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		bookRepo := bookrepository.BookRepository{}
		book = bookRepo.GetBook(db, book, id)

		json.NewEncoder(w).Encode(book)
	}
}

//AddBook tambah buku
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		bookRepo := bookrepository.BookRepository{}
		json.NewDecoder(r.Body).Decode(&book)
		bookID := bookRepo.AddBook(db, book)
		json.NewEncoder(w).Encode(bookID)
	}
}

//UpdateBook update book
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)
		bookRepo := bookrepository.BookRepository{}
		rowsUpdated := bookRepo.UpdateBook(db, book)
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

//RemoveBook Remove books
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		bookRepo := bookrepository.BookRepository{}
		rowsDeleted := bookRepo.RemoveBook(db, id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
