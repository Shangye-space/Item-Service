package item

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

// GetByCategoryIDHandler - Handles getting Items by Category ID
func GetByCategoryIDHandler(w http.ResponseWriter, r *http.Request) {
	categoryID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed.")
	}

	items := GetByCategoryID(categoryID, db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GetByCategoryID - Gets Item by Category ID
func GetByCategoryID(categoryID int, db *sql.DB) []models.Item {

	query := string(fmt.Sprintf("SELECT item.`id`, item.`name`, item.`price`, item.`sub_category_id`, item.`in_sale`, item.`added_time`, item.`last_updated`, item.`removed_time` FROM item INNER JOIN sub_category ON item.`sub_category_id` = sub_category.`id` INNER JOIN category ON sub_category.`category_id` = category.`id` WHERE category_id = %v;", categoryID))
	result, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	items := helpers.ScanItems(result)
	defer result.Close()
	return items
}
