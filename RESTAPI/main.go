package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// books struct

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// auther struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"LastName"`
}

// init books var as a slice book struct
var Books []Book

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params

	// iterate through all bokks to search specif one

	for _, item := range Books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return

		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

// To create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // mock ID
	Books = append(Books, book)
	json.NewEncoder(w).Encode(book)
}

// update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			Books = append(Books, book)
			json.NewEncoder(w).Encode(book)
			return

		}

	}
	json.NewEncoder(w).Encode(Books)
}

// to delete the book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Books)
}

func main() {
	//init router
	r := mux.NewRouter()

	//mock data
	Books = append(Books, Book{ID: "1", Isbn: "35335", Title: "Book number 1", Author: &Author{FirstName: "ashish", LastName: "sharma"}})
	Books = append(Books, Book{ID: "2", Isbn: "353245", Title: "Book number 2", Author: &Author{FirstName: "steve", LastName: "jobs"}})
	Books = append(Books, Book{ID: "3", Isbn: "36435", Title: "Book number 3", Author: &Author{FirstName: "devtron", LastName: "labs"}})

	//route handlers/endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// to run server
	log.Fatal(http.ListenAndServe(":8000", r))
}
