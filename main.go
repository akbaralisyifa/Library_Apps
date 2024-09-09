package main

import (
	"library/config"
	"library/internal/features/books"
	bkhand "library/internal/features/books/handlers"
	bkrepo "library/internal/features/books/repository"
	bksrv "library/internal/features/books/service"
	crepo "library/internal/features/categories/repository"
	rrepo "library/internal/features/recomendation/repository"
	"library/internal/features/users"
	"library/internal/features/users/handlers"
	urepo "library/internal/features/users/repository"
	"library/internal/features/users/service"
	"library/internal/routes"
	"library/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitUserRouter(db *gorm.DB) users.Handler {
	jw := utils.NewJwtUtility()
	pw := utils.NewPasswordUtility()
	um := urepo.NewUserModels(db);
	us := service.NewUserServices(um, jw, pw)
	uh := handlers.NewUserHandler(us)

	return uh
}

func InitBookRouter(db *gorm.DB) books.Handler{
	jw := utils.NewJwtUtility()
	pw := utils.NewPasswordUtility()
	um := urepo.NewUserModels(db);
	us := service.NewUserServices(um, jw, pw)
	bm := bkrepo.NewBookModels(db)
	bs := bksrv.NewBookServices(bm, us)
	bh := bkhand.NewBookHandler(bs)

	return bh
}

func main() {
	e := echo.New()
	setup := config.ImportSetting()
	connect, _ := config.ConnectDB(&setup)

	connect.AutoMigrate(&urepo.Users{}, &bkrepo.Books{}, &crepo.Categories{}, &rrepo.Recomendation{})

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	ur := InitUserRouter(connect)
	br := InitBookRouter(connect)
	routes.InitRouter(e, ur, br)


	e.Logger.Fatal(e.Start(":8888"))
}