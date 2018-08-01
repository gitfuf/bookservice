# bookservice
Test REST API / gRPC server for book managment. 
Was written in order to practice work with RESP API / gRPC and NoSQL databases like Mongodb and Redis.

Service has two flags:

```
flag.String("db", "redis", "declare what type of repository use: mongodb, redis")
flag.String("bc", "grpc", "declare what type of book controller to use: http, grpc")
```

In order to run server for mongodb as HTTP server:

`go run ./cmd/server/main.go -db=mongodb -bc=http`

In order to run server for mongodb as gRPC server:

`go run ./cmd/server/main.go -db=mongodb -bc=grpc`

### REST API routes:

Used Gorrila Mux library (github.com/gorilla/mux)


Examples with using curl

- Add new book: "/book" POST

`curl -H "Content-Type: application/json" -X POST http://localhost:8080/book -d '{"isbn":"123456789","name":"Mystery II","authors":["Kate A.", "Moo B."],"price":"$20.17"}'`

- Get book info: "/book/{isbn:[0-9]+}" GET

`curl -H "Content-Type: application/json" -X GET http://localhost:8080/book/123456789`

- Update book info: "/book/{isbn:[0-9]+}" PUT

`curl -H "Content-Type: application/json" http://localhost:8080/book/123456789 -X PUT -d '{"isbn":"123456789","name":"Mystery III","authors":["Mark A.", "Many H."],"price":"$20.17"}'`

- Delete book: "/book/{isbn:[0-9]+}" DELETE

`curl -H "Content-Type: application/json" -X DELETE http://localhost:8080/book/123456789`

- Get all books "/books" GET

`curl -H "Content-Type: application/json" -X GET http://localhost:8080/books`

- Get few books "/books/{start:[0-9]+}/{count:[0-9]+}" GET

E.g. get first 10 books

`curl -H "Content-Type: application/json" -X GET http://localhost:8080/books/0/10`


### Table model

For work is used only one simple table 'books'. Model lookes like:
```
type Book struct {
	ISBN    string   `json:"isbn" bson:"isbn"`
	Name    string   `json:"name" bson:"name"`
	Authors []string `json:"authors" bson:"authors"`
	Price   string   `json:"price" bson:"price"`
}
```
