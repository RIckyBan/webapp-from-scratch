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
	uu := usecase.NewAdminService(ur)
	ah := handler.NewAdminHandler(uu)

	// admin
	e.GET("/admin/users", ah.HandleGetAllUsers)

	// Start server
	e.Logger.Fatal(e.Start(":5001"))
}
