package repository

import (
	"errors"
	"log"
	"os"

	"github.com/gitfuf/bookservice/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName = "bookstore"
const colName = "books"

type MongoRepo struct {
	session *mgo.Session
}

func InitMongoRepo() (*MongoRepo, error) {
	ret := &MongoRepo{}
	mgoSrv := os.Getenv("MONGO_SERVER")
	if len(mgoSrv) == 0 {
		mgoSrv = "localhost"
	}
	mgoS, err := mgo.Dial(mgoSrv)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Println("Connection established to mongo server:", mgoSrv)
	ret.session = mgoS

	urlcollection := mgoS.DB(dbName).C(colName)
	if urlcollection == nil {
		err = errors.New("Collection could not be created, maybe need to create it manually")
	}
	// этот код нужен для добавления уникального индекса бд.
	index := mgo.Index{
		Key:      []string{"$text:isbn"},
		Unique:   true,
		DropDups: true,
	}
	urlcollection.EnsureIndex(index)

	return ret, err
}

func CloseMongoRepo(mongo *MongoRepo) {
	mongo.session.Close()
}

func (mr *MongoRepo) AddBook(book *models.Book) error {

	tempSession := mr.session.Copy()
	defer tempSession.Close()

	collection := tempSession.DB(dbName).C(colName)
	err := collection.Insert(book)
	if err != nil {
		log.Println("mongorepo:AddBook err =", err)
		if mgo.IsDup(err) {
			log.Println("Book with this ISBN already exists")
			return errors.New("Book with this ISBN already exists")
		}
	}
	return err
}

func (mr *MongoRepo) GetBook(isbn string) (*models.Book, error) {
	tempSession := mr.session.Copy()
	defer tempSession.Close()

	ret := models.Book{}
	err := tempSession.DB(dbName).C(colName).Find(bson.M{"isbn": isbn}).One(&ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (mr *MongoRepo) DeleteBook(isbn string) error {
	tempSession := mr.session.Copy()
	defer tempSession.Close()

	err := tempSession.DB(dbName).C(colName).Remove(bson.M{"isbn": isbn})
	return err
}

func (mr *MongoRepo) UpdateBook(isbn string, book *models.Book) error {
	tempSession := mr.session.Copy()
	defer tempSession.Close()
	err := tempSession.DB(dbName).C(colName).Update(bson.M{"isbn": isbn}, &book)
	if err != nil {
		switch err {
		default:
			return err
		case mgo.ErrNotFound:
			return errors.New("Not found")
		}
	}
	return nil
}

func (mr *MongoRepo) AllBooks() ([]models.Book, int, error) {
	tempSession := mr.session.Copy()
	defer tempSession.Close()

	var books []models.Book
	coll := tempSession.DB(dbName).C(colName)

	err := coll.Find(bson.M{}).All(&books)
	return books, len(books), err
}

func (mr *MongoRepo) Books(start uint64, count int64) ([]models.Book, int, error) {
	tempSession := mr.session.Copy()
	defer tempSession.Close()

	var books []models.Book
	coll := tempSession.DB(dbName).C(colName)

	err := coll.Find(bson.M{}).Skip(int(start)).Limit(int(count)).All(&books)
	return books, len(books), err
}
