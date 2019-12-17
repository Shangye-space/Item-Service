package iteminfo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	"github.com/Shangye-space/Item-Service/src/models"
)

// UpdateByIDHandler - Handles item info update function
func UpdateByIDHandler(w http.ResponseWriter, r *http.Request) {

	itemID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var itemInfo models.ItemInfo
	decoder := json.NewDecoder(r.Body)

	err1 := decoder.Decode(&itemInfo)
	if err1 != nil {
		panic(err)
	}

	var setProp []string

	if helpers.CheckNumberInt(itemInfo.Quantity) == nil {
		prop := fmt.Sprintf(`quantity = %v`, *itemInfo.Quantity)
		setProp = append(setProp, prop)
	}

	if helpers.CheckString(itemInfo.Description) == nil {
		prop := fmt.Sprintf(`description = "%v"`, *itemInfo.Description)
		setProp = append(setProp, prop)
	}

	if helpers.CheckNumber(itemInfo.Discount) == nil {
		prop := fmt.Sprintf(`discount = %v`, *itemInfo.Discount)
		setProp = append(setProp, prop)
	}

	if helpers.CheckString(itemInfo.Size) == nil {
		prop := fmt.Sprintf(`size = "%v"`, *itemInfo.Size)
		setProp = append(setProp, prop)
	}

	if helpers.CheckString(itemInfo.Color) == nil {
		prop := fmt.Sprintf(`color = "%v"`, *itemInfo.Color)
		setProp = append(setProp, prop)
	}

	if helpers.CheckString(itemInfo.Manufacturer) == nil {
		prop := fmt.Sprintf(`manufacturer = "%v"`, *itemInfo.Manufacturer)
		setProp = append(setProp, prop)
	}

	if helpers.CheckString(itemInfo.ItemCode) == nil {
		prop := fmt.Sprintf(`item_code = "%v"`, *itemInfo.ItemCode)
		setProp = append(setProp, prop)
	}

	if helpers.CheckString(itemInfo.Material) == nil {
		prop := fmt.Sprintf(`material = "%v"`, *itemInfo.Material)
		setProp = append(setProp, prop)
	}

	lastUpdated := fmt.Sprintf(`last_updated = "%v"`, time.Now().Format("2006-01-02 15:04:05"))
	setProp = append(setProp, lastUpdated)

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	UpdateByID(&itemID, itemInfo.Quantity, itemInfo.Description, itemInfo.Discount, itemInfo.Size, itemInfo.Color, itemInfo.Manufacturer, itemInfo.ItemCode, itemInfo.Material, db, setProp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

//UpdateByID - updates item info by id
func UpdateByID(itemID *int, Quantity *int, Description *string, Discount *float32, Size *string, Color *string, Manufacturer *string, ItemCode *string, Material *string, db *sql.DB, setProp []string) {

	query := string(fmt.Sprintf("UPDATE item_info SET %v WHERE item_id = %v;", strings.Join(setProp, ", "), strconv.Itoa(*itemID)))
	db.Exec(query)

}
