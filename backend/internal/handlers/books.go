package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/BerkeCaganDemir/book-app-backend/internal/models"
	"github.com/BerkeCaganDemir/book-app-backend/internal/services"
	"github.com/gorilla/mux"
)

// ...existing code...

type BookHandler struct {
	Service *services.BookService
}

func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	author := r.URL.Query().Get("author")

	var (
		books []models.Book
		err   error
	)
	if title != "" || author != "" {
		books, err = h.Service.Search(title, author)
	} else {
		books, err = h.Service.GetAll()
	}
	if err != nil {
		http.Error(w, "failed to load books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	created, err := h.Service.Create(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(created)
}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "missing book id", http.StatusBadRequest)
		return
	}
	var updated models.Book
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	result, err := h.Service.Update(id, updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "missing book id", http.StatusBadRequest)
		return
	}
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// UploadImage saves an uploaded image file for a book (form field name "image")
func (h *BookHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "missing book id", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB
		http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "missing image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := os.MkdirAll("uploads", 0755); err != nil {
		http.Error(w, "failed to create uploads directory", http.StatusInternalServerError)
		return
	}

	outPath := fmt.Sprintf("uploads/%s.jpg", id)
	out, err := os.Create(outPath)
	if err != nil {
		http.Error(w, "failed to save image", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		http.Error(w, "failed to write image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
