package bookrepository

import (
	"database/sql"
	"log"

	"github.com/faozimipa/belajargo/models"
)

//BookRepository struct
type BookRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//GetBooks repo to get books
func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query("select * from books")
	logFatal(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)
		books = append(books, book)
	}

	return books
}

//GetBook get single book
func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book {
	rows := db.QueryRow("select * from books where id=$1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	logFatal(err)

	return book
}

//AddBook get single book
func (b BookRepository) AddBook(db *sql.DB, book models.Book) int {
	err := db.QueryRow("insert into books(title, author, year) values($1, $2,$3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&book.ID)

	logFatal(err)

	return book.ID
}

//UpdateBook update single book
func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {

	result, err := db.Exec("update books SET title=$1, author=$2, year=$3 where id=$4 RETURNING id;",
		&book.Title, &book.Author, &book.Year, &book.ID)

	rowsUpdated, err := result.RowsAffected()

	logFatal(err)

	return rowsUpdated
}

//RemoveBook delete single book
func (b BookRepository) RemoveBook(db *sql.DB, id int) int64 {

	result, err := db.Exec("delete from books whereid=$1", id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()

	return rowsDeleted
}
