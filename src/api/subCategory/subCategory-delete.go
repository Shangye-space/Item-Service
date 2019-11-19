package subcategory

import (
	"fmt"
	"net/http"
	"strconv"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/gorilla/mux"
)

// Delete - removes sub categories from db
func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	subCategoryID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if subCategoryID == 0 {
		http.Error(w, "can't be 0", http.StatusBadRequest)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	query := fmt.Sprintf("DELETE FROM sub_category WHERE id = %v;", strconv.Itoa(subCategoryID))

	db.Exec(query)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
