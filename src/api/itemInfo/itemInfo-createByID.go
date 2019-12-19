package iteminfo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

// CreateByIDHandler - handles creating item info
func CreateByIDHandler(w http.ResponseWriter, r *http.Request) {
	var itemInfo models.ItemInfo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&itemInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ItemID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckNumberInt(itemInfo.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckString(itemInfo.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckNumber(itemInfo.Discount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckString(itemInfo.Size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckString(itemInfo.Color)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckString(itemInfo.Manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckString(itemInfo.ItemCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = helpers.CheckString(itemInfo.Material)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	CreateByID(&ItemID, itemInfo.Quantity, itemInfo.Description, itemInfo.Discount, itemInfo.Size, itemInfo.Color, itemInfo.Manufacturer, itemInfo.ItemCode, itemInfo.Material, db)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// CreateByID - creates item info in db by ID
func CreateByID(itemID *int, Quantity *int, Description *string, Discount *float32, Size *string, Color *string, Manufacturer *string, ItemCode *string, Material *string, db *sql.DB) {

	query := fmt.Sprintf(`INSERT INTO item_info (item_id, quantity, description, discount, size, color, manufacturer, item_code, material) 
	VALUES(%v, %v, "%v", %v, "%v", "%v", "%v", "%v", "%v")`, *itemID, *Quantity, *Description,
		*Discount, *Size, *Color, *Manufacturer, *ItemCode, *Material)

	db.Exec(query)
}
