package transport

import (
	"book-store/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, genreService services.GenreService) {
	genreHandler := NewGenreHandler(genreService)

	genreHandler.RegisterRoutes(router)
}
