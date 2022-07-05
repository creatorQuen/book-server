package repository

import "book-server/domain"

type BookRepository interface {
	FindAll() ([]domain.Book, error)
	FindByID(ID int) (domain.Book, error)
	Create(book domain.Book) (domain.Book, error)
	Update(book domain.Book) (domain.Book, error)
	Delete(book domain.Book) (domain.Book, error)
}
