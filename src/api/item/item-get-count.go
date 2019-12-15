package item

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/Shangye-space/Item-Service/src/db"
)

// GetCountHandler - Handles GET method for item
func GetCountHandler(w http.ResponseWriter, r *http.Request) {

	itemCount := GetCount()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemCount)
}

// GetCount - Gets Items
func GetCount() int {

	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed.")
	}

	var itemCount int
	_ = db.QueryRow(`SELECT COUNT(id) FROM item;`).Scan(&itemCount)

	return itemCount
}
