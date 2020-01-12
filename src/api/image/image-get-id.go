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

// GetByIDHandler - Handles getting images by imageID
func GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(w)

	itemID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	images, err := GetByID(itemID, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

//GetByID -
func GetByID(itemID int, db *sql.DB) ([]models.Image, error) {

	query := string(fmt.Sprintf("SELECT * FROM image WHERE item_id = %v LIMIT 1", strconv.Itoa(itemID)))
	result, err := db.Query(query)
	if err != nil {
		panic("error occured")
	}

	images := helpers.ScanImage(result)

	return images, err
}
