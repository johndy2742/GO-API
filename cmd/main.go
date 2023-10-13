package main

import (
	"api/config"
	"api/database"
	"api/delivery"
	"api/repository"
	"api/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config:", err)
		return
	}

	// Construct the connection string
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		config.DBUser, config.DBPassword, config.DBName, config.SSLMode, config.DBHost, config.DBPort)

	// Connect to PostgreSQL
	db, err := database.ConnectDB(connectionString)
	if err != nil {
		fmt.Println("Failed to connect to PostgreSQL:", err)
		return
	}
	defer db.Close()

	fmt.Println("Successfully connected to the database")

	// Initialize the PostgreSQL repository
	bookRepository := repository.NewPostgreSQLBookRepository(db)

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
