package category

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	helpers "github.com/Shangye-space/Item-Service/src/api/helpers"
)

// DeleteHandler - Handles removing category from db
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	categoryID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Delete(categoryID, db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// Delete - removes category from db
func Delete(categoryID int, db *sql.DB) {
	query := fmt.Sprintf("DELETE FROM category WHERE id = %v;", strconv.Itoa(categoryID))

	db.Exec(query)
}
