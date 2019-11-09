package item

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	err1 := decoder.Decode(&item)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
	}

	if item.ItemName == nil || len(*item.ItemName) <= 0 {
		http.Error(w, "ItemName is wrong", http.StatusBadRequest)
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
	INSERT INTO item(sub_category_id, item_name, in_sale)
	VALUES(%v, "%v", %v);`, *item.SubCategoryID, *item.ItemName, inSale)

	db.Exec(query)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
