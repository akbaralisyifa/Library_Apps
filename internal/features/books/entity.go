package books

import "github.com/labstack/echo/v4"

type Books struct {
	ID            uint
	CategoryID    uint
	Title         string
	Author        string
	PublishedYear string
}

type User struct {
    ID   uint
    Role string
}

type Handler interface {
	AddBook() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
	DeleteBook() echo.HandlerFunc
	GetAllBook() echo.HandlerFunc
	GetBookById() echo.HandlerFunc
}

type Query interface {
	AddBook(newBook Books) error
	GetAllBook(title string) ([]Books, error)
	GetBook(bookID uint) (Books, error)
	UpdateBook(bookID uint, updateBook Books) error
	DeleteBook(bookID uint) error
}

type Service interface {
	AddBook(userID uint, newBook Books) error
	GetAllBook(title string) ([]Books, error)
	GetBook(bookID uint) (Books, error)
	UpdateBook(userID, bookID uint, updateBook Books) error
	DeleteBook(userID, bookID uint) error
}