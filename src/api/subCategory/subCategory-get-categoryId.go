package subcategory

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Shangye-space/Item-Service/src/models"
	"github.com/gorilla/mux"

	database "github.com/Shangye-space/Item-Service/src/db"
)

// GetSubCategoryIDByCategoryID - Gets subcategory IDs by category ID
func GetSubCategoryIDByCategoryID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	categoryID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if categoryID == 0 {
		http.Error(w, "can't be 0", http.StatusBadRequest)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed.")
	}

	fmt.Println(categoryID)
	query := string(fmt.Sprintf("SELECT id FROM sub_category WHERE category_id = %v", strconv.Itoa(categoryID)))
	result, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var subCategory models.SubCategory
	var subCategories []models.SubCategory

	for result.Next() {
		err := result.Scan(&subCategory.ID)
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
