package repository

import (
	"gorm.io/gorm"

	"perpusbackendgo/model"
)

// BookRepository cuma tugasnya query mentah ke database, nggak ada logic
type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := r.DB.Find(&books).Error
	return books, err
}

func (r *BookRepository) FindByID(id int) (model.Book, error) {
	var book model.Book
	err := r.DB.First(&book, id).Error
	return book, err
}

func (r *BookRepository) Create(book *model.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepository) Update(book *model.Book) error {
	return r.DB.Save(book).Error
}

func (r *BookRepository) Delete(id int) error {
	return r.DB.Delete(&model.Book{}, id).Error
}
