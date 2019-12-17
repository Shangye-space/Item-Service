package subcategory

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
)

// GetByIDHandler - Handles get method for subCategory by ID
func GetByIDHandler(w http.ResponseWriter, r *http.Request) {

	subCategoryID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	subCategory := GetByID(subCategoryID, db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subCategory)

}

//GetByID - Gets unique sub category by ID
func GetByID(subCategoryID int, db *sql.DB) models.SubCategory {
	query := string(fmt.Sprintf("SELECT * FROM sub_category WHERE id = %v LIMIT 1", strconv.Itoa(subCategoryID)))
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var subCategory models.SubCategory

	for result.Next() {
		err := result.Scan(&subCategory.ID, &subCategory.Name, &subCategory.CategoryID, &subCategory.AddedTime, &subCategory.LastUpdated, &subCategory.RemovedTime)
		if err != nil {
			panic(err.Error())
		}
	}

	defer result.Close()

	return subCategory
}
