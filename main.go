package main

import (
	"log"
	"net/http"

	item "github.com/Shangye-space/Item-Service/src/api/item"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)

	http.Handle("/", r)
	r.HandleFunc("/api/items/", item.Get).Methods("GET")
	r.HandleFunc("/api/item/delete/{id}", item.Delete).Methods("GET")
	r.HandleFunc("/api/item/update/{id}", item.Update).Methods("POST")
	log.Fatal(http.ListenAndServe(":3348", r))

}
