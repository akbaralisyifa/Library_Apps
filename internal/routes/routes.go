package routes

import (
	"library/config"
	"library/internal/features/books"
	"library/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRouter(c *echo.Echo, uh users.Handler, bh books.Handler) {
	
	// jwt Key
	jwtKey := config.ImportSetting().JWTSecrat

	// USER
	c.POST("/register", uh.Register())
	c.POST("/register/:admin", uh.Register())
	c.POST("/login", uh.Login())

	ug:= c.Group("/users")
	ug.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey: []byte(jwtKey),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	));

	ug.PUT("", uh.UpdateUser())
	ug.DELETE("", uh.DeleteUser())

	// BOOK
	c.GET("/books", bh.GetAllBook())
	
	bg := c.Group("/books")
	bg.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey: []byte(jwtKey),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	));
	
	bg.POST("", bh.AddBook())
	bg.PUT("", bh.UpdateBook())
	bg.GET("/:id", bh.GetBookById())
}