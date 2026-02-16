package repository

import (
	"errors"

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
	books, err := r.GetAll()
	if err != nil {
		return models.Book{}, err
	}

	for _, b := range books {
		if b.ID == id {
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

func (r *BookRepository) Update(updated models.Book) error {
	books, err := r.GetAll()
	if err != nil {
		return err
	}

	for i, b := range books {
		if b.ID == updated.ID {
			books[i] = updated
			return r.SaveAll(books)
		}
	}

	return errors.New("book not found")
}

func (r *BookRepository) Delete(id string) error {
	books, err := r.GetAll()
	if err != nil {
		return err
	}

	newList := make([]models.Book, 0)
	found := false

	for _, b := range books {
		if b.ID == id {
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
