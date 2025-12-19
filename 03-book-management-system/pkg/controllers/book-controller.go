package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/RanitManik/go-projects/03-book-management-system/pkg/models"
	"github.com/RanitManik/go-projects/03-book-management-system/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)

	err := book.CreateBook()
	if err != nil {
		http.Error(w, "Can not create book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	book, err := models.GetBookById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = models.DeleteBookById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updatedBook models.Book
	utils.ParseBody(r, &updatedBook)

	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	book, err := models.UpdateBookById(id, &updatedBook)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
