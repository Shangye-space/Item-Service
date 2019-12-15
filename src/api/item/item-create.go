package item

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shangye-space/Item-Service/src/api/helpers"
	database "github.com/Shangye-space/Item-Service/src/db"
	"github.com/Shangye-space/Item-Service/src/models"
)

// Create - creates item in db
func Create(w http.ResponseWriter, r *http.Request) {
	db, err := database.CreateDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Declare a new Item struct.
	var item models.Item

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	item.Name, err = helpers.CheckString(item.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	item.Price, err = helpers.CheckNumber(item.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var inSale int
	if item.InSale == nil {
		http.Error(w, "InSale is wrong", http.StatusBadRequest)
	} else {
		if *item.InSale == true {
			inSale = 1
		} else {
			inSale = 0
		}
	}

	if item.SubCategoryID == nil || *item.SubCategoryID <= 0 {
		http.Error(w, "SubCategoryID is wrong", http.StatusBadRequest)
	}

	query := fmt.Sprintf(`
	INSERT INTO item(name, price, sub_category_id, in_sale)
	VALUES("%v", %v, %v, %v);`, *item.Name, *item.Price, *item.SubCategoryID, inSale)

	fmt.Println(query)

	db.Exec(query)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
