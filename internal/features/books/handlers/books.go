package handlers

import (
	"library/internal/features/books"
	"library/internal/helpers"
	"library/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BookHandler struct {
	srv books.Service
}

func NewBookHandler(s books.Service) books.Handler{
	return &BookHandler{
		srv: s,
	}
}

// Add book
func (bh *BookHandler) AddBook() echo.HandlerFunc{
	return func(c echo.Context) error {
		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		var input BookInput
		
		err := c.Bind(&input);
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}

		err = bh.srv.AddBook(uint(userID), ToModelBook(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "Book successfully created", nil))
	}
}

func (bh *BookHandler) UpdateBook() echo.HandlerFunc{
	return func(c echo.Context) error {
		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		bookID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Invalid request parameters", nil))
		}

		var input BookInput
		err = c.Bind(&input);
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}

		err = bh.srv.UpdateBook(uint(userID), uint(bookID), ToModelBook(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "Book successfully created", nil))
	}
}

// Delete Book
func (bh *BookHandler) DeleteBook() echo.HandlerFunc{
	return func(c echo.Context) error {
		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		bookID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Invalid request parameters", nil))
		}

		err = bh.srv.DeleteBook(uint(userID), uint(bookID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "Book successfully created", nil))
	}
}

// get Book
func (bh *BookHandler) GetAllBook() echo.HandlerFunc{
	return func(c echo.Context) error {
		result, err := bh.srv.GetAllBook()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(404, helpers.ResponseFormat(http.StatusBadRequest,"book not found", nil))
			}
			return c.JSON(500, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return	c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success get data", ToBookResponse(result)))
	}
}

func (bh *BookHandler) GetBookById() echo.HandlerFunc{
	return func(c echo.Context) error {
		bookID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Invalid request parameters", nil))
		}
		result, err := bh.srv.GetBook(uint(bookID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(404, helpers.ResponseFormat(http.StatusBadRequest,"book not found", nil))
			}
			return c.JSON(500, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return	c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success get data", ToBookResponseById(result)))
	
	}
}