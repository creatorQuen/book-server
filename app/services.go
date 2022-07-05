package app

import (
	"book-server/domain"
	"book-server/infrastructure/repository"
)

type service struct {
	repository repository.BookRepository
}

func NewService(repository repository.BookRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]domain.Book, error) {
	book, err := s.repository.FindAll()
	return book, err
}

func (s *service) FindByID(ID int) (domain.Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest domain.Book) (domain.Book, error) {
	price := bookRequest.Price
	rating := bookRequest.Rating

	book := domain.Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       int(price),
		Rating:      int(rating),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest domain.Book) (domain.Book, error) {
	book, err := s.repository.FindByID(ID)

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description

	price := bookRequest.Price
	rating := bookRequest.Rating
	book.Price = int(price)
	book.Rating = int(rating)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (domain.Book, error) {
	book, err := s.repository.FindByID(ID)
	deleteBook, err := s.repository.Delete(book)
	return deleteBook, err
}
