package repository

import (
	"api/models"
	"errors"
	"fmt"
	"strconv"
)

var books = []models.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
	AddBook(book *models.Book) error
	GetBookById(id string) (*models.Book, error)
	UpdateBook(id string, updatedBook *models.Book) error
	DeleteBook(id string) error
}

type inMemoryBookRepository struct {
}

func NewInMemoryBookRepository() *inMemoryBookRepository {
	return &inMemoryBookRepository{}
}

func (repo *inMemoryBookRepository) GetAllBooks() ([]models.Book, error) {
	// Return the package-level books slice
	return books, nil
}

func (repo *inMemoryBookRepository) AddBook(book *models.Book) error {
	// Increment the book ID based on the number of existing books
	book.ID = strconv.Itoa(len(books) + 1)

	// Append the new book to the books slice
	books = append(books, *book)
	return nil
}

func (repo *inMemoryBookRepository) GetBookById(id string) (*models.Book, error) {
	for _, book := range books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, errors.New("book not found")
}

func (repo *inMemoryBookRepository) UpdateBook(id string, updatedBook *models.Book) error {
	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = book.ID
			// Update the book details
			books[i] = *updatedBook
			return nil
		}
	}
	return fmt.Errorf("book with ID %s not found", id)
}

func (repo *inMemoryBookRepository) DeleteBook(id string) error {
	for i, book := range books {
		if book.ID == id {
			// Delete the book from the slice
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
