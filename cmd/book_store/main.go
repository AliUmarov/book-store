package main

import (
	"book-store/internal/config"
	"book-store/internal/models"
	"book-store/internal/repository"
	"book-store/internal/services"
	"book-store/internal/transport"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.Connect()

	if err := db.AutoMigrate(&models.Genre{}, &models.Book{}); err != nil {
		log.Fatal("Не удалось выполнить миграции", err)
	}

	genreRep := repository.NewGenreRepository(db)

	genreService := services.NewGenreService(genreRep)

	r := gin.Default()

	transport.RegisterRoutes(r, genreService)

	err := r.Run()
	if err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
