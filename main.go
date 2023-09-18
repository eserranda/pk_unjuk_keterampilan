package main

import (
	"go-prakerja/database"
	"go-prakerja/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	db, err := database.Connect()
	if err != nil {
		e.Logger.Fatal("Tidak dapat terhubung ke database")
	}

	routes.SetupRoutes(e, db)

	e.Start(":8080")
}
