// ...existing code...
package repository

import (
	"errors"
	"strconv"

	"github.com/BerkeCaganDemir/book-app-backend/internal/models"
)

type BookRepository struct {
	Store *JSONStore
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	if err := r.Store.Read(&books); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) FindByID(id string) (models.Book, error) {
	parsedID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.Book{}, errors.New("invalid id")
	}

	books, err := r.GetAll()
	if err != nil {
		return models.Book{}, err
	}
	for _, b := range books {
		if b.ID == parsedID {
			return b, nil
		}
	}
	return models.Book{}, errors.New("book not found")
}

func (r *BookRepository) SaveAll(books []models.Book) error {
	return r.Store.Write(books)
}

func (r *BookRepository) Add(book models.Book) error {
	books, err := r.GetAll()
	if err != nil {
		return err
	}
	books = append(books, book)
	return r.SaveAll(books)
}

func (r *BookRepository) Update(updatedBook models.Book) error {
	books, err := r.GetAll()
	if err != nil {
		return err
	}
	for i, b := range books {
		if b.ID == updatedBook.ID {
			books[i] = updatedBook
			return r.SaveAll(books)
		}
	}
	return errors.New("book not found")
}

func (r *BookRepository) Delete(id string) error {
	parsedID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errors.New("invalid id")
	}

	books, err := r.GetAll()
	if err != nil {
		return err
	}
	newList := make([]models.Book, 0, len(books))
	found := false

	for _, b := range books {
		if b.ID == parsedID {
			found = true
			continue
		}
		newList = append(newList, b)
	}
	if !found {
		return errors.New("book not found")
	}
	return r.SaveAll(newList)
}

// ...existing code...
