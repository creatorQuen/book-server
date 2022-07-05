package repository

import (
	"book-server/domain"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type repository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.Book, error) {
	var books []domain.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) FindByID(ID int) (domain.Book, error) {
	var book domain.Book
	err := r.db.Find(&book, ID).Error
	return book, err
}

func (r *repository) Create(book domain.Book) (domain.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) Update(book domain.Book) (domain.Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *repository) Delete(book domain.Book) (domain.Book, error) {
	err := r.db.Delete(&book).Error
	return book, err
}
