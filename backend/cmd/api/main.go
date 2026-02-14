package main

import (
	"log"
	"net/http"

	"github.com/BerkeCaganDemir/book-app-backend/internal/handlers"
	"github.com/BerkeCaganDemir/book-app-backend/internal/repository"
	"github.com/BerkeCaganDemir/book-app-backend/internal/services"

	gorillaHandlers "github.com/gorilla/handlers"
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

	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	router.HandleFunc("/books", bookHandler.GetAll).Methods("GET")
	router.HandleFunc("/books", bookHandler.Create).Methods("POST")
	router.HandleFunc("/books/{id}", bookHandler.Update).Methods("PUT")
	router.HandleFunc("/books/{id}", bookHandler.Delete).Methods("DELETE")
	router.HandleFunc("/books/{id}/image", bookHandler.UploadImage).Methods("POST")

	router.PathPrefix("/uploads/").
		Handler(http.StripPrefix("/uploads/",
			http.FileServer(http.Dir("uploads")),
		))

	//
	cors := gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:5173"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type"}),
	)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", cors(router)))
}
