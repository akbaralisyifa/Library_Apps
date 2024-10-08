package utils

import (
	"fmt"
	"library/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtUtilityInteface interface{
	GenerateJwt(id uint) (string, error)
	DecodToken(token *jwt.Token) float64
}

type JwtUtility struct{}

func NewJwtUtility() JwtUtilityInteface {
	return &JwtUtility{}
}

func (ju *JwtUtility) GenerateJwt(id uint) (string, error) {

	jwtKey := config.ImportSetting().JWTSecrat
	data := jwt.MapClaims{}

	data["id"] = id
	data["iat"] = time.Now().Unix()
	data["exp"] = time.Now().Add(time.Hour * 2).Unix() 

	processToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	result, err := processToken.SignedString([]byte(jwtKey))

	if err != nil {
		return "", err
	}

	return result, nil;
}

func (ju *JwtUtility) DecodToken(token *jwt.Token) float64 {
	var result float64

	claim := token.Claims.(jwt.MapClaims)

	for _, val := range claim {
		fmt.Println(val)
	}

	if value, found := claim["id"]; found {
		result = value.(float64)
	}

	return result;
}