package delivery

import (
	"net/http"

	"api/models"
	"api/usecase"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	BookUseCase usecase.BookUseCase
}

func (h *HTTPHandler) GetBooks(c *gin.Context) {
	books, err := h.BookUseCase.GetAllBooks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve books"})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}

func (h *HTTPHandler) AddBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	err := h.BookUseCase.AddBook(&newBook) // Call the use case to add the book
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to add the book"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Book added successfully"})
}

func (h *HTTPHandler) GetBookById(c *gin.Context) {
	id := c.Param("id")

	book, err := h.BookUseCase.GetBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func (h *HTTPHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	// Bind the request body to a Book struct to get the updated book information
	var updatedBook models.Book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	// Update the book using the use case
	book, err := h.BookUseCase.UpdateBook(id, &updatedBook)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book updated successfully", "book": book})
}

func (h *HTTPHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := h.BookUseCase.DeleteBook(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
