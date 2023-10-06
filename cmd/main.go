// main.go

package main

import (
	"api/delivery"
	"api/repository"
	"api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the repository
	bookRepository := repository.NewInMemoryBookRepository()

	// Initialize the usecase
	bookUseCase := usecase.NewBookUseCase(bookRepository)

	// Initialize HTTP handler with the usecase
	httpHandler := &delivery.HTTPHandler{
		BookUseCase: bookUseCase,
	}

	// Initialize the Gin router
	router := gin.Default()
	router.GET("/books", httpHandler.GetBooks)
	router.POST("/books", httpHandler.AddBook)
	router.GET("/books/:id", httpHandler.GetBookById)
	router.PUT("/books/:id", httpHandler.UpdateBook)
	router.DELETE("/books/:id", httpHandler.DeleteBook)

	// Run the server
	router.Run("localhost:8080")
}
