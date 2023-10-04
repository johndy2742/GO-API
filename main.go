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
	bookRepository := repository.NewBookRepository()

	// Initialize the usecase
	bookUseCase := usecase.NewBookUseCase(bookRepository)

	// Initialize HTTP handler with the usecase
	httpHandler := &delivery.HTTPHandler{
		BookUseCase: bookUseCase,
	}

	// Initialize the Gin router
	router := gin.Default()
	router.GET("/books", httpHandler.GetBooks)

	// Run the server
	router.Run("localhost:8080")
}
