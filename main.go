package main

import (
    "github.com/gorilla/mux"
    a "github.com/tommarler/rest_api/model"
    "fmt"
    "encoding/json"
    "log"
    "net/http"


)

var books []a.Book

func main() {
	router := mux.NewRouter()

	books = append(books,
		a.Book{ID: "1", Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
		a.Book{ID: "2", Title: "Goroutines", Author: "Mr. Goroutine", Year: "2011"},
		a.Book{ID: "3", Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
		a.Book{ID: "4", Title: "Golang concurrency", Author: "Mr. Currency", Year: "2013"},
		a.Book{ID: "5", Title: "Golang good parts", Author: "Mr. Good", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book a.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book a.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book a.Book

	book.ID = params["id"]

	for i, item := range books {
		if item.ID == book.ID {
			books = append(books[:i], books[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(books)
}