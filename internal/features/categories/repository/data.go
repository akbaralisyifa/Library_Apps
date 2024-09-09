package repository

import (
	bkrepo "library/internal/features/books/repository"
	"library/internal/features/categories"

	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name string
	Books bkrepo.Books `gorm:"foreignKey:CategoryID"`
}

func ToCategoryQuery(input categories.Categories) Categories{
	return Categories{
		Name: input.Name,
	}
}

func (ct *Categories) ToCategoryEntity() categories.Categories {
	return categories.Categories{
		ID:   ct.ID,
		Name: ct.Name,
	}
}