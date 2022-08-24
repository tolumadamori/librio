package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tolumadamori/librio/pkg/routes"
)

func main() {
	//create router
	r := mux.NewRouter()

	//pass router to the router function from package routes
	routes.Router(r)

	http.Handle("/", r)

	//start server
	fmt.Printf("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
