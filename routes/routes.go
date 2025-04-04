package routes

import (
	"github.com/MusaabAlfalahi/go-thingy/handlers"
	"github.com/MusaabAlfalahi/go-thingy/models"
	"github.com/MusaabAlfalahi/go-thingy/repository"
	"github.com/labstack/echo/v4"
) 

func SetUpRoutes(g *echo.Group) {
	repo := repository.NewGormBookRepository(models.DB)
	bookHandler := handlers.NewBookHandler(repo)

	g.POST("/books", bookHandler.CreateBook)
	g.GET("/books", bookHandler.GetAllBooks)
	g.GET("/books/:id", bookHandler.GetBook)
	g.PUT("/books/:id", bookHandler.UpdateBook)
	g.DELETE("/books/:id", bookHandler.DeleteBook)
}