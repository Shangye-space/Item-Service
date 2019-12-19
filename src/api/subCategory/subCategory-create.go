package subcategory

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

//CreateHandler - handles creating sub category
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	var subCategory models.SubCategory

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&subCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckString(subCategory.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckID(subCategory.CategoryID)

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Create(subCategory, db)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// Create - creates sub categories in db
func Create(subCategory models.SubCategory, db *sql.DB) {
	query := fmt.Sprintf(`INSERT INTO sub_category(name, category_id)
	VALUES("%v", %v);`, *subCategory.Name, *subCategory.CategoryID)

	db.Exec(query)
}
