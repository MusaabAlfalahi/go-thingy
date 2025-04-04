package main

import (
	"github.com/MusaabAlfalahi/go-thingy/models"
	"github.com/MusaabAlfalahi/go-thingy/routes"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := models.InitDB("books.db")
	models.MigrateDB(db)

	api := e.Group("/api")
	routes.SetUpRoutes(api)

	
	e.Logger.Fatal(e.Start(":8080"))
}