package recomendation

import "github.com/labstack/echo/v4"

type Recomendation struct {
	ID     uint
	UserID uint
	BookID uint
	Reason string
}

type Handler interface {
	AddRecommend() echo.HandlerFunc
	GetAllRecommend() echo.HandlerFunc
	UpdateRecommend() echo.HandlerFunc
	DeleteRecommend() echo.HandlerFunc
}

type Query interface {
	AddRecommend(newRecommend Recomendation) error
	GetAllRecommend() ([]Recomendation, error)
	UpdateRecommend(recommendID uint, updateRecommend Recomendation) error
	DeleteRecommend(recommendID uint) error
}

type Service interface {
	AddRecommend(newRecommend Recomendation) error
	GetAllRecommend() ([]Recomendation, error)
	UpdateRecommend(recommendID uint, updateRecommend Recomendation) error
	DeleteRecommend(recommendID uint) error
}