package controllers

import (
	"context"
	"flag"
	"log"
	"reflect"
	"testing"

	bookapi "github.com/gitfuf/bookservice/api"
	"github.com/gitfuf/bookservice/repository"
)

var dbType *string
var bc *GrpcBookController

func init() {
	dbType = flag.String("db", "redis", "declare what type of repository use: mongodb, redis")
	flag.Parse()

}

func TestMain(m *testing.M) {
	//get connect to db
	switch *dbType {
	case "mongodb":
		mgo, err := repository.InitMongoRepo("booktest", "books")
		if err != nil {
			log.Fatal(err)
		}
		defer repository.CloseMongoRepo(mgo)
		bc = NewGrpcBookController(mgo)
	case "redis":
		rds, err := repository.InitRedisRepo(1)
		if err != nil {
			log.Fatal(err)
		}
		defer repository.CloseRedisRepo(rds)
		bc = NewGrpcBookController(rds)
	}

	m.Run()
}

func TestAddBook(t *testing.T) {
	tests := []struct {
		name     string
		data     *bookapi.Book
		wantErr  bool
		err      string
		response *bookapi.SimpleResponse
	}{
		{
			"add new book",
			&bookapi.Book{Isbn: "123456", Name: "My book", Authors: []string{"Fuf", "Peter P."}, Price: "$13.67"},
			false,
			"",
			&bookapi.SimpleResponse{Ok: true, Err: ""},
		},
		{
			"empty book",
			&bookapi.Book{},
			true,
			"Empty data",
			&bookapi.SimpleResponse{Ok: false, Err: "Empty data"},
		},
		{
			"no isbn",
			&bookapi.Book{Name: "My book", Authors: []string{"Fuf", "Peter P."}, Price: "$13.67"},
			true,
			"Empty data",
			&bookapi.SimpleResponse{Ok: false, Err: "Empty data"},
		},
		{
			"no book",
			nil,
			true,
			"No book",
			&bookapi.SimpleResponse{Ok: false, Err: "No book"},
		},
	}

	clear := func() {
		bc.ClearAll()
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clear()
			defer clear()
			res, err := bc.AddBook(context.Background(), tt.data)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Got error %v, when expected none\n", err)
					return
				}
				if tt.wantErr && tt.err != err.Error() {
					t.Errorf("Expected err %s, but got %s\n", tt.err, err.Error())
					return
				}

				if tt.wantErr && !reflect.DeepEqual(res, tt.response) {
					t.Errorf("Got expected error, but response %v should be %v \n", res, tt.response)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("Expected err %s, but have none\n", tt.err)
				return
			}

			//check response
			if !reflect.DeepEqual(res, tt.response) {
				t.Errorf("Wrong response: expected=%v, got=%v", tt.response, res)
				return
			}

			//check really added
			isbn := &bookapi.ISBN{Isbn: tt.data.Isbn}
			b, err := bc.GetBook(context.Background(), isbn)
			if err != nil {
				t.Errorf("Book %v was not correctly added: %v", tt.data, err)
				return
			}
			if !reflect.DeepEqual(b, tt.data) {
				t.Errorf("Wrong get book: expected=%v, got=%v", tt.data, b)
				return
			}
		})
	}
}

func TestGetBook(t *testing.T) {
	tb := &bookapi.Book{Isbn: "123456", Name: "My book", Authors: []string{"Fuf", "Peter P."}, Price: "$13.67"}

	tests := []struct {
		name     string
		data     *bookapi.ISBN
		wantErr  bool
		err      string
		response *bookapi.Book
	}{
		{
			"get book",
			&bookapi.ISBN{Isbn: "123456"},
			false,
			"",
			&bookapi.Book{Isbn: "123456", Name: "My book", Authors: []string{"Fuf", "Peter P."}, Price: "$13.67"},
		},
		{
			"empty isbn",
			&bookapi.ISBN{},
			true,
			"Empty isbn",
			nil,
		},
		{
			"non-existing isbn",
			&bookapi.ISBN{Isbn: "123"},
			true,
			"No book with this ISBN",
			nil,
		},
		{
			"no isbn",
			nil,
			true,
			"No isbn",
			nil,
		},
	}

	clear := func() {
		bc.ClearAll()
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clear()
			defer clear()

			_, err := bc.AddBook(context.Background(), tb)
			if err != nil {
				t.Error("Can't add data to db")
				return
			}

			res, err := bc.GetBook(context.Background(), tt.data)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Got error %v, when expected none\n", err)
					return
				}
				if tt.wantErr && tt.err != err.Error() {
					t.Errorf("Expected err %s, but got %s\n", tt.err, err.Error())
					return
				}

				if tt.wantErr && !reflect.DeepEqual(res, tt.response) {
					t.Errorf("Got expected error, but response %v should be %v \n", res, tt.response)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("Expected err %s, but have none\n", tt.err)
				return
			}

			//check response
			if !reflect.DeepEqual(res, tt.response) {
				t.Errorf("Wrong response: expected=%v, got=%v", tt.response, res)
				return
			}
		})
	}
}

