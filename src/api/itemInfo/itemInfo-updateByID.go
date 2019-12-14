package iteminfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
	"github.com/gorilla/mux"
)

// UpdateByID - updates item info by id
func UpdateByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if itemID == 0 {
		http.Error(w, "can't be 0", http.StatusBadRequest)
	}

	var iteminfo models.ItemInfo

	decoder := json.NewDecoder(r.Body)

	err1 := decoder.Decode(&iteminfo)
	if err1 != nil {
		panic(err)
	}

	var setProp []string

	if iteminfo.Quantity != nil && *iteminfo.Quantity > 0 {
		prop := fmt.Sprintf(`quantity = %v`, *iteminfo.Quantity)
		setProp = append(setProp, prop)
	}

	if iteminfo.Description != nil && len(*iteminfo.Description) > 0 {
		prop := fmt.Sprintf(`description = "%v"`, *iteminfo.Description)
		setProp = append(setProp, prop)
	}

	if iteminfo.Discount != nil && *iteminfo.Discount > 0 {
		prop := fmt.Sprintf(`discount = %v`, *iteminfo.Discount)
		setProp = append(setProp, prop)
	}

	if iteminfo.Size != nil && len(*iteminfo.Size) > 0 {
		prop := fmt.Sprintf(`size = "%v"`, *iteminfo.Size)
		setProp = append(setProp, prop)
	}

	if iteminfo.Color != nil && len(*iteminfo.Color) > 0 {
		prop := fmt.Sprintf(`color = "%v"`, *iteminfo.Color)
		setProp = append(setProp, prop)
	}

	if iteminfo.Manufacturer != nil && len(*iteminfo.Manufacturer) > 0 {
		prop := fmt.Sprintf(`manufacturer = "%v"`, *iteminfo.Manufacturer)
		setProp = append(setProp, prop)
	}

	if iteminfo.ItemCode != nil && len(*iteminfo.ItemCode) > 0 {
		prop := fmt.Sprintf(`item_code = "%v"`, *iteminfo.ItemCode)
		setProp = append(setProp, prop)
	}

	if iteminfo.Material != nil && len(*iteminfo.Material) > 0 {
		prop := fmt.Sprintf(`material = "%v"`, *iteminfo.Material)
		setProp = append(setProp, prop)
	}

	db, err := database.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	lastUpdated := fmt.Sprintf(`last_updated = "%v"`, time.Now().Format("2006-01-02 15:04:05"))
	setProp = append(setProp, lastUpdated)

	query := string(fmt.Sprintf("UPDATE item_info SET %v WHERE item_id = %v;", strings.Join(setProp, ", "), strconv.Itoa(itemID)))
	db.Exec(query)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
