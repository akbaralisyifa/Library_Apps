package repository

import "gorm.io/gorm"

type Recomendation struct {
	gorm.Model
	UserID uint
	BookID uint
	Reason string
}