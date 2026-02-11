package main

import (
	"log"
	"net/http"

	"github.com/BerkeCaganDemir/book-app-backend/internal/handlers"
	"github.com/BerkeCaganDemir/book-app-backend/internal/repository"
	"github.com/BerkeCaganDemir/book-app-backend/internal/services"
	"github.com/gorilla/mux"
)

func main() {

	store := repository.NewJSONStore("data/books.json")

	bookRepo := &repository.BookRepository{
		Store: store,
	}

	bookService := &services.BookService{
		Repo: bookRepo,
	}

	bookHandler := &handlers.BookHandler{
		Service: bookService,
	}
	uploadHandler := &handlers.UploadHandler{}

	router := mux.NewRouter()
	router.HandleFunc("/books", bookHandler.GetAll).Methods("GET")
	router.HandleFunc("/books", bookHandler.Create).Methods("POST")
	router.HandleFunc("/books/{id}", bookHandler.Update).Methods("PUT")
	router.HandleFunc("/books/{id}", bookHandler.Delete).Methods("DELETE")
	router.HandleFunc("/books/{id}/image", uploadHandler.UploadImage).Methods("POST")

	router.PathPrefix("/uploads/").
		Handler(http.StripPrefix("/uploads/",
			http.FileServer(http.Dir("uploads")),
		))

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
