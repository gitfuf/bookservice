package repository

import "github.com/gitfuf/bookservice/models"

//DbRepo is an interface for work with database
type DbRepo interface {
	BookRepo
}

//BookRepo interface for work with Book table
type BookRepo interface {
	AddBook(book *models.Book) error
	GetBook(isbn string) (*models.Book, error)
	UpdateBook(isbn string, book *models.Book) error
	DeleteBook(isbn string) error
	AllBooks() ([]models.Book, int, error)
	Books(start uint64, count int64) ([]models.Book, int, error)
	ClearAll() error
}
