package services

import (
	"book-store/internal/models"
	"book-store/internal/repository"
)

type BookService interface {
	Create(id uint, req models.BookCreateRequest) error
	GetAllBooks() ([]models.Book, error)
	Update(id uint, req models.BookUpdateRequest) (*models.Book, error)
	Delete(id uint) error
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{bookRepo: bookRepo}
}

func (s *bookService) Create(id uint, req models.BookCreateRequest) error {
	var book = models.Book{
		Title:   req.Title,
		GenreID: id,
	}

	return s.bookRepo.Create(&book)
}

func (h *bookService) GetAllBooks() ([]models.Book, error) {
	books, err := h.bookRepo.GetAllBooks()

	if err != nil {
		return nil, err
	}

	return books, err
}

func (h *bookService) Update(id uint, req models.BookUpdateRequest) (*models.Book, error) {
	var updated models.Book

	if req.Title != nil {
		updated.Title = *req.Title
	}

	err := h.bookRepo.Update(id, updated)

	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (h *bookService) Delete(id uint) error {
	return h.bookRepo.Delete(id)
}
