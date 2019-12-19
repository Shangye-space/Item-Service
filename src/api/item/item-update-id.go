package item

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

// UpdateHandler - Handles item update function
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	itemID, err := helpers.CheckIDWithRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var item models.Item
	decoder := json.NewDecoder(r.Body)
	err1 := decoder.Decode(&item)
	if err1 != nil {
		panic(err1)
	}

	var setProp []string

	if helpers.CheckString(item.Name) == nil {
		prop := fmt.Sprintf(`name = "%v"`, *item.Name)
		setProp = append(setProp, prop)
	}

	if helpers.CheckNumber(item.Price) == nil {
		prop := fmt.Sprintf(`price = %v`, *item.Price)
		setProp = append(setProp, prop)
	}

	inSaleNum, err := helpers.CheckBool(item.InSale)
	if err == nil {
		prop := fmt.Sprintf(`in_sale = %v`, inSaleNum)
		setProp = append(setProp, prop)
	}

	if helpers.CheckID(item.SubCategoryID) == nil {
		prop := fmt.Sprintf(`sub_category_id = %v`, *item.SubCategoryID)
		setProp = append(setProp, prop)
	}

	lastUpdated := fmt.Sprintf(`last_updated = "%v"`, time.Now().Format("2006-01-02 15:04:05"))
	setProp = append(setProp, lastUpdated)

	db, err := helpers.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Update(itemID, db, setProp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//Update - updates item
func Update(itemID int, db *sql.DB, setProp []string) {

	query := string(fmt.Sprintf("UPDATE item SET %v WHERE id = %v;", strings.Join(setProp, ", "), strconv.Itoa(itemID)))
	db.Exec(query)
}
