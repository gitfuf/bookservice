package repository

import "github.com/gitfuf/bookservice/models"

type DbRepo interface {
	BookRepo
}

type BookRepo interface {
	AddBook(book *models.Book) error
	GetBook(isbn string) (*models.Book, error)
	UpdateBook(isbn string, book *models.Book) error
	DeleteBook(isbn string) error
	AllBooks() ([]models.Book, int, error)
	Books(start uint64, count int64) ([]models.Book, int, error)
}
