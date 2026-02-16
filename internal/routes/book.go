package routes

import (
	"github.com/Sandesh-Siddhewar/eBook/BOOKAPI/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(routes *gin.Engine) {
	bookRoutes := routes.Group("/api/")
	{
		bookRoutes.GET("books", handler.GetBooks)
		bookRoutes.GET("books/:id", handler.GetBookByID)
		bookRoutes.POST("books", handler.CreateBook)
		bookRoutes.PUT("books/:id", handler.UpdateBook)
		bookRoutes.DELETE("books/:id", handler.DeleteBook)
	}
}
