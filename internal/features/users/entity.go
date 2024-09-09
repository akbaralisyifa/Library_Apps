package users

import "github.com/labstack/echo/v4"

type Users struct {
	ID       uint
	Username string
	Email    string
	Password string
	Role     string
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}

type Query interface {
	Register(newUser Users) error
	Login(email string) (Users, error)
	UpdateUser(id uint, updateUser Users) error
	DeleteUser(id uint) error
	GetUser(userID uint) (Users, error)
}

type Service interface {
	Register(newUser Users) error
	Login(email string, password string) (Users, string, error)
	UpdateUser(id uint, updateUser Users) error
	DeleteUser(id uint) error
	GetUser(userID uint) (Users, error)
}