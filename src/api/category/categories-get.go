package category

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

// GetHandler - Handles GET method for categories
func GetHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	db, err := helpers.CreateDatabase()
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

	categories := helpers.ScanCategories(result)
	defer result.Close()

	return categories

}
