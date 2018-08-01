package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	bookapi "github.com/gitfuf/bookservice/api"
	"google.golang.org/grpc"
)

var bookCtl bookapi.BookControllerClient

func TestMain(m *testing.M) {
	grpcConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)

	if err != nil {
		fmt.Println("can't connect to grpc")
		os.Exit(1)
	}

	defer grpcConn.Close()

	bookCtl = bookapi.NewBookControllerClient(grpcConn)

	m.Run()
}

func TestAddBook(t *testing.T) {
	tests := []struct {
		name    string
		data    bookapi.Book
		wantErr bool
	}{
		{
			"add new book",
			bookapi.Book{Isbn: "123456", Name: "My book", Authors: []string{"Fuf", "Peter P."}, Price: "$13.67"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := bookCtl.AddBook(context.Background(), &tt.data)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Got error %v, when expected none", err)
					return
				}
			}

			if tt.wantErr {
				t.Error("Expected err, but have none")
				return
			}

			t.Log("response=", res)
		})
	}
}

func TestGetBook(t *testing.T) {
	tests := []struct {
		name    string
		data    bookapi.ISBN
		wantErr bool
	}{
		{
			"get book",
			bookapi.ISBN{Isbn: "123456"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := bookCtl.GetBook(context.Background(), &tt.data)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Got error %v, when expected none", err)
					return
				}
			}

			if tt.wantErr {
				t.Error("Expected err, but have none")
				return
			}

			t.Log("response=", res)
		})
	}
}

func TestUpdateBook(t *testing.T) {
	ub := &bookapi.Book{Isbn: "123456", Name: "My book II", Authors: []string{"Fuf", "Peter P."}, Price: "$20.67"}
	tests := []struct {
		name    string
		data    bookapi.UpdateBookRequest
		wantErr bool
	}{
		{
			"update book",
			bookapi.UpdateBookRequest{Isbn: "123456", Book: ub},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := bookCtl.UpdateBook(context.Background(), &tt.data)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Got error %v, when expected none", err)
					return
				}
			}

			if tt.wantErr {
				t.Error("Expected err, but have none")
				return
			}

			t.Log("response=", res)
		})
	}
}

func TestAllBooks(t *testing.T) {
	tests := []struct {
		name    string
		data    bookapi.Empty
		wantErr bool
	}{
		{
			"get books",
			bookapi.Empty{},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := bookCtl.AllBooks(context.Background(), &tt.data)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Got error %v, when expected none", err)
					return
				}
			}

			if tt.wantErr {
				t.Error("Expected err, but have none")
				return
			}

			t.Log("response=", res)
		})
	}
}

func TestBooks(t *testing.T) {
	tests := []struct {
		name    string
		data    bookapi.Range
		wantErr bool
	}{
		{
			"get books",
			bookapi.Range{Start: 0, Count: 1},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := bookCtl.Books(context.Background(), &tt.data)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Got error %v, when expected none", err)
					return
				}
			}

			if tt.wantErr {
				t.Error("Expected err, but have none")
				return
			}

			t.Log("response=", res)
		})
	}
}

func TestDeleteBook(t *testing.T) {
	tests := []struct {
		name    string
		data    bookapi.ISBN
		wantErr bool
	}{
		{
			"delete book",
			bookapi.ISBN{Isbn: "123456"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := bookCtl.DeleteBook(context.Background(), &tt.data)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Got error %v, when expected none", err)
					return
				}
			}
			if tt.wantErr {
				t.Error("Expected err, but have none")
				return
			}

			t.Log("response=", res)
		})
	}
}
