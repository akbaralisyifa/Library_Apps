package repository

import (
	bkrepo "library/internal/features/books/repository"

	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name string
	Books bkrepo.Books `gorm:"foreignKey:CategoryID"`
}