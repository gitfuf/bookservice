package controllers

import (
	"context"
	"errors"
	"log"

	bookapi "github.com/gitfuf/bookservice/api"
	"github.com/gitfuf/bookservice/models"
	"github.com/gitfuf/bookservice/repository"
)

//GrpcBookController struct
type GrpcBookController struct {
	db repository.DbRepo
}

//NewBookController is a func which is return new BookController with selected database handler
func NewGrpcBookController(repo repository.DbRepo) *GrpcBookController {
	return &GrpcBookController{db: repo}
}

func (bc *GrpcBookController) AddBook(ctx context.Context, book *bookapi.Book) (*bookapi.SimpleResponse, error) {
	var err error
	if book == nil {
		err = errors.New("No book")
		return formSimpleResponse("AddBook", err), err
	}
	log.Println("AddBook start:", book)
	b := grpcBookToModelBook(book)
	err = checkBook(b)
	if err != nil {
		return formSimpleResponse("AddBook", err), err
	}
	err = bc.db.AddBook(b)

	return formSimpleResponse("AddBook", err), err

}

func (bc *GrpcBookController) GetBook(ctx context.Context, isbn *bookapi.ISBN) (*bookapi.Book, error) {
	var err error
	if isbn == nil {
		err = errors.New("No isbn")
		return nil, err
	}

	if isbn.Isbn == "" {
		err := errors.New("Empty isbn")
		return nil, err
	}

	book, err := bc.db.GetBook(isbn.Isbn)
	if err != nil {
		return nil, err
	}
	ret := modelBookToGrpc(book)
	return ret, err
}

func (bc *GrpcBookController) UpdateBook(ctx context.Context, data *bookapi.UpdateBookRequest) (*bookapi.SimpleResponse, error) {
	var err error
	if data == nil {
		err = errors.New("No book")
		return formSimpleResponse("UpdateBook", err), err
	}
	if data.Isbn == "" {
		err = errors.New("Empty isbn")
		return formSimpleResponse("UpdateBook", err), err
	}
	mb := grpcBookToModelBook(data.Book)
	err = checkBook(mb)
	if err != nil {
		return formSimpleResponse("UpdateBook", err), err
	}
	err = bc.db.UpdateBook(data.Isbn, mb)
	return formSimpleResponse("UpdateBook", err), err
}

func (bc *GrpcBookController) DeleteBook(ctx context.Context, isbn *bookapi.ISBN) (*bookapi.SimpleResponse, error) {
	var err error
	if isbn == nil {
		err = errors.New("No isbn")
		return formSimpleResponse("DeleteBook", err), err
	}
	if isbn.Isbn == "" {
		err = errors.New("Empty isbn")
		return formSimpleResponse("DeleteBook", err), err
	}
	err = bc.db.DeleteBook(isbn.Isbn)
	return formSimpleResponse("DeleteBook", err), err
}

func (bc *GrpcBookController) AllBooks(ctx context.Context, _ *bookapi.Empty) (*bookapi.BooksResponse, error) {
	books, amount, err := bc.db.AllBooks()
	if err != nil {
		return &bookapi.BooksResponse{Amount: 0, Books: nil}, err
	}
	var grpcBooks []*bookapi.Book
	for _, b := range books {
		gb := modelBookToGrpc(&b)
		grpcBooks = append(grpcBooks, gb)
	}
	return &bookapi.BooksResponse{Amount: int32(amount), Books: grpcBooks}, nil
}

func (bc *GrpcBookController) Books(ctx context.Context, limits *bookapi.Range) (*bookapi.BooksResponse, error) {
	if limits == nil {
		err := errors.New("No range data")
		return &bookapi.BooksResponse{Amount: 0, Books: nil}, err
	}
	ret := &bookapi.BooksResponse{}

	books, amount, err := bc.db.Books(uint64(limits.Start), int64(limits.Count))
	if err != nil {
		return &bookapi.BooksResponse{Amount: 0, Books: nil}, err
	}
	var grpcBooks []*bookapi.Book
	for _, b := range books {
		gb := modelBookToGrpc(&b)
		grpcBooks = append(grpcBooks, gb)
	}
	return &bookapi.BooksResponse{Amount: int32(amount), Books: grpcBooks}, nil

	return ret, nil
}

func (bc *GrpcBookController) ClearAll() {
	bc.db.ClearAll()
}

func formSimpleResponse(method string, err error) *bookapi.SimpleResponse {
	var ret bookapi.SimpleResponse
	if err == nil {
		ret.Ok = true
		ret.Err = ""
	} else {
		ret.Ok = false
		ret.Err = err.Error()
	}
	log.Printf("%s simple response = %v\n", method, ret)
	return &ret
}

//check if book entry has necessary fields
func checkBook(book *models.Book) error {
	if book.ISBN == "" || book.Name == "" || book.Price == "" {
		return errors.New("Empty data")
	}
	return nil
}

func grpcBookToModelBook(book *bookapi.Book) *models.Book {
	return &models.Book{
		ISBN:    book.Isbn,
		Name:    book.Name,
		Authors: book.Authors,
		Price:   book.Price,
	}
}

func modelBookToGrpc(book *models.Book) *bookapi.Book {
	return &bookapi.Book{
		Isbn:    book.ISBN,
		Name:    book.Name,
		Authors: book.Authors,
		Price:   book.Price,
	}
}
