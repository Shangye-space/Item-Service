package iteminfo

import (
	"fmt"
	"net/http"
	"strconv"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/gorilla/mux"
)

//DeleteByID - removes item from db by id
func DeleteByID(w http.ResponseWriter, r *http.Request) {

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

	query := fmt.Sprintf("DELETE FROM item_info WHERE item_id = %v;", strconv.Itoa(itemID))
	fmt.Printf(query)
	db.Exec(query)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
