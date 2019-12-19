package item

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	helpers "github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

// GetByIDHandler - Handles get method for Item by ID
func GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	itemID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	item := GetByID(itemID, db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

//GetByID - Gets unique item by ID
func GetByID(itemID int, db *sql.DB) []models.Item {
	query := string(fmt.Sprintf("SELECT * FROM item WHERE id = %v LIMIT 1", strconv.Itoa(itemID)))
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	item := helpers.ScanItems(result)
	defer result.Close()
	return item
}
