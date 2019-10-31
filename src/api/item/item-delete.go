package item

import (
	"fmt"
	"net/http"
	"strconv"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/gorilla/mux"
)

// Delete - removes item from db
func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if itemID == 0 {
		http.Error(w, "can't be 0", http.StatusBadRequest)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	query := fmt.Sprintf("DELETE FROM Items WHERE ItemID = %v;", strconv.Itoa(itemID))

	db.Exec(query)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
