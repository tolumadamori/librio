package routes

import (
	"github.com/gorilla/mux"
	"github.com/tolumadamori/librio/pkg/controllers"
)

// Router function takes a variable r *mux.Router and chains it to the appropriate HandleFunc for the appropriate path.
func Router(r *mux.Router) {
	r.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
}
