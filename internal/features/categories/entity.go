package categories

import "github.com/labstack/echo/v4"

type Categories struct {
	ID   uint
	Name string
}

type Handler interface {
	AddCategory() echo.HandlerFunc
	GetAllCategory() echo.HandlerFunc
	UpdateCategory() echo.HandlerFunc
	DeleteCategory() echo.HandlerFunc
}

type Query interface {
	AddCategory(newCategory Categories) error
	GetAllCategory() ([]Categories, error)
	UpdateCategory(categoryID uint, updateCategory Categories) error
	DeleteCategory(categoryID uint) error
}

type Service interface {
	AddCategory(userID uint, newCategory Categories) error
	GetAllCategory() ([]Categories, error)
	UpdateCategory(userID, categoryID uint, updateCategory Categories) error
	DeleteCategory(userID, categoryID uint) error
}