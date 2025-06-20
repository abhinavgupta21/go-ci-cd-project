package routes

import (
	"github.com/abhinavgupta21/go-ci-cd-project/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Config struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func RegisterRoutes(c *Config) {
	bookController := controllers.NewBookController(c.DB)

	c.Router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the Book API!",
		})
	})
	c.Router.POST("/books", bookController.CreateBook)
	c.Router.GET("/books", bookController.GetBooks)
	c.Router.GET("/books/:id", bookController.GetBookByID)
	c.Router.PUT("/books/:id", bookController.UpdateBook)
	c.Router.DELETE("/books/:id", bookController.DeleteBook)
}
