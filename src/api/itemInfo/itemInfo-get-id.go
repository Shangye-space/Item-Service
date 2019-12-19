package iteminfo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

// GetByIDHandler - Handles get method for Item info by ID
func GetByIDHandler(w http.ResponseWriter, r *http.Request) {

	itemID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		log.Fatal("Conection to DB has failed")
	}

	itemInfo := GetByID(itemID, db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemInfo)
}

// GetByID - GetByID Item Info
func GetByID(itemID int, db *sql.DB) []models.ItemInfo {
	query := string(fmt.Sprintf("SELECT * FROM item_info WHERE item_id = %v", strconv.Itoa(itemID)))
	result, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	itemInfo := helpers.ScanItemInfo(result)
	defer result.Close()
	return itemInfo
}
