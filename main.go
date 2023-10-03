package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

var lastBookID int

func init() {
	lastBookID = len(books)
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func AddBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	lastBookID++
	newBook.ID = strconv.Itoa(lastBookID)

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, gin.H{"book": newBook, "message": "Book added successfully"})
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func UpdateBookById(c *gin.Context) {
	id := c.Param("id")

	var updatedBook book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	for i, b := range books {
		if b.ID == id {
			updatedBook.ID = b.ID
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, gin.H{"book": updatedBook, "message": "Book updated successfully"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func DeleteBookById(c *gin.Context) {
	id := c.Param("id")

	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})

}

func main() {
	router := gin.Default()
	router.GET("/books", GetBooks)
	router.GET("/books/:id", GetBookById)
	router.POST("/books", AddBook)
	router.PUT("/books/:id", UpdateBookById)
	router.DELETE("/books/:id", DeleteBookById)
	router.Run("localhost:8080")
}
