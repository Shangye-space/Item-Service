package item

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
)

// DeleteHandler - Handles removing item from db
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	itemID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Delete(itemID, db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// Delete - removes item from db
func Delete(itemID int, db *sql.DB) {
	query := fmt.Sprintf("DELETE FROM item WHERE id = %v;", strconv.Itoa(itemID))
	db.Exec(query)
}
