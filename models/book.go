package models

//Book struct
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

//Response setting respons api
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Book
}