func TestUpdateBook(t *testing.T) {
	tb := &bookapi.Book{Isbn: "123456", Name: "My book", Authors: []string{"Mary F.", "Peter P."}, Price: "$10.67"}
	ub := &bookapi.Book{Isbn: "123456", Name: "My book II", Authors: []string{"Fuf", "Peter P."}, Price: "$20.67"}
	tests := []struct {
		name     string
		data     *bookapi.UpdateBookRequest
		wantErr  bool
		err      string
		response *bookapi.SimpleResponse
	}{
		{
			"update book",
			&bookapi.UpdateBookRequest{Isbn: "123456", Book: ub},
			false,
			"",
			&bookapi.SimpleResponse{Ok: true, Err: ""},
		},
		{
			"update book with isbn",
			&bookapi.UpdateBookRequest{Isbn: "123456", Book: &bookapi.Book{Isbn: "123457", Name: "My book III", Authors: []string{"Fuf", "Peter P."}, Price: "$20.67"}},
			false,
			"",
			&bookapi.SimpleResponse{Ok: true, Err: ""},
		},
		{
			"update book with non-existing isbn",
			&bookapi.UpdateBookRequest{Isbn: "123", Book: ub},
			true,
			"No book with this ISBN",
			&bookapi.SimpleResponse{Ok: false, Err: "No book with this ISBN"},
		},
		{
			"empty isbn",
			&bookapi.UpdateBookRequest{Isbn: "", Book: ub},
			true,
			"Empty isbn",
			&bookapi.SimpleResponse{Ok: false, Err: "Empty isbn"},
		},
		{
			"no update data",
			nil,
			true,
			"No book",
			&bookapi.SimpleResponse{Ok: false, Err: "No book"},
		},
	}

	clear := func() {
		bc.ClearAll()
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clear()
			defer clear()

			_, err := bc.AddBook(context.Background(), tb)
			if err != nil {
				t.Error("Can't add data to db")
				return
			}

			res, err := bc.UpdateBook(context.Background(), tt.data)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Got error %v, when expected none\n", err)
					return
				}
				if tt.wantErr && tt.err != err.Error() {
					t.Errorf("Expected err %s, but got %s\n", tt.err, err.Error())
					return
				}

				if tt.wantErr && !reflect.DeepEqual(res, tt.response) {
					t.Errorf("Got expected error, but response %v should be %v \n", res, tt.response)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("Expected err %s, but have none\n", tt.err)
				return
			}

			//check response
			if !reflect.DeepEqual(res, tt.response) {
				t.Errorf("Wrong response: expected=%v,\n got=%v\n", tt.response, res)
				return
			}
			//check really added
			isbn := &bookapi.ISBN{Isbn: tt.data.Book.Isbn}
			b, err := bc.GetBook(context.Background(), isbn)
			if err != nil {
				t.Errorf("Book %v was not correctly updated: %v\n", tt.data, err)
				return
			}
			if !reflect.DeepEqual(b, tt.data.Book) {
				t.Errorf("Wrong get book: expected=%v,\n got=%v", tt.data, b)
				return
			}

		})
	}
}

func TestDeleteBook(t *testing.T) {
	tb := &bookapi.Book{Isbn: "123456", Name: "My book", Authors: []string{"Mary F.", "Peter P."}, Price: "$10.67"}
	tests := []struct {
		name     string
		data     *bookapi.ISBN
		wantErr  bool
		err      string
		response *bookapi.SimpleResponse
	}{
		{
			"delete book",
			&bookapi.ISBN{Isbn: "123456"},
			false,
			"Wrong data",
			&bookapi.SimpleResponse{Ok: true, Err: ""},
		},
		{
			"non-existing ISBN",
			&bookapi.ISBN{Isbn: "123"},
			true,
			"No book with this ISBN",
			&bookapi.SimpleResponse{Ok: false, Err: "No book with this ISBN"},
		},
		{
			"empty ISBN",
			&bookapi.ISBN{},
			true,
			"Empty isbn",
			&bookapi.SimpleResponse{Ok: false, Err: "Empty isbn"},
		},
		{
			"nil ISBN",
			nil,
			true,
			"No isbn",
			&bookapi.SimpleResponse{Ok: false, Err: "No isbn"},
		},
	}

	clear := func() {
		bc.ClearAll()
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			clear()
			defer clear()

			_, err := bc.AddBook(context.Background(), tb)
			if err != nil {
				t.Error("Can't add data to db")
				return
			}

			res, err := bc.DeleteBook(context.Background(), tt.data)

			if err != nil {
				if !tt.wantErr {
					t.Errorf("Got error %v, when expected none\n", err)
					return
				}
				if tt.wantErr && tt.err != err.Error() {
					t.Errorf("Expected err %s, but got %s\n", tt.err, err.Error())
					return
				}

				if tt.wantErr && !reflect.DeepEqual(res, tt.response) {
					t.Errorf("Got expected error, but response %v should be %v \n", res, tt.response)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("Expected err %s, but have none\n", tt.err)
				return
			}

			//check response
			if !reflect.DeepEqual(res, tt.response) {
				t.Errorf("Wrong response: expected=%v, got=%v", tt.response, res)
				return
			}
		})
	}
}

