package handlers

import (
	"errors"
	"strconv"

	"github.com/MusaabAlfalahi/go-thingy/dto"
	"github.com/labstack/echo/v4"
)

func validateBook (b *dto.Book) error {
	if b.Title == "" {
		return errors.New("Title is required")
	}

	if b.Author == "" {
		return errors.New("Author is required")
	}

	if b.Published <= 0 {
		return errors.New("Published year must be greater than 0")
	}

	return nil
}

func getIntId(c echo.Context) (uint, error) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return uint(intId), nil
}