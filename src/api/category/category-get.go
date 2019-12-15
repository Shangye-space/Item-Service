package category

import (
	"database/sql"
	"encoding/json"
	"net/http"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
)

// GetHandler - Handles GET method for categories
func GetHandler(w http.ResponseWriter, r *http.Request) {

	db, err := database.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	categories := Get(db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)

}

// Get - Gets Categories
func Get(db *sql.DB) []models.Category {

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

	return categories

}
