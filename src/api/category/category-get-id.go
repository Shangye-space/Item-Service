package category

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
	"github.com/gorilla/mux"
)

// GetByID - Gets Category by ID
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
		log.Fatal("Connection to DB has failed")
	}

	result, err := db.Query(`SELECT * FROM category WHERE id = %v LIMIT 1`)

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
