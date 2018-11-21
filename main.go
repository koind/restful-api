package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
)

type Book struct {
	ID 	   int 	  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "Book 1", Author: "Gother 1", Year: "2001"},
		Book{ID: 2, Title: "Book 2", Author: "Gother 2", Year: "2002"},
		Book{ID: 3, Title: "Book 3", Author: "Gother 3", Year: "2003"},
		Book{ID: 4, Title: "Book 4", Author: "Gother 4", Year: "2004"},
		Book{ID: 5, Title: "Book 5", Author: "Gother 5", Year: "2005"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get one book")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add new book")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update a book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove a book")
}