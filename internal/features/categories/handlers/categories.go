package handlers

import (
	"library/internal/features/categories"
	"library/internal/helpers"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	srv categories.Service
}

func NewCategoryHandler(s categories.Service) categories.Handler {
	return &CategoryHandler{
		srv: s,
	}
}

func (ch *CategoryHandler) AddCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input categoryInput

		err := c.Bind(&input);
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}

		err = ch.srv.AddCategory(ToModelCategory(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "Category successfully created", nil))
	}
}

func (ch *CategoryHandler) GetAllCategory() echo.HandlerFunc {
	return func(c echo.Context) error {

		result, err := ch.srv.GetAllCategory()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(404, helpers.ResponseFormat(http.StatusBadRequest,"category not found", nil))
			}
			return c.JSON(500, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return	c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success get data", ToCategoryResponse(result)))
	
	}
}

func (ch *CategoryHandler) UpdateCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		categoryID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Invalid request parameters", nil))
		}
		var input categoryInput
		err = c.Bind(&input)
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}

		err = ch.srv.UpdateCategory(uint(categoryID), ToModelCategory(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "category successfully updated", nil))
	}
}

func (ch *CategoryHandler) DeleteCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		categoryID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Invalid request parameters", nil))
		}

		err = ch.srv.DeleteCategory(uint(categoryID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "category successfully deleted", nil))
	
	}
}