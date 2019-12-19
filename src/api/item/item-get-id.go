package item

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Shangye-space/Item-Service/src/models"
	"github.com/gorilla/mux"

	database "github.com/Shangye-space/Item-Service/src/db"
)

// GetByID - Gets Item by ID
func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if itemID == 0 {
		http.Error(w, "can't be 0", http.StatusBadRequest)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed.")
	}

	fmt.Println(itemID)
	query := string(fmt.Sprintf("SELECT * FROM item WHERE id = %v LIMIT 1", strconv.Itoa(itemID)))
	result, err := db.Query(query)

	fmt.Println(result)

	if err != nil {
		panic(err.Error())
	}

	var item models.Item
	var items []models.Item

	for result.Next() {
		err := result.Scan(&item.ID, &item.Name, &item.SubCategoryID, &item.InSale, &item.AddedTime, &item.LastUpdated, &item.RemovedTime)
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
