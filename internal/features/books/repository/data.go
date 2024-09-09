package repository

import (
	"library/internal/features/books"
	rrepo "library/internal/features/recomendation/repository"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	CategoryID	  uint
	Title 		  string
	Author 	 	  string
	PublishedYear string
	Recomendation []rrepo.Recomendation `gorm:"foreignKey:BookID"`
}

func (b *Books) ToBookEntity() books.Books{
	return books.Books{
		ID:            b.ID, 
		CategoryID:    b.CategoryID,
		Title:    	   b.Title,
		Author:        b.Author,
		PublishedYear: b.PublishedYear,
	}
}

func ToBookQuery(input books.Books) Books {
	return Books{
		CategoryID:    input.CategoryID,
		Title:         input.Title,
		Author: 	   input.Author,
		PublishedYear: input.PublishedYear,
	}
}