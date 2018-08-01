package repository

import (
	"errors"
	"log"
	"os"

	"github.com/gitfuf/bookservice/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
const dbName = "bookstore"
const colName = "books"
*/
//MongoRepo struct for store mongodb connection
type MongoRepo struct {
	session *mgo.Session
	dbName  string
	colName string
}

//InitMongoRepo is a func for initialize connection with mongodb.
//If connection successfull then connection saved into MongoRepo object
func InitMongoRepo(db, col string) (*MongoRepo, error) {
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
	ret.dbName = db
	ret.colName = col

	urlcollection := mgoS.DB(db).C(col)
	if urlcollection == nil {
		return nil, errors.New("Collection could not be created, maybe need to create it manually")
	}
	// этот код нужен для добавления уникального индекса бд.
	index := mgo.Index{
		Key:      []string{"$text:isbn"},
		Unique:   true,
		DropDups: true,
	}
	err = urlcollection.EnsureIndex(index)

	return ret, err
}

//CloseMongoRepo is a func for close mongodb connection
func CloseMongoRepo(mongo *MongoRepo) {
	mongo.session.Close()
}

//AddBook is a method of MongoRepo. Add book into DB
func (mr *MongoRepo) AddBook(book *models.Book) error {

	tempSession := mr.session.Copy()
	defer tempSession.Close()

	collection := tempSession.DB(mr.dbName).C(mr.colName)
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

//GetBook is a method of MongoRepo. Get book from DB using isbn as a key
func (mr *MongoRepo) GetBook(isbn string) (*models.Book, error) {
	tempSession := mr.session.Copy()
	defer tempSession.Close()

	ret := models.Book{}
	err := tempSession.DB(mr.dbName).C(mr.colName).Find(bson.M{"isbn": isbn}).One(&ret)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, errors.New("No book with this ISBN")
		}
		return nil, err
	}

	return &ret, nil
}

//DeleteBook is a method of MongoRepo. Delete book from DB using isbn as a key
func (mr *MongoRepo) DeleteBook(isbn string) error {
	tempSession := mr.session.Copy()
	defer tempSession.Close()

	err := tempSession.DB(mr.dbName).C(mr.colName).Remove(bson.M{"isbn": isbn})
	if err == mgo.ErrNotFound {
		return errors.New("No book with this ISBN")
	}
	return err
}

//UpdateBook is a method of MongoRepo. Update book entry in the DB using isbn as an old key and book object with new data
func (mr *MongoRepo) UpdateBook(isbn string, book *models.Book) error {
	tempSession := mr.session.Copy()
	defer tempSession.Close()
	err := tempSession.DB(mr.dbName).C(mr.colName).Update(bson.M{"isbn": isbn}, &book)
	if err != nil {
		switch err {
		default:
			return err
		case mgo.ErrNotFound:
			return errors.New("No book with this ISBN")
		}
	}
	return nil
}

//AllBooks return all books from the DB
func (mr *MongoRepo) AllBooks() ([]models.Book, int, error) {
	tempSession := mr.session.Copy()
	defer tempSession.Close()

	var books []models.Book
	coll := tempSession.DB(mr.dbName).C(mr.colName)

	err := coll.Find(bson.M{}).All(&books)
	return books, len(books), err
}

//Books return range of books from the DB using start and count values
func (mr *MongoRepo) Books(start uint64, count int64) ([]models.Book, int, error) {
	tempSession := mr.session.Copy()
	defer tempSession.Close()

	var books []models.Book
	coll := tempSession.DB(mr.dbName).C(mr.colName)

	err := coll.Find(bson.M{}).Skip(int(start)).Limit(int(count)).All(&books)
	return books, len(books), err
}

func (mr *MongoRepo) ClearAll() error {
	tempSession := mr.session.Copy()
	defer tempSession.Close()

	_, err := tempSession.DB(mr.dbName).C(mr.colName).RemoveAll(bson.M{})
	return err
}
