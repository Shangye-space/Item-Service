package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRouter(router *mux.Router) {
	router.
		Methods("GET").
		Path("/").
		HandlerFunc(postFunction)
}

func postFunction(w http.ResponseWriter, r *http.Request) {
	log.Println("You called a thing!")
	fmt.Fprintf(w, "Welcome to go Service!")

}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
