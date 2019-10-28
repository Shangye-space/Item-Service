package main

import (
	"log"
	"net/http"

	items "github.com/Shangye-space/Item-Service/src/api/Items"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)

	http.Handle("/", r)
	r.HandleFunc("/api/items/", items.GetItems)
	log.Fatal(http.ListenAndServe(":3348", r))

}
