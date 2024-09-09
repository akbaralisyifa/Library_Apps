package main

import (
	"library/config"
	"library/internal/features/books"
	bkhand "library/internal/features/books/handlers"
	bkrepo "library/internal/features/books/repository"
	bksrv "library/internal/features/books/service"
	"library/internal/features/categories"
	chand "library/internal/features/categories/handlers"
	crepo "library/internal/features/categories/repository"
	csrv "library/internal/features/categories/service"
	"library/internal/features/recomendation"
	rhand "library/internal/features/recomendation/handlers"
	rrepo "library/internal/features/recomendation/repository"
	rsrv "library/internal/features/recomendation/service"
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
	um := urepo.NewUserModels(db)
	us := service.NewUserServices(um, jw, pw)
	bm := bkrepo.NewBookModels(db)
	bs := bksrv.NewBookServices(bm, us)
	bh := bkhand.NewBookHandler(bs)

	return bh
}

func InitCategoryRoter(db *gorm.DB) categories.Handler{
	cm := crepo.NewCategoryModels(db)
	cs := csrv.NewCategoryServices(cm)
	ch := chand.NewCategoryHandler(cs)

	return ch
}

func InitRecommendRouter(db *gorm.DB) recomendation.Handler{
	rm := rrepo.NewRecommendModels(db)
	rs := rsrv.NewRecommendServices(rm)
	rh := rhand.NewRecommnedHandler(rs)

	return rh
}

func main() {
	e := echo.New()
	setup := config.ImportSetting()
	connect, _ := config.ConnectDB(&setup)

	connect.AutoMigrate(&urepo.Users{})
	connect.AutoMigrate(&bkrepo.Books{})
	connect.AutoMigrate(&crepo.Categories{})
	connect.AutoMigrate(&rrepo.Recomendation{})

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	ur := InitUserRouter(connect)
	br := InitBookRouter(connect)
	cr := InitCategoryRoter(connect)
	rr := InitRecommendRouter(connect)
	routes.InitRouter(e, ur, br, cr, rr)


	e.Logger.Fatal(e.Start(":3333"))
}