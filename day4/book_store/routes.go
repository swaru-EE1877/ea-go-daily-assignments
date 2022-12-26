package bookstore

import "github.com/gin-gonic/gin"

func BookStore() *gin.Engine {
	router := gin.Default()

	bookRoutes := router.Group("/books")
	{
		bookRoutes.GET("", getBooks)
		bookRoutes.GET("/:id", getBookDetails)
		bookRoutes.POST("", addBook)
		bookRoutes.DELETE("/:id", deleteBook)
	}
	return router
}
