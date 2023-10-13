package repository

import (
	"api/models"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgreSQLBookRepository struct {
	db *sql.DB
}

type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
	AddBook(book *models.Book) error
	GetBookById(id string) (*models.Book, error)
	UpdateBook(id string, updatedBook *models.Book) error
	DeleteBook(id string) error
}

// NewPostgreSQLBookRepository creates a new PostgreSQLBookRepository instance.
func NewPostgreSQLBookRepository(db *sql.DB) *PostgreSQLBookRepository {
	return &PostgreSQLBookRepository{db: db}
}

// GetAllBooks retrieves all books from the database.
func (repo *PostgreSQLBookRepository) GetAllBooks() ([]models.Book, error) {
	// Query to retrieve all books
	rows, err := repo.db.Query("SELECT id, title, author, quantity FROM book")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity); err != nil {
			fmt.Println("Error scanning rows:", err)
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating through rows:", err)
		return nil, err
	}

	fmt.Println("Retrieved books:", books)

	return books, nil
}

// AddBook adds a new book to the database.
func (repo *PostgreSQLBookRepository) AddBook(book *models.Book) error {
	_, err := repo.db.Exec("INSERT INTO book (title, author, quantity) VALUES ($1, $2, $3)",
		book.Title, book.Author, book.Quantity)
	if err != nil {
		return err
	}
	return nil
}

// GetBookById retrieves a book by its ID.
func (repo *PostgreSQLBookRepository) GetBookById(id string) (*models.Book, error) {
	var book models.Book
	err := repo.db.QueryRow("SELECT id, title, author, quantity FROM book WHERE id = $1", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
	if err == sql.ErrNoRows {
		return nil, errors.New("book not found")
	} else if err != nil {
		return nil, err
	}

	return &book, nil
}

// UpdateBook updates an existing book in the database.
func (repo *PostgreSQLBookRepository) UpdateBook(id string, updatedBook *models.Book) error {
	_, err := repo.db.Exec("UPDATE book SET title = $1, author = $2, quantity = $3 WHERE id = $4",
		updatedBook.Title, updatedBook.Author, updatedBook.Quantity, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBook deletes a book by its ID.
func (repo *PostgreSQLBookRepository) DeleteBook(id string) error {
	// Check if the book exists
	var count int
	err := repo.db.QueryRow("SELECT COUNT(*) FROM book WHERE id = $1", id).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		// Book with the provided ID doesn't exist
		return errors.New("book not found")
	}

	_, err = repo.db.Exec("DELETE FROM book WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
