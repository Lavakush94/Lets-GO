package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//get a book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r) // extracting the params

	log.Println(params)

	for _, item := range books {
		if item.Id == params["id"] {
			log.Println("matched the conditon")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, book := range books {
		if book.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var updatedBook Book
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.Id = book.Id
			books = append(books, updatedBook)
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, book := range books {
		if book.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// book model
type Book struct {
	Id     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author model
type Author struct {
	FristName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// creating slice of books
var books []Book

func main() {
	//init router
	r := mux.NewRouter()

	// sample book data
	books = append(books, Book{Id: "1", Isbn: "44842", Title: "The one", Author: &Author{FristName: "John the", LastName: "Don"}})
	books = append(books, Book{Id: "2", Isbn: "54842", Title: "The Two", Author: &Author{FristName: "Motu the", LastName: "Master"}})
	books = append(books, Book{Id: "3", Isbn: "34842", Title: "The Three", Author: &Author{FristName: "Patalu the", LastName: "Geek"}})
	books = append(books, Book{Id: "4", Isbn: "24842", Title: "The Four", Author: &Author{FristName: "Dr. Jhatka the", LastName: "Scientist"}})
	books = append(books, Book{Id: "5", Isbn: "14842", Title: "The Five", Author: &Author{FristName: "Ghaseeta the", LastName: "Tajurba Guy"}})

	// Router handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
