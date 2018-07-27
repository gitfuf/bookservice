package controllers

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gitfuf/bookservice/models"
	"github.com/gitfuf/bookservice/repository"
	"github.com/gorilla/mux"
)

//BookController struct
type BookController struct {
	db repository.DbRepo
}

type booksResponse struct {
	Amount int
	Books  []models.Book
}

//NewBookController is a func which is return new BookController with selected database handler
func NewBookController(repo repository.DbRepo) *BookController {
	return &BookController{db: repo}
}

//AddBook method of BookController processes POST "/book" route
func (bc *BookController) AddBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = bc.db.AddBook(&book)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, book)
}

//GetBook method of BookController processes GET "/book" route
func (bc *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["isbn"]

	book, err := bc.db.GetBook(isbn)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}

//DeleteBook method of BookController processes DELETE "/book/{isbn:[0-9]+}" route
func (bc *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["isbn"]

	err := bc.db.DeleteBook(isbn)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

//UpdateBook method of BookController processes PUT "/book/{isbn:[0-9]+}" route
func (bc *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	vars := mux.Vars(r)
	isbn := vars["isbn"]

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = bc.db.UpdateBook(isbn, &book)
	if err != nil {
		if err.Error() == "Not found" {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

//AllBooks method of BookController processes  GET "/books" route
func (bc *BookController) AllBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("AllBooks begin")
	books, amount, err := bc.db.AllBooks()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}
	log.Println("books amount:", amount)
	resp := &booksResponse{
		Books:  books,
		Amount: amount,
	}
	log.Println(resp)

	respondWithJSON(w, http.StatusOK, resp)
}

//Books method of BookController processes GET "/books/{start:[0-9]+}/{count:[0-9]+}" route
func (bc *BookController) Books(w http.ResponseWriter, r *http.Request) {
	log.Println("Get books")
	vars := mux.Vars(r)
	countS := vars["count"]
	startS := vars["start"]
	start, err := strconv.Atoi(startS)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Bad params")
		return
	}
	count, err := strconv.Atoi(countS)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Bad params")
		return
	}
	log.Printf("count: %d, start: %d", count, start)
	books, amount, err := bc.db.Books(uint64(start), int64(count))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}
	log.Println("books amount:", amount)
	resp := &booksResponse{
		Books:  books,
		Amount: amount,
	}
	respondWithJSON(w, http.StatusOK, resp)

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})

}

func respondWithJSON(w http.ResponseWriter, code int, payload ...interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println("respondWithJSON err: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		log.Println("respondWithJSON err: ", err)
	}
}
