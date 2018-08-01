package main

import (
	"log"

	"flag"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"

	bookapi "github.com/gitfuf/bookservice/api"
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

	dbType := flag.String("db", "redis", "declare what type of repository use: mongodb, redis")
	BcType := flag.String("bc", "grpc", "declare what type of book controller to use: http, grpc")
	flag.Parse()

	switch *dbType {
	case "mongodb":
		mgo, err := ConnectToMongoDB("bookstore", "books")
		if err != nil {
			log.Fatal(err)
		}
		defer repository.CloseMongoRepo(mgo)

		if *BcType == "http" {
			bc := controllers.NewHttpBookController(mgo)
			r, port := setupHttpRoutes(bc)
			log.Fatal(http.ListenAndServe(":"+port, r))
		} else {
			bc := controllers.NewGrpcBookController(mgo)
			booksrv, lis := setupGrpcServer(bc)
			booksrv.Serve(lis)
		}

	case "redis":
		rds, err := ConnectToRedis(0)
		if err != nil {
			log.Fatal(err)
		}
		defer repository.CloseRedisRepo(rds)
		if *BcType == "http" {
			bc := controllers.NewHttpBookController(rds)
			r, port := setupHttpRoutes(bc)
			log.Fatal(http.ListenAndServe(":"+port, r))
		} else {
			bc := controllers.NewGrpcBookController(rds)
			booksrv, lis := setupGrpcServer(bc)
			log.Fatal(booksrv.Serve(lis))
		}
	}
}

func setupGrpcServer(bc *controllers.GrpcBookController) (*grpc.Server, net.Listener) {
	port := os.Getenv("GRPC_PORT")
	if len(port) == 0 {
		port = "8081"
	}
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("can't listen port", err)
	}
	srv := grpc.NewServer()
	bookapi.RegisterBookControllerServer(srv, bc)
	log.Println("starting bookservice at " + port)
	return srv, l
}

func setupHttpRoutes(bc *controllers.HttpBookController) (*mux.Router, string) {
	port := os.Getenv("HTTP_PORT")
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
	return router, port
}

//ConnectToMongoDB func returns pointer to MongoRepo struct
func ConnectToMongoDB(db, coll string) (*repository.MongoRepo, error) {
	//init mongo session
	mgo, err := repository.InitMongoRepo(db, coll)
	if err != nil {
		log.Println("can't connect to mongodb:", err)
		return nil, err
	}
	return mgo, nil

}

//ConnectToRedis unc returns pointer to RedisRepo struct
func ConnectToRedis(dbNum int) (*repository.RedisRepo, error) {
	rds, err := repository.InitRedisRepo(dbNum)
	if err != nil {
		log.Println("can't connect to redis:", err)
		return nil, err
	}
	return rds, nil
}
