package handlers

import (
	"library/internal/features/users"
	"library/internal/helpers"
	"library/internal/utils"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	srv users.Service
}

func NewUserHandler(s users.Service) users.Handler{
	return &UserHandler{
		srv: s,
	}
}


func (uh *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var result RegisterRequest

		err := c.Bind(&result)
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}

		err = uh.srv.Register(ToModelUser(result))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "register success", nil))
	}
}

func (uh *UserHandler) Login() echo.HandlerFunc{
	return func(c echo.Context) error {
		var input LoginRequest

		err := c.Bind(&input)
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}

		value, token, err := uh.srv.Login(input.Email, input.Password)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}


		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "login success", ToLoginResponse(value, token)))
	}
}

func (uh *UserHandler) UpdateUser() echo.HandlerFunc{
	return func(c echo.Context) error {

		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		
		var input RegisterRequest

		err := c.Bind(&input);
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}

		err = uh.srv.UpdateUser(uint(ID), ToModelUser(input))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "update success", nil))
	}
}

func (uh *UserHandler) DeleteUser() echo.HandlerFunc{
	return func(c echo.Context) error {
		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		err := uh.srv.DeleteUser(uint(ID))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "delete success", nil))
	}
}