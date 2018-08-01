package main

import (
	"context"
	"fmt"
	"log"

	bookapi "github.com/gitfuf/bookservice/api"
	"google.golang.org/grpc"
)

func main() {
	grpcConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatal("can't connect to grpc")
	}

	defer grpcConn.Close()

	bookCtl := bookapi.NewBookControllerClient(grpcConn)
	book := bookapi.Book{
		Isbn:    "123456",
		Name:    "My book",
		Authors: []string{"Fuf", "Peter P."},
		Price:   "$13.67",
	}
	res, err := bookCtl.AddBook(context.Background(), &book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
