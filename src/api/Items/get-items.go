package items

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/models"

	database "github.com/Shangye-space/Item-Service/src/db"
)

// GetItems - Gets Items
func GetItems(w http.ResponseWriter, r *http.Request) {

	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed.")
	}

	result, err := db.Query("SELECT * FROM Items;")
	if err != nil {
		panic(err.Error())
	}

	var item models.Items
	var items []models.Items

	for result.Next() {
		err := result.Scan(&item.ItemID, &item.ItemName, &item.Quantity, &item.Description, &item.Price, &item.Discount, &item.InSale, &item.Category, &item.SubCategory, &item.AddedTime, &item.RemovedTime)
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
