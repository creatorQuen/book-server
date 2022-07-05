package app

import (
	"book-server/domain"
)

type BookService interface {
	FindAll() ([]domain.Book, error)
	FindByID(ID int) (domain.Book, error)
	Create(bookRequest domain.Book) (domain.Book, error)
	Update(ID int, bookRequest domain.Book) (domain.Book, error)
	Delete(ID int) (domain.Book, error)
}
