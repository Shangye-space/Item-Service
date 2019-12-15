package item

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/models"

	database "github.com/Shangye-space/Item-Service/src/db"
)

// GetHandler - Handles GET method for items
func GetHandler(w http.ResponseWriter, r *http.Request) {

	db, err := database.CreateDatabase()
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

	var item models.Item
	var items []models.Item

	for result.Next() {
		err := result.Scan(&item.ID, &item.Name, &item.Price, &item.SubCategoryID, &item.InSale, &item.AddedTime, &item.LastUpdated, &item.RemovedTime)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, item)
	}

	defer result.Close()

	return items
}
