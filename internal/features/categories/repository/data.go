package repository

import (
	"library/internal/features/books/repository"

	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name string
	Books repository.Books `gorm:"foreignKey:CategoryID"`
}