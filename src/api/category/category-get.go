package category

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
)

// Get - Gets Categories
func Get(w http.ResponseWriter, r *http.Request) {

	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed")
	}

	result, err := db.Query(`SELECT * FROM category`)

	if err != nil {
		panic(err.Error())
	}

	var category models.Category
	var categories []models.Category

	for result.Next() {
		err := result.Scan(&category.ID, &category.Name, &category.AddedTime, &category.LastUpdated, &category.RemovedTime)
		if err != nil {
			panic(err.Error())
		}
		categories = append(categories, category)
	}

	defer result.Close()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)

}
