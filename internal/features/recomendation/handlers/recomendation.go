package handlers

import (
	"library/internal/features/recomendation"
	"library/internal/helpers"
	"library/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RecommendHandler struct {
	srv recomendation.Service
}

func NewRecommnedHandler(s recomendation.Service) recomendation.Handler {
	return &RecommendHandler{
		srv: s,
	}
}


// Add recommned book
func (rh *RecommendHandler) AddRecommend() echo.HandlerFunc{
	return func(c echo.Context) error {
		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		bookID, err := strconv.Atoi(c.Param("bookID"))
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Invalid request parameters", nil))
		}

		var input RecommnedInput
		err = c.Bind(&input)
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}

		err = rh.srv.AddRecommend(ToModelRecommend(uint(userID), uint(bookID), input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "recommend book successfully created", nil))
	}
}

// get all recommend book
func (rh *RecommendHandler) GetAllRecommend() echo.HandlerFunc{
	return func(c echo.Context) error {

		result, err := rh.srv.GetAllRecommend()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(404, helpers.ResponseFormat(http.StatusBadRequest,"recommend book not found", nil))
			}
			return c.JSON(500, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}
		
		return	c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success get data", ToRecommendResponse(result)))
	}
}

// update recommend book
func (rh *RecommendHandler) UpdateRecommend() echo.HandlerFunc{
	return func(c echo.Context) error {
		recommnedID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Invalid request parameters", nil))
		}

		var input RecommnedInput
		err = c.Bind(&input)
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}
		err = rh.srv.UpdateRecommend(uint(recommnedID), ToModelRecommendUpdate(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "recommend book successfully updated", nil))
	}
}

// delete recommend book
func (rh *RecommendHandler) DeleteRecommend() echo.HandlerFunc{
	return func(c echo.Context) error {
		recommnedID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Invalid request parameters", nil))
		}

		err = rh.srv.DeleteRecommend(uint(recommnedID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "recommend book successfully deleted", nil))
	}
}