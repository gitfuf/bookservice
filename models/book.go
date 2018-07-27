package models

type Book struct {
	ISBN    string   `json: "isbn" bson: "isbn"`
	Name    string   `json: "name" bson: "name"`
	Authors []string `json: "authors" bson: "authors"`
	Price   string   `json: "price" bson: "price"`
}
