package repository

import (
	"book-store/internal/models"

	"gorm.io/gorm"
)

type GenreRepository interface {
	Create(genre *models.Genre) error
	GetGenres() ([]models.Genre, error)
	Update(id uint, genre models.Genre) error
	Delete(id uint) error
}

type gormGenreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) GenreRepository {
	return &gormGenreRepository{db: db}
}

func (r *gormGenreRepository) Create(genre *models.Genre) error {
	return r.db.Create(genre).Error
}

func (r *gormGenreRepository) GetGenres() ([]models.Genre, error) {
	var genres []models.Genre

	err := r.db.Find(&genres).Error

	if err != nil {
		return nil, err
	}

	return genres, err
}

func (r *gormGenreRepository) Update(id uint, genre models.Genre) error {
	return r.db.Where("id = ?", id).Updates(genre).Error
}

func (r *gormGenreRepository) Delete(id uint) error {
	return r.db.Delete(&models.Genre{}, id).Error
}