func TestAllBook(t *testing.T) {
	var books []*bookapi.Book
	b1 := &bookapi.Book{Isbn: "123456", Name: "My book", Authors: []string{"Fuf", "Peter P."}, Price: "$20.67"}
	b2 := &bookapi.Book{Isbn: "123457", Name: "My book II", Authors: []string{"Fly K.", "Peter P."}, Price: "$15.69"}
	books = append(books, b1)
	books = append(books, b2)

	tests := []struct {
		name     string
		data     *bookapi.Empty
		wantErr  bool
		err      string
		response *bookapi.BooksResponse
	}{
		{
			"get books",
			&bookapi.Empty{},
			false,
			"",
			&bookapi.BooksResponse{Amount: 2, Books: books},
		},
	}

	clear := func() {
		bc.ClearAll()
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clear()
			defer clear()

			bc.AddBook(context.Background(), b1)
			bc.AddBook(context.Background(), b2)

			res, err := bc.AllBooks(context.Background(), tt.data)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Got error %v, when expected none\n", err)
					return
				}
				if tt.wantErr && tt.err != err.Error() {
					t.Errorf("Expected err %s, but got %s\n", tt.err, err.Error())
					return
				}

				if tt.wantErr && !reflect.DeepEqual(res, tt.response) {
					t.Errorf("Got expected error, but response %v should be %v \n", res, tt.response)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("Expected err %s, but have none\n", tt.err)
				return
			}

			//check response
			if !reflect.DeepEqual(res, tt.response) {
				t.Errorf("Wrong response: expected=%v, got=%v", tt.response, res)
				return
			}
		})
	}
}

func TestBooks(t *testing.T) {
	var (
		books1 []*bookapi.Book
		books2 []*bookapi.Book
		books3 []*bookapi.Book
	)

	b1 := &bookapi.Book{Isbn: "123456", Name: "My book", Authors: []string{"Fuf", "Peter P."}, Price: "$20.67"}
	b2 := &bookapi.Book{Isbn: "123457", Name: "My book II", Authors: []string{"Fly K.", "Peter P."}, Price: "$15.69"}
	books1 = append(books1, b1)

	books2 = append(books2, b1)
	books2 = append(books2, b2)

	books3 = append(books3, b2)

	tests := []struct {
		name     string
		data     *bookapi.Range
		wantErr  bool
		err      string
		response *bookapi.BooksResponse
	}{
		{
			"get books start=0 count=1",
			&bookapi.Range{Start: 0, Count: 1},
			false,
			"",
			&bookapi.BooksResponse{Amount: 1, Books: books1},
		},
		{
			"get books start=0 count=2 ",
			&bookapi.Range{Start: 0, Count: 2},
			false,
			"",
			&bookapi.BooksResponse{Amount: 2, Books: books2},
		},
		{
			"get books start=1 count=1",
			&bookapi.Range{Start: 1, Count: 1},
			false,
			"",
			&bookapi.BooksResponse{Amount: 1, Books: books3},
		},
		{
			"get books no range",
			nil,
			true,
			"No range data",
			&bookapi.BooksResponse{Amount: 0, Books: nil},
		},
		/*{
			"get books wrong range",
			&bookapi.Range{Start: -1, Count: 1},
			true,
			"Wrong range data",
			&bookapi.BooksResponse{Amount: 0, Books: nil},
		},*/
	}

	clear := func() {
		bc.ClearAll()
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clear()
			defer clear()

			bc.AddBook(context.Background(), b1)
			bc.AddBook(context.Background(), b2)

			res, err := bc.Books(context.Background(), tt.data)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Got error %v, when expected none\n", err)
					return
				}
				if tt.wantErr && tt.err != err.Error() {
					t.Errorf("Expected err %s, but got %s\n", tt.err, err.Error())
					return
				}

				if tt.wantErr && !reflect.DeepEqual(res, tt.response) {
					t.Errorf("Got expected error, but response %v should be %v \n", res, tt.response)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("Expected err %s, but have none\n", tt.err)
				return
			}

			//check response
			if !reflect.DeepEqual(res, tt.response) {
				t.Errorf("Wrong response: expected=%v,\n got=%v\n", tt.response, res)
				return
			}
		})
	}
}
