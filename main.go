package main

import (
	"log"
	"net/http"

	category "github.com/Shangye-space/Item-Service/src/api/category"
	image "github.com/Shangye-space/Item-Service/src/api/image"
	item "github.com/Shangye-space/Item-Service/src/api/item"
	iteminfo "github.com/Shangye-space/Item-Service/src/api/itemInfo"
	subCategory "github.com/Shangye-space/Item-Service/src/api/subCategory"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)

	/* Items */
	http.Handle("/", r)
	r.HandleFunc("/api/items", item.GetHandler).Methods("GET")
	r.HandleFunc("/api/items/sub_category/{id}", item.GetBySubCategoryIDHandler).Methods("GET")
	r.HandleFunc("/api/items/category/{id}", item.GetByCategoryIDHandler).Methods("GET")
	r.HandleFunc("/api/item/{id}", item.GetByIDHandler).Methods("GET")
	r.HandleFunc("/api/item/create", item.CreateHandler).Methods("POST")
	r.HandleFunc("/api/item/update/{id}", item.UpdateHandler).Methods("POST")
	r.HandleFunc("/api/item/delete/{id}", item.DeleteHandler).Methods("GET")
	r.HandleFunc("/api/items/count", item.GetCountHandler).Methods("GET")

	/* Item info */
	r.HandleFunc("/api/item_info/{id}", iteminfo.GetByIDHandler).Methods("GET")
	r.HandleFunc("/api/item_info/create/{id}", iteminfo.CreateByIDHandler).Methods("POST")
	r.HandleFunc("/api/item_info/update/{id}", iteminfo.UpdateByIDHandler).Methods("POST")
	r.HandleFunc("/api/item_info/delete/{id}", iteminfo.DeleteByIDHandler).Methods("GET")

	/* Image */
	//r.HandleFunc("/api/image/{id}", image.Get).Methods("GET")
	r.HandleFunc("/api/image/upload", image.Create).Methods("POST")
	//r.HandleFunc("/api/image/update/{id}", image.Update).Methods("POST")
	//r.HandleFunc("/api/image/delete/{id}", image.Delete).Methods("GET")

	/* Categories */
	r.HandleFunc("/api/categories", category.GetHandler).Methods("GET")
	r.HandleFunc("/api/category/{id}", category.GetByIDHandler).Methods("GET")
	r.HandleFunc("/api/category/create", category.CreateHandler).Methods("POST")
	r.HandleFunc("/api/category/update/{id}", category.UpdateHandler).Methods("POST")
	r.HandleFunc("/api/category/delete/{id}", category.DeleteHandler).Methods("GET")

	/* Sub Categories */
	r.HandleFunc("/api/sub_categories", subCategory.GetHandler).Methods("GET")
	r.HandleFunc("/api/sub_category/{id}", subCategory.GetByIDHandler).Methods("GET")
	r.HandleFunc("/api/sub_categories/category/{id}", subCategory.GetSubCategoryByCategoryIDHandler).Methods("GET")
	r.HandleFunc("/api/sub_category/create", subCategory.CreateHandler).Methods("POST")
	r.HandleFunc("/api/sub_category/update/{id}", subCategory.UpdateHandler).Methods("POST")
	r.HandleFunc("/api/sub_category/delete/{id}", subCategory.DeleteHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":3348", r))

}
