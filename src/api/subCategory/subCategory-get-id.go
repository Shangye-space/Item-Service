package subcategory

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

// GetByIDHandler - Handles get method for subCategory by ID
func GetByIDHandler(w http.ResponseWriter, r *http.Request) {

	subCategoryID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	subCategory := GetByID(subCategoryID, db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subCategory)

}

//GetByID - Gets unique sub category by ID
func GetByID(subCategoryID int, db *sql.DB) []models.SubCategory {
	query := string(fmt.Sprintf("SELECT * FROM sub_category WHERE id = %v LIMIT 1", strconv.Itoa(subCategoryID)))
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	subCategory := helpers.ScanSubCategories(result)
	defer result.Close()

	return subCategory
}
