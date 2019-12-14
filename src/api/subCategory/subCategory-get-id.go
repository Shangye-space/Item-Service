package subcategory

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
	"github.com/gorilla/mux"
)

//GetByID - gets sub categories by ID
func GetByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if itemID == 0 {
		http.Error(w, "can't be 0", http.StatusBadRequest)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed")
	}

	result, err := db.Query(`SELECT * FROM sub_category WHERE id = %v LIMIT 1`)

	if err != nil {
		panic(err.Error())
	}

	var subCategory models.SubCategory
	var subCategories []models.SubCategory

	for result.Next() {
		err := result.Scan(&subCategory.ID, &subCategory.Name, &subCategory.CategoryID, &subCategory.AddedTime, &subCategory.LastUpdated, &subCategory.RemovedTime)
		if err != nil {
			panic(err.Error())
		}
		subCategories = append(subCategories, subCategory)
	}

	defer result.Close()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subCategories)

}
