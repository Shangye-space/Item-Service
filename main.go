package main

import (
	"log"
	"net/http"

	category "github.com/Shangye-space/Item-Service/src/api/category"
	item "github.com/Shangye-space/Item-Service/src/api/item"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)

	http.Handle("/", r)
	r.HandleFunc("/api/items", item.Get).Methods("GET")
	r.HandleFunc("/api/item/create", item.Create).Methods("POST")
	r.HandleFunc("/api/item/update/{id}", item.Update).Methods("POST")
	r.HandleFunc("/api/item/delete/{id}", item.Delete).Methods("GET")

	/* Categories */
	r.HandleFunc("/api/categories", category.Get).Methods("GET")
	r.HandleFunc("/api/category/create", category.Create).Methods("POST")
	r.HandleFunc("/api/category/update/{id}", category.Update).Methods("POST")
	r.HandleFunc("/api/category/delete/{id}", category.Delete).Methods("GET")

	log.Fatal(http.ListenAndServe(":3348", r))

}
