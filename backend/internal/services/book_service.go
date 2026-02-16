package services

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/BerkeCaganDemir/book-app-backend/internal/models"
	"github.com/BerkeCaganDemir/book-app-backend/internal/repository"
)

type BookService struct {
	Repo *repository.BookRepository
}

func (s *BookService) GetAll() ([]models.Book, error) {
	return s.Repo.GetAll()
}

func (s *BookService) GetByID(id string) (models.Book, error) {
	return s.Repo.FindByID(id)
}

func (s *BookService) Create(book models.Book) (models.Book, error) {
	if book.Title == "" {
		return models.Book{}, errors.New("title cannot be empty")
	}
	if book.Author == "" {
		return models.Book{}, errors.New("author cannot be empty")
	}

	now := time.Now().Unix()

	book.ID = uuid.New().String() //  STRING ID
	book.CreatedAt = now
	book.UpdatedAt = now

	if err := s.Repo.Add(book); err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (s *BookService) Update(id string, updated models.Book) (models.Book, error) {
	existing, err := s.Repo.FindByID(id)
	if err != nil {
		return models.Book{}, err
	}

	existing.Title = updated.Title
	existing.Author = updated.Author
	existing.Notes = updated.Notes
	existing.ImageUrl = updated.ImageUrl
	existing.BuyURL = updated.BuyURL
	existing.UpdatedAt = time.Now().Unix()

	if err := s.Repo.Update(existing); err != nil {
		return models.Book{}, err
	}

	return existing, nil
}

func (s *BookService) Delete(id string) error {
	return s.Repo.Delete(id)
}

func (s *BookService) Search(title string, author string) ([]models.Book, error) {
	books, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []models.Book

	for _, b := range books {
		if title != "" && !containsIgnoreCase(b.Title, title) {
			continue
		}
		if author != "" && !containsIgnoreCase(b.Author, author) {
			continue
		}
		result = append(result, b)
	}

	return result, nil
}

func containsIgnoreCase(source, search string) bool {
	return strings.Contains(
		strings.ToLower(source),
		strings.ToLower(search),
	)
}
