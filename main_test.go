package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MusaabAlfalahi/go-thingy/handlers"
	"github.com/MusaabAlfalahi/go-thingy/models"
	"github.com/MusaabAlfalahi/go-thingy/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	bookJson = `{"title":"One Piece","author":"Oda","published":1997}`
)

func TestCreateBook(t *testing.T) {
	db := models.InitDB("test.db")
	models.MigrateDB(db)
	defer db.Migrator().DropTable(&models.Book{})

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(bookJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	repo := repository.NewGormBookRepository(models.DB)
	h := handlers.NewBookHandler(repo)

	if assert.NoError(t, h.CreateBook(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var createdBook models.Book
		if err := db.First(&createdBook).Error; err != nil {
			t.Fatalf("Failed to retrieve created book: %v", err)
		}
		assert.Equal(t, "One Piece", createdBook.Title)
	}
}

func TestGetBook(t *testing.T) {
	db := models.InitDB("test.db")
	models.MigrateDB(db)
	defer db.Migrator().DropTable(&models.Book{})

	e := echo.New()

	createReq := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(bookJson))
	createReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	createRec := httptest.NewRecorder()
	createContext := e.NewContext(createReq, createRec)
	
	repo := repository.NewGormBookRepository(models.DB)
	h := handlers.NewBookHandler(repo)

	if assert.NoError(t, h.CreateBook(createContext)) {
		assert.Equal(t, http.StatusCreated, createRec.Code)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/books/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, h.GetBook(c)) {
		var retrievedBook models.Book
		if err := db.First(&retrievedBook, 1).Error; err != nil {
			t.Fatalf("Failed to retrieve book: %v", err)
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "One Piece", retrievedBook.Title)
	}
}

func TestUpdateBook(t *testing.T) {
	db := models.InitDB("test.db")
	models.MigrateDB(db)
	defer db.Migrator().DropTable(&models.Book{})

	e := echo.New()

	createReq := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(bookJson))
	createReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	createRec := httptest.NewRecorder()
	createContext := e.NewContext(createReq, createRec)
	
	repo := repository.NewGormBookRepository(models.DB)
	h := handlers.NewBookHandler(repo)

	if assert.NoError(t, h.CreateBook(createContext)) {
		assert.Equal(t, http.StatusCreated, createRec.Code)
	}

	updateReq := httptest.NewRequest(http.MethodPut, "/api/books/1", strings.NewReader(`{"published":1999}`))
	updateReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	updateRec := httptest.NewRecorder()
	updateContext := e.NewContext(updateReq, updateRec)
	updateContext.SetParamNames("id")
	updateContext.SetParamValues("1")

	if assert.NoError(t, h.GetBook(updateContext)) {
		var updatedBook models.Book
		if err := db.First(&updatedBook, 1).Error; err != nil {
			t.Fatalf("Failed to retrieve book: %v", err)
		}
		assert.Equal(t, http.StatusOK, updateRec.Code)
		assert.Equal(t, 1997, updatedBook.Published)
	}
}

func TestDeleteBook(t *testing.T) {
	db := models.InitDB("test.db")
	models.MigrateDB(db)
	defer db.Migrator().DropTable(&models.Book{})

	e := echo.New()

	createReq := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(bookJson))
	createReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	createRec := httptest.NewRecorder()
	createContext := e.NewContext(createReq, createRec)
	
	repo := repository.NewGormBookRepository(models.DB)
	h := handlers.NewBookHandler(repo)

	if assert.NoError(t, h.CreateBook(createContext)) {
		assert.Equal(t, http.StatusCreated, createRec.Code)
	}

	req := httptest.NewRequest(http.MethodDelete, "/api/books/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, h.DeleteBook(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)

		var deletedBook models.Book
		if err := db.First(&deletedBook, 1).Error; err == nil {
			t.Fatalf("Book was not deleted: %v", err)
		}
	}
}