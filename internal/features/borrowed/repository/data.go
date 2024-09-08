package repository

import "gorm.io/gorm"

type Borrowed struct {
	gorm.Model
	UserID uint
	BookID uint
	Status string
}