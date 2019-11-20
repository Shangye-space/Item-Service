package subcategory

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
)

//Get - gets sub categories
func Get(w http.ResponseWriter, r *http.Request) {

	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed")
	}

	result, err := db.Query(`SELECT * FROM sub_category`)

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
