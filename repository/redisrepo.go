package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gitfuf/bookservice/models"
	"github.com/go-redis/redis"
)

const booklist = "books"

//RedisRepo for save redis client connection and use methods to work with redis
type RedisRepo struct {
	client *redis.Client
}

//InitRedisRepo is a func to connect to the redis. If connect was successfull: saved connection into RedisRepo object
func InitRedisRepo(dbNum int) (*RedisRepo, error) {
	ret := &RedisRepo{}
	redisSrv := os.Getenv("REDIS_SERVER")
	if len(redisSrv) == 0 {
		redisSrv = "localhost"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisSrv + ":6379",
		Password: "", // no password set
		DB:       dbNum,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	ret.client = client
	return ret, nil
}

//CloseRedisRepo as a func to close connection with redis
func CloseRedisRepo(repo *RedisRepo) {
	log.Println(repo.client.Close())
}

//AddBook is a RedisRepo method. It is save book into DB using ISBN as key. Also add this ISBN to the booklist
func (r *RedisRepo) AddBook(book *models.Book) error {
	key := book.ISBN
	bs, err := json.Marshal(book)
	if err != nil {
		log.Printf("Unable to marshal book entry into bytes: %s\n", err)
		return err
	}
	if err = r.client.Set(key, bs, 0).Err(); err != nil {
		log.Printf("Unable to store book entry into redis: %s\n", err)
		return err
	}

	if err = r.client.RPush(booklist, key).Err(); err != nil {
		log.Printf("Unable to store book entry into redis list: %s\n", err)
		return err
	}
	log.Printf("AddBook %v is successful \n", book)

	return nil
}

//GetBook is a RedisRepo method. It is get book from DB using ISBN as key
func (r *RedisRepo) GetBook(isbn string) (*models.Book, error) {
	ret, err := r.getBookUsingKey(isbn)
	if err != nil {
		if err == redis.Nil {
			return nil, errors.New("No book with this ISBN")
		}
		return nil, err
	}
	log.Printf("GetBook %v is successful \n", ret)
	return ret, nil
}

//UpdateBook is a RedisRepo method.
//It is update book entry in the DB using isbn as an old key and book object with new data
func (r *RedisRepo) UpdateBook(isbn string, book *models.Book) error {
	bs, err := json.Marshal(book)
	if err != nil {
		log.Printf("Unable to mashal book entry (isbn=%s) into bytes: %s\n", isbn, err)
		return err
	}
	if isbn == book.ISBN {
		if err = r.client.SetXX(isbn, bs, 0).Err(); err != nil {
			log.Printf("Unable to update book entry (isbn=%s) into redis: %s\n", isbn, err)
			return err
		}
		log.Printf("UpdatBook %s is successful \n", isbn)
		return nil
	}
	//Key ISBN was changed
	//remove old key
	err = r.deleteKey(isbn)
	if err != nil {
		return err
	}
	err = r.deleteKeyFromList(isbn, 0)
	if err != nil {
		return err
	}

	//add new one entry
	if err = r.client.Set(book.ISBN, bs, 0).Err(); err != nil {
		log.Printf("Unable to store book entry into redis: %s\n", err)
		return err
	}
	if err = r.client.RPush(booklist, book.ISBN).Err(); err != nil {
		log.Printf("Unable to store book entry into redis list: %s\n", err)
		return err
	}

	return nil
}

//DeleteBook is a RedisRepo method. It is delete book from DB using ISBN as key, also delete from booklist
func (r *RedisRepo) DeleteBook(isbn string) error {
	err := r.deleteKey(isbn)
	if err != nil {
		return err
	}

	err = r.deleteKeyFromList(isbn, 0)
	if err != nil {
		return err
	}

	log.Println("DeleteBook successfull")
	return nil
}

//AllBooks is a RedisRepo method. It is return all books from the DB
func (r *RedisRepo) AllBooks() ([]models.Book, int, error) {
	var (
		books []models.Book
	)

	keys, err := r.client.LRange(booklist, 0, -1).Result()
	if err != nil {
		log.Println("Can't get booklist from db:", err)
		return nil, 0, err
	}

	for _, key := range keys {
		book, err := r.getBookUsingKey(key)
		if err != nil {
			continue
		}
		books = append(books, *book)
	}
	log.Printf("Found %d books", len(books))
	/*for _, bk := range books {
		for _, author := range bk.Authors {
			log.Println("author: ", author)
		}

	}*/
	return books, len(books), nil
}

//Books is a RedisRepo method. It is return range of books from the DB using start and count values
func (r *RedisRepo) Books(start uint64, count int64) ([]models.Book, int, error) {
	var books []models.Book

	//if use "lrange 0 1" returned 2 rows so use "lrange 0 0" for 1
	if start == 0 {
		count--
	}
	keys, err := r.client.LRange(booklist, int64(start), count).Result()
	if err != nil {
		log.Println("Can't get booklist from db:", err)
		return nil, 0, err
	}

	for _, key := range keys {
		book, err := r.getBookUsingKey(key)
		if err != nil {
			continue
		}
		books = append(books, *book)
	}
	log.Printf("Found %d books", len(books))

	return books, len(books), nil
}

func (r *RedisRepo) getBookUsingKey(key string) (*models.Book, error) {
	ret := models.Book{}
	data, err := r.client.Get(key).Bytes()
	if err != nil {
		fmt.Printf("Unable to retrieve book entry from redis: %s \n", err)
		return nil, err
	}
	if err := json.Unmarshal(data, &ret); err != nil {
		log.Printf("Unable to unmarshal book entry: %s \n", err)
		return nil, err
	}
	return &ret, nil
}

func (r *RedisRepo) deleteKeyFromList(key string, flag int) error {
	res, err := r.client.LRem(booklist, 0, key).Result()
	if err != nil {
		log.Printf("Unable delete book (isbn=%s) from booklist err: %v", key, err)
		return err
	}

	if res == 0 {
		return errors.New("No book in the list with this ISBN")
	}

	return nil
}

func (r *RedisRepo) deleteKey(key string) error {
	res, err := r.client.Del(key).Result()
	if err != nil {
		log.Printf("Unable delete book (isbn=%s) err: %v", key, err)
		return err
	}
	if res == 0 {
		return errors.New("No book with this ISBN")
	}
	return nil
}

func (r *RedisRepo) ClearAll() error {
	r.client.FlushDB()
	return nil
}
