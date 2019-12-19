package subcategory

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
)

// DeleteHandler - Handles removing sub category from db
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	subCategoryID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Delete(subCategoryID, db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// Delete - removes sub category from db
func Delete(subCategoryID int, db *sql.DB) {
	query := fmt.Sprintf("DELETE FROM sub_category WHERE id = %v;", strconv.Itoa(subCategoryID))

	db.Exec(query)
}
