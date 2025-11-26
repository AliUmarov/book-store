package services

import (
	"book-store/internal/models"
	"book-store/internal/repository"
)

type GenreService interface {
	Create(req models.GenreCreateRequest) error
	GetAllGenres() ([]models.Genre, error)
	Update(id uint, req *models.GenreUpdateRequest) (*models.Genre, error)
	Delete(id uint) error
}

type genreService struct {
	genreRep repository.GenreRepository
}

func NewGenreService(genreRep repository.GenreRepository) GenreService {
	return &genreService{genreRep: genreRep}
}

func (s *genreService) Create(req models.GenreCreateRequest) error {
	var genre = models.Genre{
		Name: req.Name,
	}

	return s.genreRep.Create(&genre)
}

func (s *genreService) GetAllGenres() ([]models.Genre, error) {

	genre, err := s.genreRep.GetGenres()

	if err != nil {
		return nil, err
	}

	return genre, err
}

func (s *genreService) Update(id uint, req *models.GenreUpdateRequest) (*models.Genre, error) {
	var updated *models.Genre

	if req.Name != nil {
		updated.Name = *req.Name
	}

	err := s.genreRep.Update(id, updated)

	if err != nil {
		return nil, err
	}

	return updated, err
}

func (s *genreService) Delete(id uint) error {
	return s.genreRep.Delete(id)
}
