package delivery

import (
	"net/http"

	"api/usecase"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	BookUseCase usecase.BookUseCase // Changed field name to start with uppercase
}

func (h *HTTPHandler) GetBooks(c *gin.Context) {
	books, err := h.BookUseCase.GetAllBooks() // Changed field access to use uppercase field name
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve books"})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}
