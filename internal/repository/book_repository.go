package repository

import (
	"book-store/internal/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book models.Book) error
	GetAllBooks() ([]models.Book, error)
	Update(book *models.Book) error
	Delete(id uint) error
}

type gormBookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &gormBookRepository{db: db}
}

func (r *gormBookRepository) Create(book models.Book) error {
	return r.db.Create(book).Error
}

func (r *gormBookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book

	err := r.db.Find(books).Error

	if err != nil {
		return nil, err
	}

	return books, err
}

func (r *gormBookRepository) Update(book *models.Book) error {
	return r.db.Updates(&book).Error
}

func (r *gormBookRepository) Delete(id uint) error {
	return r.db.Delete(id).Error
}
