package main

import (
	"github.com/RIckyBan/webapp-from-scratch/backend/app/handler"
	"github.com/RIckyBan/webapp-from-scratch/backend/app/infra/mysql"
	"github.com/RIckyBan/webapp-from-scratch/backend/app/usecase"
	"github.com/RIckyBan/webapp-from-scratch/backend/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := db.OpenDB()
	if err != nil {
		panic(err)
	}

	ur := mysql.NewUserRepository(db)
	cr := mysql.NewCartRepository(db)

	// usecase
	au := usecase.NewAdminService(ur)
	cu := usecase.NewCartService(cr)

	// handler
	ah := handler.NewAdminHandler(au)
	ch := handler.NewCartHandler(cu)

	// admin
	e.GET("/admin/users", ah.HandleGetAllUsers)

	// cart
	e.GET("/cart/items", ch.HandleGetAllItemsInCart)
	e.POST("/cart/add-item", ch.HandleAddItemToCart)

	// Start server
	e.Logger.Fatal(e.Start(":5001"))
}
