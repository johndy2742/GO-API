package usecase

import (
	"api/models"
	"api/repository"
)

type BookUseCase interface {
	GetAllBooks() ([]models.Book, error)
	AddBook(book *models.Book) error
	GetBookById(id string) (*models.Book, error)
	UpdateBook(id string, updatedBook *models.Book) (*models.Book, error)
	DeleteBook(id string) error
}

type bookUseCase struct {
	bookRepository repository.BookRepository
}

func NewBookUseCase(bookRepository repository.BookRepository) BookUseCase {
	return &bookUseCase{
		bookRepository: bookRepository,
	}
}

func (uc *bookUseCase) GetAllBooks() ([]models.Book, error) {
	return uc.bookRepository.GetAllBooks()
}

func (uc *bookUseCase) AddBook(book *models.Book) error {
	return uc.bookRepository.AddBook(book)
}

func (uc *bookUseCase) GetBookById(id string) (*models.Book, error) {
	return uc.bookRepository.GetBookById(id)
}

func (uc *bookUseCase) UpdateBook(id string, updatedBook *models.Book) (*models.Book, error) {
	err := uc.bookRepository.UpdateBook(id, updatedBook)
	if err != nil {
		return nil, err
	}
	return updatedBook, nil
}

func (uc *bookUseCase) DeleteBook(id string) error {
	return uc.bookRepository.DeleteBook(id)
}
