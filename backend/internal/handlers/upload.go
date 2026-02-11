package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type UploadHandler struct {
	ImageURL string `json:"imageUrl"`
}

type UploadResponse struct {
	ImageURL string `json:"imageUrl"`
}

func (h *UploadHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	// parse multipart form (10 MB)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "could not parse multipart form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "could not get uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		http.Error(w, "only jpg, jpeg and png files are allowed", http.StatusBadRequest)
		return
	}

	uploadDir := "./uploads/images"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, "could not create upload directory", http.StatusInternalServerError)
		return
	}

	fileName := uuid.New().String() + ext
	filePath := filepath.Join(uploadDir, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "could not create file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		http.Error(w, "could not save file", http.StatusInternalServerError)
		return
	}

	// convert to forward slashes and ensure leading slash for URL safety
	urlPath := filepath.ToSlash(filePath)
	urlPath = strings.TrimPrefix(urlPath, ".")
	if !strings.HasPrefix(urlPath, "/") {
		urlPath = "/" + urlPath
	}

	resp := UploadResponse{
		ImageURL: urlPath,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
