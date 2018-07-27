package main

import (
	"log"

	"flag"
	"net/http"
	"os"

	"github.com/gitfuf/bookservice/controllers"
	"github.com/gitfuf/bookservice/repository"
	"github.com/gorilla/mux"
)

func main() {
	file, err := os.OpenFile("db.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file.Sync()
		file.Close()
	}()
	log.SetOutput(file)

	dbType := flag.String("db", "mongodb", "declare what type of repository use: mongodb, redis")
	flag.Parse()
	var bc *controllers.BookController
	switch *dbType {
	case "mongodb":
		mgo, err := ConnectToMongoDB()
		if err != nil {
			log.Fatal(err)
		}
		defer repository.CloseMongoRepo(mgo)
		bc = controllers.NewBookController(mgo)

	case "redis":
		rds, err := ConnectToRedis()
		if err != nil {
			log.Fatal(err)
		}
		defer repository.CloseRedisRepo(rds)
		bc = controllers.NewBookController(rds)

	}

	var port string
	port = os.Getenv("HTTP_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	router := mux.NewRouter()
	router.HandleFunc("/book", bc.AddBook).Methods("POST")
	router.HandleFunc("/books", bc.AllBooks).Methods("GET")
	router.HandleFunc("/books/{start:[0-9]+}/{count:[0-9]+}", bc.Books).Methods("GET")
	router.HandleFunc("/book/{isbn:[0-9]+}", bc.GetBook).Methods("GET")
	router.HandleFunc("/book/{isbn:[0-9]+}", bc.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{isbn:[0-9]+}", bc.UpdateBook).Methods("PUT")

	http.ListenAndServe(":"+port, router)

}

func ConnectToMongoDB() (*repository.MongoRepo, error) {
	//init mongo session
	mgo, err := repository.InitMongoRepo()
	if err != nil {
		log.Println("can't connect to mongodb:", err)
		return nil, err
	}
	return mgo, nil

}

func ConnectToRedis() (*repository.RedisRepo, error) {
	rds, err := repository.InitRedisRepo()
	if err != nil {
		log.Println("can't connect to redis:", err)
		return nil, err
	}
	return rds, nil
}
