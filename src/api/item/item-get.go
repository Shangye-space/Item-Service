package item

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/models"

	database "github.com/Shangye-space/Item-Service/src/db"
)

// Get - Gets Items
func Get(w http.ResponseWriter, r *http.Request) {

	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed.")
	}

	result, err := db.Query(`
	SELECT * FROM item`)

	if err != nil {
		panic(err.Error())
	}

	var item models.Item
	var items []models.Item

	for result.Next() {
		err := result.Scan(&item.ItemID, &item.SubCategoryID, &item.ItemName, &item.InSale, &item.AddedTime, &item.LastUpdated, &item.RemovedTime)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, item)
	}

	defer result.Close()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
