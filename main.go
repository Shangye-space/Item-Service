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
	r.HandleFunc("/api/item_info/{id}", iteminfo.GetByID).Methods("GET")
	r.HandleFunc("/api/item_info/create/{id}", iteminfo.CreateByID).Methods("POST")
	r.HandleFunc("/api/item_info/update/{id}", iteminfo.UpdateByID).Methods("POST")
	r.HandleFunc("/api/item_info/delete/{id}", iteminfo.DeleteByID).Methods("GET")

	/* Image */
	r.HandleFunc("/api/image/{id}", image.GetByIDHandler).Methods("GET")
	//r.HandleFunc("/api/image/upload", image.UploadHandler).Methods("POST")
	r.HandleFunc("/api/image/create/{id}", image.CreateHandler).Methods("POST")
	r.HandleFunc("/api/image/update/{id}", image.UpdateHandler).Methods("POST")
	r.HandleFunc("/api/image/delete/{id}", image.DeleteHandler).Methods("GET")
	r.HandleFunc("/api/image/return/{id}", image.ReturnImage).Methods("GET")
	r.HandleFunc("/api/image/return_link/{id}", image.ReturnImageLink).Methods("GET")
	r.HandleFunc("/api/image/return_links/{ids}", image.ReturnImageLinks).Methods("GET")

	/* Categories */
	r.HandleFunc("/api/categories", category.GetHandler).Methods("GET")
	r.HandleFunc("/api/category/{id}", category.GetByIDHandler).Methods("GET")
	r.HandleFunc("/api/category/create", category.Create).Methods("POST")
	r.HandleFunc("/api/category/update/{id}", category.Update).Methods("POST")
	r.HandleFunc("/api/category/delete/{id}", category.Delete).Methods("GET")

	/* Sub Categories */
	r.HandleFunc("/api/sub_categories", subCategory.Get).Methods("GET")
	r.HandleFunc("/api/sub_category/{id}", subCategory.GetByIDHandler).Methods("GET")
	r.HandleFunc("/api/sub_categories/category/{id}", subCategory.GetSubCategoryIDByCategoryID).Methods("GET")
	r.HandleFunc("/api/sub_category/create", subCategory.Create).Methods("POST")
	r.HandleFunc("/api/sub_category/update/{id}", subCategory.Update).Methods("POST")
	r.HandleFunc("/api/sub_category/delete/{id}", subCategory.Delete).Methods("GET")

	log.Fatal(http.ListenAndServe(":3348", r))
}
