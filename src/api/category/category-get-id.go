package category

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	helpers "github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

// GetByIDHandler - Handles get method for Category by ID
func GetByIDHandler(w http.ResponseWriter, r *http.Request) {

	categoryID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	category := GetByID(categoryID, db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)

}

//GetByID - Gets unique category by ID
func GetByID(categoryID int, db *sql.DB) []models.Category {
	query := string(fmt.Sprintf("SELECT * FROM category WHERE id = %v LIMIT 1", strconv.Itoa(categoryID)))
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	category := helpers.ScanCategories(result)
	defer result.Close()

	return category
}
