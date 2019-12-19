package subcategory

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

// GetHandler - Handles GET method for sub categories
func GetHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	subCategories := Get(db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subCategories)

}

//Get - gets sub sub categories
func Get(db *sql.DB) []models.SubCategory {
	result, err := db.Query(`SELECT * FROM sub_category`)

	if err != nil {
		panic(err.Error())
	}

	subCategories := helpers.ScanSubCategories(result)
	defer result.Close()

	return subCategories
}
