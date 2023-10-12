package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"api/delivery"
	"api/repository"
	"api/usecase"

	_ "github.com/lib/pq"
)

type Config struct {
	PGUser     string `json:"PG_USER"`
	PGPassword string `json:"PG_PASSWORD"`
	PGDBName   string `json:"PG_DBNAME"`
}

func loadConfig() (*Config, error) {
	configFile, err := os.Open("../config/config.json")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	// Load configuration
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Failed to load configuration:", err)
		return
	}

	// Construct the PostgreSQL connection string
	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.PGUser, config.PGPassword, config.PGDBName)

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", connString)
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
