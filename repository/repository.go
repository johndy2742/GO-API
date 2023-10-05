package repository

import "api/models"

type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
}

type bookRepository struct {
}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (repo *bookRepository) GetAllBooks() ([]models.Book, error) {
	books := []models.Book{
		{ID: "1", Title: "Book 1", Author: "Author 1", Quantity: 10},
		{ID: "2", Title: "Book 2", Author: "Author 2", Quantity: 5},
		{ID: "3", Title: "21212121", Author: "Author 2", Quantity: 5},
	}

	return books, nil
}
