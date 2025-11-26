package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title   string `json:"title" gorm:"type:varchar(100);not null"`
	GenreID uint   `json:"genre_id"`
}
