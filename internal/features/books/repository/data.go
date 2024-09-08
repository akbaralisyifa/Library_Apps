package repository

import (
	brepo "library/internal/features/borrowed/repository"
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
	Borrowed	  []brepo.Borrowed `gorm:"foreignKey:BookID"`
}