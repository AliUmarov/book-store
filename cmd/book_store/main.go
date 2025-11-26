package main

import (
	"book-store/internal/config"
	"book-store/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.Connect()

	if err := db.AutoMigrate(&models.Genre{}, &models.Book{}); err != nil {
		log.Fatal("Не удалось выполнить миграции", err)
	}

	r := gin.Default()

	err := r.Run()
	if err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
