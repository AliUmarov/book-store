package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model

	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type GenreCreateRequest struct {
	Name string `json:"name"`
}

type GenreUpdateRequest struct {
	Name *string `json:"name"`
}
