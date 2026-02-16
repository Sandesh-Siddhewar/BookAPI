package handler

import (
	"log"
	"net/http"

	"github.com/Sandesh-Siddhewar/eBook/BOOKAPI/internal/database"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []database.Book
	err := database.DB.Find(&books).Error
	if err != nil {
		log.Printf("Error fetching books: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	var book database.Book

	id := c.Param("id")
	err := database.DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		log.Printf("Error fetching book by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book database.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := database.DB.Create(&book).Error; err != nil {
		log.Printf("Error creating book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book database.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := database.DB.Where("id = ?", id).Updates(&book).Error; err != nil {
		log.Printf("Error updating book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).Delete(&database.Book{}).Error; err != nil {
		log.Printf("Error deleting book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
