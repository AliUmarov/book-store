package transport

import (
	"book-store/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, genreService services.GenreService, bookService services.BookService) {
	genreHandler := NewGenreHandler(genreService)
	bookHandler := NewBookHandler(bookService)

	genreHandler.RegisterRoutes(router)
	bookHandler.RegisterRoutes(router)
}
