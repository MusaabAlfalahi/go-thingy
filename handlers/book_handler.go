package handlers

import (
	"net/http"
	"github.com/MusaabAlfalahi/go-thingy/dto"
	"github.com/MusaabAlfalahi/go-thingy/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type BookHandler struct {
	Repo repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{Repo: repo}
}

func (b *BookHandler) CreateBook(c echo.Context) error {
	book := new(dto.Book)
	if err := c.Bind(&book); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validateBook(book); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := b.Repo.CreateBook(book)
	if err != nil {
		log.Error("Failed to create book: ", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create book")
	}

	return c.JSON(http.StatusCreated, book)
}

func (b *BookHandler) GetAllBooks(c echo.Context) error {
	books, err := b.Repo.GetAllBooks()
	if err != nil {
		log.Error("Failed to get books: ", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get books")
	}

	return c.JSON(http.StatusOK, books)
}

func (b *BookHandler) GetBook(c echo.Context) error {
	id, err := getIntId(c)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	book, err := b.Repo.GetBookByID(id)
	if err != nil {
		log.Error("Failed to get book: ", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get book")
	}

	return c.JSON(http.StatusOK, book)
}

func (b *BookHandler) UpdateBook(c echo.Context) error {
	id, err := getIntId(c)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	book := new(dto.Book)
	if err := c.Bind(&book); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validateBook(book); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updatedBook, err := b.Repo.UpdateBook(id, book)
	if err != nil {
		log.Error("Failed to update book: ", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update book")
	}

	return c.JSON(http.StatusOK, updatedBook)
}

func (b *BookHandler) DeleteBook(c echo.Context) error {
	id, err := getIntId(c)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = b.Repo.DeleteBook(id)
	if err != nil {
		log.Error("Failed to delete book: ", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to delete book")
	}

	return c.NoContent(http.StatusNoContent)
}
