package transport

import (
	"book-store/internal/models"
	"book-store/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GenreHandler struct {
	service services.GenreService
}

func NewGenreHandler(service services.GenreService) *GenreHandler {
	return &GenreHandler{service: service}
}

func (h *GenreHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/genres")
	{
		api.POST("/", h.Create)
		api.GET("/", h.GetAll)
		api.PATCH("/:id", h.Update)
		api.DELETE("/:id", h.Delete)
	}
}

func (h *GenreHandler) Create(c *gin.Context) {
	var req models.GenreCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	err := h.service.Create(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, req)
}

func (h *GenreHandler) GetAll(c *gin.Context) {

	genres, err := h.service.GetAllGenres()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, genres)
}

func (h *GenreHandler) Update(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input models.GenreUpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	updated, err := h.service.Update(uint(id), input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *GenreHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid id"})
		return
	}

	err = h.service.Delete(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "genre deleted"})
}
