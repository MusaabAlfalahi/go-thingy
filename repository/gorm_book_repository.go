package repository

import (
	"github.com/MusaabAlfalahi/go-thingy/dto"
	"github.com/MusaabAlfalahi/go-thingy/models"
	"gorm.io/gorm"
)

type GormBookRepository struct {
	DB *gorm.DB
}

func NewGormBookRepository(db *gorm.DB) *GormBookRepository {
	return &GormBookRepository{DB: db}
}

func (r *GormBookRepository) CreateBook(book *dto.Book) error {
	return r.DB.Create(book).Error
}

func (r *GormBookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *GormBookRepository) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *GormBookRepository) UpdateBook(id uint, updatedBook *dto.Book) (*models.Book, error) {
	book, err := r.GetBookByID(id)
	if err != nil {
		return nil, err
	}

	if updatedBook.Title != "" {
		book.Title = updatedBook.Title
	}
	if updatedBook.Author != "" {
		book.Author = updatedBook.Author
	}
	if updatedBook.Published > 0 {
		book.Published = updatedBook.Published
	}

	err = r.DB.Save(&book).Error

	return book, err
}

func (r *GormBookRepository) DeleteBook(id uint) error {
	return r.DB.Delete(&models.Book{}, id).Error
}
