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

	if item.Name == nil || len(*item.Name) <= 0 {
		http.Error(w, "Name is wrong", http.StatusBadRequest)
	}

	if item.Price == nil || *item.Price <= 0 {
		http.Error(w, "Price is wrong", http.StatusBadRequest)
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
	INSERT INTO item(sub_category_id, name, price, in_sale)
	VALUES(%v, "%v", %v, %v);`, *item.SubCategoryID, *item.Name, *item.Price, inSale)

	db.Exec(query)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
