package main

import (
	"api/delivery"
	"api/repository"
	"api/usecase"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Load configuration from environment variables using Viper
	viper.SetConfigFile("../.env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Failed to read .env file:", err)
		return
	}
	viper.AutomaticEnv() // Automatically read from environment variables

	// Retrieve PostgreSQL configuration from environment variables
	dbUser := viper.GetString("PG_USER")
	dbPassword := viper.GetString("PG_PASSWORD")
	dbName := viper.GetString("PG_DBNAME")
	sslMode := viper.GetString("PG_SSLMODE")

	fmt.Println("SSL Mode:", os.Getenv("PG_SSLMODE"))

	// Construct the connection string
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbUser, dbPassword, dbName, sslMode)

	// Connect to PostgreSQL using configuration
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Failed to connect to PostgreSQL:", err)
		return
	}
	defer db.Close()

	// Try to ping the database to check the connection
	if err = db.Ping(); err != nil {
		fmt.Println("Failed to ping the database:", err)
		return
	}

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
