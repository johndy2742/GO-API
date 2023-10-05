// usecase/book_usecase.go

package usecase

import (
	"api/models"
	"api/repository"
)

type BookUseCase interface {
	GetAllBooks() ([]models.Book, error)
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
