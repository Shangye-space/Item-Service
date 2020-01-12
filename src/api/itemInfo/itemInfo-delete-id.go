package iteminfo

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	database "github.com/Shangye-space/Item-Service/src/db"
)

//DeleteByIDHandler - Handles removing item info from db
func DeleteByIDHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(w)

	itemID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	DeleteByID(itemID, db)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//DeleteByID - removes item info from db
func DeleteByID(itemID int, db *sql.DB) {
	query := fmt.Sprintf("DELETE FROM item_info WHERE item_id = %v;", strconv.Itoa(itemID))
	db.Exec(query)
}
