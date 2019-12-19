package item

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

// GetBySubCategoryIDHandler - Handles Get Item by SubCategory ID function
func GetBySubCategoryIDHandler(w http.ResponseWriter, r *http.Request) {
	subCategoryID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	items := GetBySubCategoryID(subCategoryID, db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

//GetBySubCategoryID - Gets Items by SubCategory ID
func GetBySubCategoryID(subCategoryID int, db *sql.DB) []models.Item {

	query := string(fmt.Sprintf("SELECT * FROM item WHERE sub_category_id = %v", strconv.Itoa(subCategoryID)))
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	items := helpers.ScanItems(result)
	defer result.Close()
	return items
}
