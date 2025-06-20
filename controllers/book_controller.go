package controllers

import (
	"net/http"

	"github.com/abhinavgupta21/go-ci-cd-project/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{DB: db}
}

// getBookByID is a helper to fetch a book by ID and return error if not found
func (bc *BookController) getBookByID(id string) (*models.Book, error) {
	var book models.Book
	err := bc.DB.First(&book, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &book, err
}

func (bc *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		respondJSON(ctx, http.StatusBadRequest, "Validation failed", nil, err.Error())
		return
	}

	if err := bc.DB.Create(&book).Error; err != nil {
		respondJSON(ctx, http.StatusInternalServerError, "Failed to create book", nil, err.Error())
		return
	}

	respondJSON(ctx, http.StatusOK, "Book created successfully", book)
}

func (bc *BookController) GetBooks(ctx *gin.Context) {
	var books []models.Book
	err := bc.DB.Find(&books).Error
	if err != nil {
		respondJSON(ctx, http.StatusInternalServerError, "Failed to fetch books", nil, err.Error())
		return
	}

	respondJSON(ctx, http.StatusOK, "Books retrieved successfully", books)
}

func (bc *BookController) GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := bc.getBookByID(id)
	if err != nil {
		respondJSON(ctx, http.StatusOK, "Book not found", nil)
		return
	}

	respondJSON(ctx, http.StatusOK, "Book retrieved successfully", book)
}

func (bc *BookController) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := bc.getBookByID(id)
	if err != nil {
		respondJSON(ctx, http.StatusOK, "Book not found", nil)
		return
	}

	var input models.Book
	if err := ctx.ShouldBindJSON(&input); err != nil {
		respondJSON(ctx, http.StatusBadRequest, "Invalid input", nil, err.Error())
		return
	}

	book.Title = input.Title
	book.Author = input.Author
	book.PublishedYear = input.PublishedYear
	book.Price = input.Price

	if err := bc.DB.Save(book).Error; err != nil {
		respondJSON(ctx, http.StatusInternalServerError, "Failed to update book", nil, err.Error())
		return
	}

	respondJSON(ctx, http.StatusOK, "Book updated successfully", book)
}

func (bc *BookController) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := bc.getBookByID(id)
	if err != nil {
		respondJSON(ctx, http.StatusOK, "Book not found", nil)
		return
	}

	if err := bc.DB.Delete(book).Error; err != nil {
		respondJSON(ctx, http.StatusInternalServerError, "Failed to delete book", nil, err.Error())
		return
	}

	respondJSON(ctx, http.StatusOK, "Book deleted successfully", nil)
}

// respondJSON is a helper for consistent JSON responses
func respondJSON(ctx *gin.Context, status int, message string, data interface{}, errMsg ...string) {
	resp := gin.H{
		"status":  http.StatusText(status),
		"message": message,
	}
	if len(errMsg) > 0 {
		resp["error"] = errMsg[0]
	}
	if data != nil {
		resp["data"] = data
	}
	ctx.JSON(status, resp)
}
