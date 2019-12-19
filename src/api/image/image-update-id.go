package image

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

//UpdateHandler - handles updating image
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	itemID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var image models.Image
	decoder := json.NewDecoder(r.Body)
	err1 := decoder.Decode(&image)
	if err1 != nil {
		panic(err1)
	}

	var path string

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Update(itemID, db, path)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//Update - updates item
func Update(itemID int, db *sql.DB, newPath string) {

	query := string(fmt.Sprintf("UPDATE item SET %v WHERE id = %v;", newPath, strconv.Itoa(itemID)))
	db.Exec(query)
}
