package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tolumadamori/librio/pkg/config"
	"github.com/tolumadamori/librio/pkg/models"
	"github.com/tolumadamori/librio/pkg/utils"
)

// Wrapper function to check for marshalling errors. Returns a bad request status in the header if there is an error.
func errorChecker(book *models.Book, w http.ResponseWriter) {
	if res, err := json.Marshal(book); err != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// Wrapper function to check for parse errors and return a bad request status in the header if there is an eror.
func parseErrorChecker(a []byte, w http.ResponseWriter) []byte {
	if a != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(a)
		return a
	}
	return nil
}

// Handler func for the create book route. You must include an ID in the request body.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	db.AutoMigrate(&models.Book{})
	newbook := &models.Book{}
	if err := parseErrorChecker(utils.ParseBody(r, newbook), w); err != nil {
		return
	}
	book := newbook.CreateBook(db)
	errorChecker(book, w)
}

// Handler func for the get books route
func GetBooks(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	db.AutoMigrate(&models.Book{})

	books := models.FindBooks(db)
	if res, err := json.Marshal(books); err != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// Hnndler func for the get book by id route.
func GetBook(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	db.AutoMigrate(&models.Book{})

	vars := mux.Vars(r)
	ID, _ := strconv.ParseInt(vars["id"], 0, 0)

	book := models.FindBookByID(db, ID)
	errorChecker(book, w)
}

// Handler func for the delete book route.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	db.AutoMigrate(&models.Book{})
	vars := mux.Vars(r)
	ID, _ := strconv.ParseInt(vars["id"], 0, 0)
	book := models.DeleteBook(db, ID)
	errorChecker(book, w)
}

// Handler func for the Update book route. You must include an ID in the request body.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	db.AutoMigrate(&models.Book{})

	vars := mux.Vars(r)
	ID, _ := strconv.ParseInt(vars["id"], 0, 0)
	updatebook := models.Book{}
	if err := parseErrorChecker(utils.ParseBody(r, &updatebook), w); err != nil {
		return
	}
	book := models.UpdateBook(db, ID, updatebook)
	errorChecker(book, w)
}
