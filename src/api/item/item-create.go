package item

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

// CreateHandler - handles creating item
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckString(item.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckNumber(item.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	inSale, err := helpers.CheckBool(item.InSale)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckID(item.SubCategoryID)

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Create(item, inSale, db)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// Create - creates item
func Create(item models.Item, inSale int, db *sql.DB) {
	query := fmt.Sprintf(`
	INSERT INTO item(name, price, sub_category_id, in_sale)
	VALUES("%v", %v, %v, %v);`, *item.Name, *item.Price, *item.SubCategoryID, inSale)
	db.Exec(query)
}
