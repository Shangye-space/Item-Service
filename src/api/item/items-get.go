package item

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

// GetHandler - Handles GET method for items
func GetHandler(w http.ResponseWriter, r *http.Request) {

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	items := Get(db)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Get - Gets Items
func Get(db *sql.DB) []models.Item {

	result, err := db.Query(`SELECT * FROM item`)
	if err != nil {
		panic(err.Error())
	}

	items := helpers.ScanItems(result)
	defer result.Close()
	return items
}
