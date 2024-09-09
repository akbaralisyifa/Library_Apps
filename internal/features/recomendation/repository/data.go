package repository

import (
	"library/internal/features/recomendation"

	"gorm.io/gorm"
)

type Recomendation struct {
	gorm.Model
	UserID uint
	BookID uint
	Reason string
}


func ToRecommendQuery(input recomendation.Recomendation) Recomendation{
	return Recomendation{
		UserID: input.UserID,
		BookID: input.BookID,
		Reason: input.Reason,
	}
}

func (rm *Recomendation) ToRecommendEntity() recomendation.Recomendation {
	return recomendation.Recomendation{
		ID:     rm.ID, 
		UserID: rm.UserID,
		BookID: rm.BookID,
		Reason: rm.Reason,
	}
}
