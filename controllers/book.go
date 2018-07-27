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

type BookController struct {
	db repository.DbRepo
}

type booksResponse struct {
	Amount int
	Books  []models.Book
}

func NewBookController(repo repository.DbRepo) *BookController {
	return &BookController{db: repo}
}

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

func (bc *BookController) Books(w http.ResponseWriter, r *http.Request) {
	log.Println("Get books")
	vars := mux.Vars(r)
	countS := vars["count"]
	startS := vars["start"]
	start, _ := strconv.Atoi(startS)
	count, _ := strconv.Atoi(countS)
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
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
