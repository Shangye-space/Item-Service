package subcategory

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

// GetSubCategoryByCategoryIDHandler - handles get method for subCategory by CategoryID
func GetSubCategoryByCategoryIDHandler(w http.ResponseWriter, r *http.Request) {

	categoryID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		log.Fatal("Connection to DB has failed.")
	}

	subCategory := GetSubCategoryIDByCategoryID(categoryID, db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subCategory)
}

// GetSubCategoryIDByCategoryID - Gets subcategory IDs by category ID
func GetSubCategoryIDByCategoryID(categoryID int, db *sql.DB) []models.SubCategory {
	query := string(fmt.Sprintf("SELECT * FROM sub_category WHERE category_id = %v", strconv.Itoa(categoryID)))
	result, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	subCategory := helpers.ScanSubCategories(result)
	defer result.Close()

	return subCategory

}
