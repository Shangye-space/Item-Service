package item

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
)

// GetCountHandler - Handles GET method for item
func GetCountHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	itemCount := GetCount(db)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemCount)
}

// GetCount - Gets Items
func GetCount(db *sql.DB) int {

	var itemCount int
	_ = db.QueryRow(`SELECT COUNT(id) FROM item;`).Scan(&itemCount)

	return itemCount
}
