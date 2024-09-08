package main

import (
	"library/config"
	bkrepo "library/internal/features/books/repository"
	brepo "library/internal/features/borrowed/repository"
	crepo "library/internal/features/categories/repository"
	rrepo "library/internal/features/recomendation/repository"
	urepo "library/internal/features/users/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	setup := config.ImportSetting()
	connect, _ := config.ConnectDB(&setup)

	connect.AutoMigrate(&urepo.Users{}, &bkrepo.Books{}, &brepo.Borrowed{}, &crepo.Categories{}, &rrepo.Recomendation{})

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
}